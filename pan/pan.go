package pan

import (
	"bytes"
	"context"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/sevki/google-roughtime/go/config"

	"github.com/sevki/roughtime"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"sevki.org/x/debug"
	"willnorris.com/go/newbase60"
)

var (
	l = log.New(os.Stdout, "pan: ", 0)
)

type pan struct {
	fs http.FileSystem

	servers   []config.Server
	prev      *roughtime.Roughtime
	cache     string
	templates map[string]string
}

// New returns a new TitanFS
func New(fs http.FileSystem, options ...Option) http.FileSystem {
	servers, _, _ := roughtime.LoadConfig("/go/src/github.com/cloudflare/roughtime/ecosystem.json")

	dir, _ := ioutil.TempDir("", "pan")
	p := &pan{
		fs:        fs,
		cache:     dir,
		templates: make(map[string]string),
		servers:   servers,
	}
	for _, option := range options {
		option(p)
	}
	return p
}

type Option func(*pan)

func WithTemplate(format, path string) Option {
	return func(p *pan) { p.templates[format] = path }
}

func (p *pan) Open(name string) (http.File, error) {
	oname := name
	const (
		none = ""
		html = ".html"
		pdf  = ".pdf"
	)

	ext := path.Ext(name)
	namePlain := strings.TrimSuffix(name, ext)
	l.Printf("get: name=%s", name)

	switch ext {
	case pdf, html:
		mname := fmt.Sprintf("%s.%s", namePlain, "md")
		f, err := p.fs.Open(mname)
		if err == nil {
			//	_, fileName := path.Split(oname)
			l.Printf("render: name=%s", name)
			return p.render(name, oname, f)
		}
	}

	l.Printf("open: name=%s", name)
	f, err := p.fs.Open(name)
	if err != nil {
		l.Printf("open errored: name=%s err=%v", name, err)
	}
	return f, err
}

type saturnFile struct {
	http.File

	name        string
	contentType string
}

type saturnInfo struct {
	os.FileInfo

	name string
}

func (f *saturnFile) ContentType() string {
	switch f.contentType {
	case "html5":
		return "text/html"
	case "latex":
		return "application/pdf"

	}
	return "application/octet-stream"
}

func (i *saturnInfo) Name() string {
	l.Println(i.name)
	return i.name
}

func (f *saturnFile) Stat() (os.FileInfo, error) {
	stat, err := f.File.Stat()
	if err != nil {
		return nil, err
	}
	return &saturnInfo{stat, f.name}, nil
}

func hash(s string) string {
	h := fnv.New32a()
	io.WriteString(h, s)
	a := int(h.Sum32())
	return newbase60.EncodeInt(a)
}

func (p *pan) render(name, fake string, f http.File) (http.File, error) {

	stt, _ := f.Stat()
	if stt.IsDir() {
		return f, nil
	}

	bytz, err := ioutil.ReadAll(f)
	if err != nil {
		l.Println(errors.Wrap(err, "pan read all"))

		return nil, errors.Wrap(err, "pan read all")
	}

	dir, file := path.Split(name)

	cache := path.Join("/x/", hash(name)+"-"+stt.ModTime().Format("20060102150405"))
	outputFormat := "html5"
	switch path.Ext(name) {
	case ".pdf":
		outputFormat = "latex"
	}
	if fx, err := os.Open(path.Join(cache, file)); err == nil {
		return &saturnFile{fx, name, outputFormat}, nil
	}

	buf := bytes.NewBuffer(nil)

	os.Mkdir(cache, 0644)

	{ // run pre processor
		l.Println("pp", "running")

		args := []string{fmt.Sprintf("-img=%s", cache)}
		cmd := exec.CommandContext(context.Background(), "pp", args...)
		cmd.Stdin = bytes.NewBuffer(bytz)
		stdErr := bytes.NewBuffer(nil)
		cmd.Stdout = buf
		cmd.Stderr = stdErr
		if err := cmd.Run(); err != nil {
			debug.Indent(stdErr, 1)
			l.Printf(`pp: %v
%s`, args, stdErr.String())
			l.Println("pp", err.Error())
			return nil, err
		}
	}

	{ // run pandoc

		results := roughtime.Do(p.servers, 4, time.Second*1, p.prev)
		chain := roughtime.NewChain(results)
		if cool, err := chain.Verify(p.prev); err != nil && cool {
			return nil, err
		}
		p.prev = results[0].Roughtime
		args := []string{
			"-f",
			"markdown",
			"-V", fmt.Sprintf("permalink:%s", dir),
			"-V", fmt.Sprintf("rt:%s", results[0].String()),
		}

		args = append(args, "-o", path.Join(cache, file))

		args = append(args, "-t", outputFormat)
		if path, ok := p.templates[outputFormat]; ok {
			args = append(args, "--template", path)
		} else {
			args = append(args, "-s")
		}

		cmd := exec.CommandContext(context.Background(), "pandoc", args...)
		cmd.Stdin = buf
		stdOut := bytes.NewBuffer(nil)
		stdErr := bytes.NewBuffer(nil)
		cmd.Stdout = stdOut
		cmd.Stderr = stdErr
		fmt.Printf("running pandoc: %v\n", args)
		if err := cmd.Run(); err != nil {
			debug.Indent(stdErr, 1)
			l.Printf(`pandoc: %v
	args= %v
%s
`,
				err,
				args,
				stdErr.String(),
			)
			return nil, err
		}

	}
	fx, _ := os.Open(path.Join(cache, file))
	return &saturnFile{fx, name, outputFormat}, nil
}

func h1(r io.Reader) string {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		tok := z.Token()
		if tok.Type == html.StartTagToken && tok.Data == "h1" {
			if z.Next() == html.TextToken {
				return z.Token().Data
			}
		}
	}
	return ""
}

func pandocVersion() string {
	cmd := exec.Command("pandoc", "--version")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "deafbeef"
	}
	return string(stdoutStderr)
}

func version() string {
	cmd := exec.Command("git", "describe", "--always")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "deafbeef"
	}
	return string(stdoutStderr)
}
