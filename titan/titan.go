package titan

import (
	"log"
	"net/http"
	"os"
	fpath "path/filepath"
	"time"

	"github.com/pkg/errors"
	"upspin.io/client"
	ue "upspin.io/errors"
	"upspin.io/path"
	"upspin.io/upspin"
)

var (
	l = log.New(os.Stdout, "titan: ", 0)
)

type titan struct {
	upspin.Client

	dir upspin.DirServer

	prefix   string
	userName string
	cfg      upspin.Config
}

type Option func(*titan)

func WithDirServer(dir upspin.DirServer) Option {
	return func(t *titan) { dir = dir }
}

func WithRootUser(name string) Option {
	return func(t *titan) { t.userName = name }
}

func WithPrefix(prefix string) Option {
	return func(t *titan) { t.prefix = prefix }
}

func New(cfg upspin.Config, options ...Option) http.FileSystem {
	x := &titan{userName: string(cfg.UserName()), cfg: cfg}
	for _, opt := range options {
		opt(x)
	}
	x.Client = client.New(x.cfg)
	return x
}

func (t *titan) Open(name string) (http.File, error) {
	upPath := upspin.PathName(fpath.Join(t.userName, t.prefix, name))
	l.Println(name, upPath)

	f, err := t.Client.Open(upPath)
	if err != nil {
		if ue.Is(ue.IsDir, err) {
			return &titanFile{nil, t.Client, upPath, t.prefix}, nil
		} else if ue.Is(ue.NotExist, err) {
			return nil, os.ErrNotExist
		}
		return nil, os.ErrNotExist
	}
	return &titanFile{f, t.Client, f.Name(), t.prefix}, nil
}

type titanFile struct {
	upspin.File

	c      upspin.Client
	upName upspin.PathName
	prefix string
}

func (tf *titanFile) Name() string {
	p, _ := path.Parse(tf.upName)
	return p.FilePath()
}
func (tf *titanFile) Close() error {
	return nil
}
func (tf *titanFile) Readdir(count int) ([]os.FileInfo, error) {
	stat, err := tf.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "tf.readdir.stat")
	}
	if !stat.IsDir() {
		return nil, os.ErrInvalid
	}
	glob := upspin.AllFilesGlob(tf.upName)
	dir, err := tf.c.DirServer(upspin.PathName(glob))
	if err != nil {
		return nil, errors.Wrap(err, "tf.Readdir.DirServer")
	}
	des, err := dir.Glob(glob)
	if err != nil {
		return nil, errors.Wrap(err, "tf.Stat.Lookup")
	}
	var infos []os.FileInfo
	for _, de := range des {
		infos = append(infos, titanFileInfo{*de, fpath.Join(tf.Name(), "")})
		if len(infos) == count {
			break
		}
	}
	return infos, nil
}

func (tf *titanFile) Stat() (os.FileInfo, error) {
	dir, err := tf.c.DirServer(tf.upName)
	if err != nil {
		return nil, errors.Wrap(err, "tf.Stat.DirServer")
	}
	de, err := dir.Lookup(tf.upName)
	if err != nil {
		return nil, errors.Wrap(err, "tf.Stat.Lookup")
	}
	return titanFileInfo{*de, tf.prefix}, nil
}

type titanFileInfo struct {
	upspin.DirEntry

	prefix string
}

func (tfi titanFileInfo) Name() string {
	parsed, _ := path.Parse(tfi.DirEntry.Name)
	fullpath := string(parsed.FilePath())

	prefixParsed, _ := path.Parse(upspin.PathName(tfi.prefix))

	if parsed.HasPrefix(prefixParsed) {
		fullpath = fullpath[len(tfi.prefix)+1:]
	}
	return fullpath
}

func (tfi titanFileInfo) Size() int64 {
	size, err := tfi.DirEntry.Size()
	if err != nil {
		return -1
	}
	return size
}

func (tfi titanFileInfo) Mode() os.FileMode {
	de := tfi.DirEntry
	if de.IsDir() {
		return os.ModeDir
	}
	return os.FileMode(0777)
}

func (tfi titanFileInfo) ModTime() time.Time {
	return tfi.Time.Go()
}

func (tfi titanFileInfo) IsDir() bool {
	return tfi.DirEntry.IsDir()
}

func (tfi titanFileInfo) Sys() interface{} {
	return tfi.DirEntry
}
