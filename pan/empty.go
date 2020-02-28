package pan

import (
	"os"
	"time"
)

type rootdir struct {
}

type rootfileinfo struct {
}

func (r *rootdir) Close() error { return nil }

func (r *rootdir) Read(p []byte) (n int, err error) {
	panic("not implemented")
}

func (r *rootdir) Seek(offset int64, whence int) (int64, error) {
	panic("not implemented")
}

func (r *rootdir) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (r *rootdir) Stat() (os.FileInfo, error) {
	return &rootfileinfo{}, nil
}

func (i *rootfileinfo) Name() string {
	panic("not implemented")
}

func (i *rootfileinfo) Size() int64 {
	panic("not implemented")
}

func (i *rootfileinfo) Mode() os.FileMode {
	panic("not implemented")
}

func (i *rootfileinfo) ModTime() time.Time {
	panic("not implemented")
}

func (i *rootfileinfo) IsDir() bool { return true }

func (i *rootfileinfo) Sys() interface{} {
	panic("not implemented")
}
