package atlas

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	l = log.New(os.Stdout, "atlas: ", log.Lshortfile)
)

func New(fs http.FileSystem) http.FileSystem {
	return &atlasFS{fs: fs}
}

type atlasFS struct {
	fs http.FileSystem
}

func (a *atlasFS) Open(name string) (http.File, error) {
	return tryFiles(name, a.fs)
}

// GuessPossibleNames for a given URL guesses file names
func guessPossibleNames(name string) []string {
	ext := path.Ext(name)
	namePlain := strings.TrimSuffix(name, ext)

	guesses := []string{}
	switch ext {
	case "":
		guesses = append(guesses, fmt.Sprintf("%s.%s", namePlain, "md"))

		guesses = append(guesses, path.Join(namePlain, "index.html"))
		guesses = append(guesses, path.Join(namePlain, "index.md"))
	case ".pdf", ".html":
		guesses = append(guesses, fmt.Sprintf("%s.%s", namePlain, "md"))
		guesses = append(guesses, fmt.Sprintf("%s.%s", namePlain, "html"))
	default:
		return []string{name}
	}

	return guesses
}

type emptyDir struct{}

func (*emptyDir) Close() error                                 { return nil }
func (*emptyDir) Read([]byte) (int, error)                     { return 0, io.EOF }
func (*emptyDir) Readdir(count int) ([]os.FileInfo, error)     { return nil, io.EOF }
func (*emptyDir) Seek(offset int64, whence int) (int64, error) { return 0, nil }
func (*emptyDir) Stat() (os.FileInfo, error)                   { return &dirInfo{}, nil }

type dirInfo struct{}

func (*dirInfo) Name() string       { return "/" }
func (*dirInfo) Size() int64        { return 0 }
func (*dirInfo) Mode() os.FileMode  { return os.ModeDir | 0555 }
func (*dirInfo) ModTime() time.Time { return time.Now() }
func (*dirInfo) IsDir() bool        { return true }
func (*dirInfo) Sys() interface{}   { return nil }

// TryFiles guesses a bunch of file by their names
func tryFiles(name string, fs http.FileSystem) (http.File, error) {
	var f http.File
	if stat, err := os.Stat(name); err == nil && stat != nil {
		return os.Open(name)
	}
	if name == "/" {
		return &emptyDir{}, nil
	}
	for _, guess := range guessPossibleNames(name) {
		var err error
		rendered := path.Ext(guess) == ".md"
		f, err = fs.Open(guess)
		if err == nil {
			s, err := f.Stat()
			if err == nil && s.IsDir() {
				rendered = false
			}

			l.Printf("MATCH name: %s guess: %s rendered: %v isDir: %v\n", s.Name(), guess, rendered, s.IsDir())
			return &atlasFile{
				name:       name,
				underlying: f,
				sys: SysInfo{
					ShouldRender:  rendered,
					CanonicalName: guess,
					Name:          name,
				},
			}, nil
		}
		l.Printf("FAIL name: %s guess: %s rendered: %v\n", name, guess, rendered)
	}
	return nil, os.ErrNotExist
}

type SysInfo struct {
	ShouldRender  bool
	CanonicalName string
	Name          string
}

type atlasFile struct {
	name       string
	underlying http.File
	sys        SysInfo
}

func (atlasFile) Close() error                               { return nil }
func (c atlasFile) Read(p []byte) (int, error)               { return c.underlying.Read(p) }
func (c atlasFile) Seek(a int64, b int) (int64, error)       { return c.Seek(a, b) }
func (c atlasFile) Readdir(count int) ([]os.FileInfo, error) { return c.underlying.Readdir(count) }
func (c atlasFile) Stat() (os.FileInfo, error) {
	stat, err := c.underlying.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "atlasFile stat")
	}

	return &atlasFileInfo{
		size:    stat.Size(),
		mode:    stat.Mode(),
		sys:     c.sys,
		name:    c.name,
		dir:     stat.IsDir(),
		modTime: stat.ModTime(),
	}, nil
}

type atlasFileInfo struct {
	size    int64
	mode    os.FileMode
	sys     interface{}
	name    string
	modTime time.Time
	dir     bool
}

func (t *atlasFileInfo) Name() string       { return t.name }
func (t *atlasFileInfo) Size() int64        { return t.size }
func (t *atlasFileInfo) Mode() os.FileMode  { return t.mode }
func (t *atlasFileInfo) ModTime() time.Time { return t.modTime }
func (t *atlasFileInfo) IsDir() bool        { return t.dir }
func (t *atlasFileInfo) Sys() interface{}   { return t.sys }
