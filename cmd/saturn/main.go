package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"sevki.org/saturn/pan"
	"sevki.org/saturn/titan"
	"upspin.io/config"
	"upspin.io/transports"
)

func main() {

	var (
		confName = flag.String("conf", "/root/upspin/config", "upspin-config")
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
		fs := ufs
		pan := http.FileServer(
			pan.New(fs,
				pan.WithTemplate("latex", "/go/src/sevki.org/saturn/templates/default.latex"),
				pan.WithTemplate("html5", "/go/src/sevki.org/saturn/templates/web.html"),
			),
		)
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Exit(0)
	}()

	http.Handle("/x/", http.StripPrefix("/x/", http.FileServer(http.Dir("/x"))))
	http.Handle("/", renderer)

	log.Fatal(
		http.ListenAndServe(":8080", nil),
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
