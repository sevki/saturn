package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"path"

	"sevki.org/saturn/atlas"
	"sevki.org/saturn/pan"
	"sevki.org/saturn/titan"
	"upspin.io/config"
	"upspin.io/transports"
)

func main() {

	var (
		confName = flag.String("conf", "", "upspin-config")
		userName = flag.String("root-user", "", "upspin.User that namespace belongs to")
		prefix   = flag.String("prefix", "", "prefix to be added to the urls")
	)

	flag.Parse()

	cfg, err := config.FromFile(*confName)
	if err != nil {
		log.Fatal(err)
	}
	transports.Init(cfg)
	opts := []titan.Option{}
	if userName != nil {
		opts = append(opts, titan.WithRootUser(*userName))
	}
	if prefix != nil {
		opts = append(opts, titan.WithPrefix(*prefix))
	}

	ufs := titan.New(cfg, opts...)

	renderer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/_status":
			io.WriteString(w, "OK.")
			return
		}
		ext := path.Ext(r.URL.Path)
		fs := atlas.New(ufs)

		pan := http.FileServer(pan.New(fs))
		switch ext {
		case ".latex":
			w.Header().Set("Content-Type", "application/x-latex")
			pan.ServeHTTP(w, r)
		case ".pdf":
			w.Header().Set("Content-Type", "application/pdf")
			pan.ServeHTTP(w, r)
		case ".html":
			w.Header().Set("Content-Type", "text/html")
			pan.ServeHTTP(w, r)
		default:
			pan.ServeHTTP(w, r)
		}
	})

	log.Fatal(
		http.ListenAndServe(":8080", renderer),
	)
}
func redir(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/index"
		}
		h.ServeHTTP(w, r)
	})
}
