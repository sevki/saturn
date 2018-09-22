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

	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"willnorris.com/go/newbase60"
)

var (
	l = log.New(os.Stdout, "pan: ", 0)
)

type pan struct {
	fs    http.FileSystem
	cache string

	templates map[string]string
}

// New returns a new TitanFS
func New(fs http.FileSystem, options ...Option) http.FileSystem {
	dir, _ := ioutil.TempDir("", "pan")
	p := &pan{
		fs:        fs,
		cache:     dir,
		templates: make(map[string]string),
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
	const (
		none = ""
		html = ".html"
		pdf  = ".pdf"
	)

	ext := path.Ext(name)
	namePlain := strings.TrimSuffix(name, ext)

	switch ext {
	case none:
	case pdf, html:
		mdName := fmt.Sprintf("%s.%s", namePlain, "md")
		f, err := p.fs.Open(mdName)
		if err == nil {
			_, fileName := path.Split(name)
			return p.render(fileName, f)
		}
	}

	f, err := p.fs.Open(name)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func hash(s string) string {
	h := fnv.New32a()
	io.WriteString(h, s)
	a := int(h.Sum32())
	return newbase60.EncodeInt(a)
}

func (p *pan) render(name string, f http.File) (http.File, error) {
	stt, _ := f.Stat()
	if stt.IsDir() {
		return f, nil
	}

	bytz, err := ioutil.ReadAll(f)
	if err != nil {
		l.Println( errors.Wrap(err, "titan read all"))

		return nil, errors.Wrap(err, "titan read all")
	}
	buf := bytes.NewBuffer(nil)

	cache := path.Join("/x/", hash(name)+"-"+stt.ModTime().Format("20060102150405"))
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
			l.Println("pp", stdErr.String())
			l.Println("pp", err.Error())
			return nil, err
		}
	}
	{ // run pandoc
		outputFormat := "html5"
		permalink, _ := path.Split(stt.Name())
		args := []string{"-f", "markdown", "-V", fmt.Sprintf("permalink:%s", permalink)}
		switch path.Ext(name) {
		case "", ".html":

			args = append(args, "-o", path.Join(cache, name))
		case ".pdf":
			outputFormat = "latex"
			args = append(args, "-o", path.Join(cache, name))
		}
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

		if err := cmd.Run(); err != nil {
			l.Println("pandoc", stdErr.String())
			return nil, err
		}

	}
	fx, _ := os.Open(path.Join(cache, name))
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
