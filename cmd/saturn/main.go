package main

import (
	"flag"
	"io"
	"log"
	"mime"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"sevki.org/saturn/pan"
)

var (
	root = flag.String("root", "/pub", "root")
)

func main() {

	flag.Parse()

	renderer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/_status":
			io.WriteString(w, "OK.")
			return
		}
		ext := path.Ext(r.URL.Path)

		pan := http.FileServer(
			pan.New(
				http.Dir(*root),
				pan.WithTemplate("latex", "/sevki.org/saturn/templates/default.latex"),
				pan.WithTemplate("html5", "/sevki.org/saturn/templates/web.html"),
			),
		)

		switch ext {
		case ".latex":
			w.Header().Set("Content-Type", "application/x-latex")
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
	http.Handle("/", mimeTypeHandler(renderer))

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}

func mimeTypeHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mimetype := "application/octet-stream"
		ext := path.Ext(r.URL.Path)
		if strings.HasSuffix(r.URL.Path, "/") {
			ext = ".html"
		}
		mt := mime.TypeByExtension(ext)
		if mt != "" {
			w.Header().Set("Content-Type", mt)
			h.ServeHTTP(w, r)
			return
		} else {
			rw := httptest.NewRecorder()
			h.ServeHTTP(rw, r)

			contentType := http.DetectContentType(rw.Body.Bytes())
			w.Header().Set("Content-Type", contentType)
			io.Copy(w, rw.Body)
			return
		}
		w.Header().Set("Content-Type", mimetype)
		h.ServeHTTP(w, r)
		return
	})
}

func sync() {
	f, err := os.Create("/pub/_log")
	if err != nil {
		log.Fatal(err)
	}
	for {

		cmd := exec.Command("gsutil",
			"rsync",
			"-d",
			"-r",
			"gs://sevki-io",
			*root,
		)

		cmd.Stdout = io.MultiWriter(f, os.Stdout)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Minute)
	}

}