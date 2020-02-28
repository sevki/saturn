package main

import (
	"bufio"
	"bytes"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/araddon/dateparse"
	"github.com/gorilla/feeds"
	yaml "gopkg.in/yaml.v2"
	"sevki.org/saturn/pan"
	"sevki.org/saturn/qr"
	"sevki.org/x/pretty"
)

var (
	l  = log.New(os.Stdout, "atlas: ", 0)
	fs http.FileSystem
)

var (
	root = flag.String("root", "/pub", "root")
)

func main() {
	p := pan.New(
		http.Dir(*root),
		pan.WithTemplate("html5", "/sevki.org/saturn/templates/blog.html"),
		pan.WithTemplate("latex", "/sevki.org/saturn/templates/default.latex"),
	)
	a := &atlas{fs: p}
	go a.reload()
	http.Handle("/x/", http.StripPrefix("/x/", http.FileServer(http.Dir("/x"))))
	http.Handle("/qr/", http.StripPrefix("/qr/", http.HandlerFunc(qr.Qart)))

	http.Handle("/", a)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Exit(0)
	}()

	l.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}

type atlas struct {
	fs http.FileSystem

	mu     sync.Mutex
	posts  []post
	shorts map[string]string
}

type postTime struct {
	time.Time
}

// YYYY-MM-DD hh:mm:ss tz
const fuzzyFormat = "2006-01-02 15:04:05-07:00"

func (t *postTime) UnmarshalText(text []byte) error {
	postTime, err := dateparse.ParseAny(string(text))
	if err != nil {
		log.Println(err)
	}
	t.Time = postTime
	return nil
}

type post struct {
	Title  string
	Date   postTime
	Slug   string
	Author []struct {
		Name        string
		Affiliation string
	}
	Abstract string
	Tags     []string
}
type byDate []post

func (a byDate) Len() int           { return len(a) }
func (a byDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDate) Less(i, j int) bool { return a[i].Date.After(a[j].Date.Time) }

func (a *atlas) reload() {
	for {
		f, _ := a.fs.Open("/")

		files, err := f.Readdir(1000)
		if err != nil && len(files) < 1 {
			l.Println("read dir", err)
			time.Sleep(time.Second * 30)
			continue
		}
		f.Close()
		x := []post{}
		rdr := make(map[string]string)

		for _, file := range files {
			if file.IsDir() {
				slug := file.Name()
				f, err = a.fs.Open(path.Join(slug, "index.md"))
				if err != nil {
					continue
				}
				p := parseHeader(f)
				p.Slug = slug
				rdr[qr.Hash(slug)] = slug
				l.Println(qr.Hash(slug), slug)
				x = append(x, p)
				f.Close()
			}
		}
		sort.Sort(byDate(x))
		a.posts = x
		a.shorts = rdr
		time.Sleep(time.Second * 30)
	}

}
func parseHeader(r io.Reader) post {
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(r)
	yamlFiles := 0
	for scanner.Scan() {
		buf.Write(append(scanner.Bytes(), '\n'))
		if strings.TrimSpace(scanner.Text()) == "---" {
			yamlFiles++
		}
		if yamlFiles > 1 {
			break
		}
	}
	dec := yaml.NewDecoder(buf)
	var p = post{}
	if err := dec.Decode(&p); err != nil {
		log.Println(err)
	}
	return p
}

var (
	redirPath = regexp.MustCompile("/([[:alnum:]]{5})#")
)

func (a *atlas) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Path) >= 6 {
		possibleCode := r.URL.Path[1:6]
		l.Println(pretty.JSON(a.shorts), possibleCode)
		if slug, ok := a.shorts[possibleCode]; ok {
			http.Redirect(w, r, "https://"+path.Join("fistfulofbytes.com/", slug), http.StatusTemporaryRedirect)
			return
		}
	}

	switch r.URL.Path {
	case "/_status":
		io.WriteString(w, "OK.")
		return
	case "/feed.atom", "/feed.rss", "/feed.json":
		a.feed(w, r)
		return
	case "/":
		a.index(w, r)
		return
	}

	ext := path.Ext(r.URL.Path)

	panServer := http.FileServer(a.fs)
	switch ext {
	case ".latex":
		w.Header().Set("Content-Type", "application/x-latex")
		panServer.ServeHTTP(w, r)
	case ".pdf":
		w.Header().Set("Content-Type", "application/pdf")
		panServer.ServeHTTP(w, r)
	case ".html":
		w.Header().Set("Content-Type", "text/html")
		panServer.ServeHTTP(w, r)
	default:
		panServer.ServeHTTP(w, r)
	}
}
func (a *atlas) redir(w http.ResponseWriter, r *http.Request) {

}
func (a *atlas) feed(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	u, _ := url.Parse("https://fistfulofbytes.com")
	feed := &feeds.Feed{
		Title:       "fistful of bytes",
		Link:        &feeds.Link{Href: u.String()},
		Description: "sevki's blog",
		Author:      &feeds.Author{Name: "Sevki", Email: "s@sevki.org"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{}
	for _, post := range a.posts {
		x := *u
		x.Path = post.Slug
		feed.Items = append(feed.Items, &feeds.Item{
			Id:          qr.Hash(post.Slug),
			Title:       post.Title,
			Link:        &feeds.Link{Href: x.String() +"/"},
			Description: post.Abstract,
			Author: &feeds.Author{
				Name:  post.Author[0].Name,
				Email: post.Author[0].Affiliation,
			},
			Created: post.Date.Time,
		})
	}
	switch r.URL.Path {
	case "/feed.atom":
		feed.WriteAtom(w)
		return
	case "/feed.rss":
		feed.WriteRss(w)
		return
	case "/feed.json":
		feed.WriteJSON(w)
		return
	}
}
func (a *atlas) index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if a.posts == nil {
		io.WriteString(w, "Server not ready")
		return
	}

	instanceName, _ := metadata.InstanceID()
	containterName, _ := os.Hostname()
	pageTemplate, err := template.New("").Parse(index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = pageTemplate.Execute(w, map[string]interface{}{
		"Files":         a.posts,
		"RenderTime":    time.Now().Sub(start),
		"RenderedAt":    time.Now(),
		"GoVersion":     runtime.Version(),
		"Container":     containterName,
		"Instance":      instanceName,
		"PandocVersion": pandocVersion(),
	})
	return
}

var index = `<html>
	<head>
		<title>fistfulofbytes</title>
		<link rel="stylesheet" type="text/css" href="/style.css">
  		<link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/octicons/4.4.0/font/octicons.css">
		<link rel="icon" type="image/png" href="/qr/">

	</head>
	<body>
  <header>
    <a href="/"><img src="/qr/" alt="qr" class="qr"/></a>
        <div>
          <h1>
            <a href="/">fistfulofbytes</a>
          </h1>
	<ul>
		<li><a href="/">Index</a></li>
		<li><a href="https://sevki.io">About</a></li>
		<li>
			<details>
 				<summary>Feed <span style="" class="octicon octicon-rss"></span></summary>
				<ul>
					<li><a href="/feed.atom">Atom</a></li>
					<li><a href="/feed.rss">RSS</a></li>
					<li><a href="/feed.json">JSON</a></li>
				</ul>
			</details>
		</li>
	</ul>
        </div>
      </header>
	<main>
		{{- range .Files -}}
 		<p>
			<a href="/{{ .Slug }}" >{{ .Title }}</a>
			<br/>
			{{ range .Author }} {{ .Name }} {{ end }}
			<br/>
			{{ .Date.Format "2 Jan 2006" }}

		</p>
 		{{- end -}}
	</main>
		<footer>
		    <p>
			Except as noted,
			the content of this page is licensed under
			the <a href="//creativecommons.org/licenses/by/3.0/legalcode">
			Creative Commons Attribution 3.0 License</a>
		    </p>
		</footer>
	</body>
</html>
`

func pandocVersion() string {
	cmd := exec.Command("pandoc", "--version")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "deafbeef"
	}
	return string(stdoutStderr)
}