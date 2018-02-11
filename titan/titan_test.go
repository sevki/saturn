package titan

import (
	"testing"

	"sevki.org/x/pretty"
	"upspin.io/config"
	"upspin.io/transports"
)

func TestFile(t *testing.T) {
	cfg, err := config.FromFile("/home/sevki/upspin/web-config")
	transports.Init(cfg)
	fs := New(cfg, WithUser("s@sevki.io"))
	x, err := fs.Open("web/hello")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	stat, err := x.Stat()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Logf(pretty.JSON(stat.Name()))
}
func TestDir(t *testing.T) {
	cfg, err := config.FromFile("/home/sevki/upspin/web-config")
	transports.Init(cfg)
	fs := New(cfg, WithUser("s@sevki.io"))
	x, err := fs.Open("web/")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	files, err := x.Readdir(-1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Logf(pretty.JSON(files))
}