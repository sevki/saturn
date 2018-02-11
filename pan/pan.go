package pan

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"sevki.org/saturn/atlas"
)

type pan struct {
	fs    http.FileSystem
	cache string
}

// New returns a new TitanFS
func New(fs http.FileSystem) http.FileSystem {
	dir, _ := ioutil.TempDir("", "pan")
	return &pan{fs: fs, cache: dir}
}

func (p *pan) Open(name string) (http.File, error) {
	start := time.Now()
	f, err := p.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stt, _ := f.Stat()
	if atlasSys, ok := stt.Sys().(atlas.SysInfo); ok && !atlasSys.ShouldRender {
		return f, nil
	} else if !ok {
		return f, nil
	}

	bytz, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "titan read all")
	}
	buf := bytes.NewBuffer(nil)

	{ // run pre processor
		args := []string{fmt.Sprintf("-img=%s", p.cache)}
		cmd := exec.CommandContext(context.Background(), "pp", args...)
		cmd.Stdin = bytes.NewBuffer(bytz)
		stdErr := bytes.NewBuffer(nil)
		cmd.Stdout = buf
		cmd.Stderr = stdErr
		if err := cmd.Run(); err != nil {
			log.Println("pp", stdErr.String())
			log.Println("pp", err.Error())
			return nil, err
		}
	}
	{ // run pandoc
		outputFormat := "html5"
		args := []string{"-f", "markdown"}
		html := false
		switch path.Ext(name) {
		case "", ".html":
			html = true
		case ".pdf":
			outputFormat = "latex"
			args = append(args, "-o", path.Join(p.cache, name))
		}
		args = append(args, "-t", outputFormat)

		cmd := exec.CommandContext(context.Background(), "pandoc", args...)
		cmd.Stdin = buf
		stdOut := bytes.NewBuffer(nil)
		stdErr := bytes.NewBuffer(nil)
		cmd.Stdout = stdOut
		cmd.Stderr = stdErr

		if err := cmd.Run(); err != nil {
			log.Println("pandoc", stdErr.String())
			return nil, err
		}
		body := stdOut.String()
		if html {
			title := h1(bytes.NewBufferString(body))
			fx, err := os.Create(path.Join(p.cache, name))
			if err != nil {
				return nil, err
			}
			pageTemplate, err := template.New("").Parse(tmpl)
			containterName, _ := os.Hostname()
			instanceName, _ := metadata.InstanceID()
			err = pageTemplate.Execute(fx, map[string]interface{}{
				"Title":         title,
				"Name":          name,
				"Stripped":      strings.TrimSuffix(name, path.Ext(name)),
				"RenderTime":    time.Now().Sub(start),
				"RenderedAt":    time.Now(),
				"GoVersion":     runtime.Version(),
				"LastMod":       stt.ModTime(),
				"Container":     containterName,
				"Instance":      instanceName,
				"Body":          template.HTML(body),
				"PandocVersion": pandocVersion(),
			})
			if err != nil {
				return nil, err
			}
			fx.Close()
		}
	}
	fx, _ := os.Open(path.Join(p.cache, name))
	return fx, nil
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

var tmpl = `
<!DOCTYPE html>
<html>
<head>
  <title>{{ .Title }}</title>

  <meta charset="utf-8">
  <link rel="stylesheet" type="text/css" href="/saturn.css">
</head>
	<body>
		{{ .Body }}
		<hr/>
		<footer>
			<a href="javascript:history.back()">↩</a>
	 		<a href="#">⇪</a>
			<a href="{{ .Stripped }}.pdf">PDF</a>
			<details>
				render-time: {{ .RenderTime }}
				<br/>
				rendered-at: {{ .RenderedAt }}
				<br/>
 				last-modified: {{ .LastMod }}
				<br/>
				container:  {{ .Container }}
				<br/>
				go: {{ .GoVersion }}
				<br/>
				pandoc: {{ .PandocVersion }}
 			</details>
		</footer>
	</body>
</html>`

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
