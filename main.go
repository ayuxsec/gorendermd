package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"strconv"
)

// PageData holds data to be passed to the HTML template
type PageData struct {
	MarkDown string
}

func main() {
	if err := parseArgs(); err != nil {
		log.Fatalf("failed to parse args: %v", err)
	}
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/index.html"))
	fs := http.FileServer(http.FS(staticFS))

	http.Handle("/static/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			MarkDown: rawMD,
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			//return
		}
	})

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	tcpAddr := ln.Addr().(*net.TCPAddr)

	listenerURL := "http://localhost:" + strconv.Itoa(tcpAddr.Port)
	log.Printf("Listening on %s", listenerURL)

	if *browser != "" {
		if !*skipOpenBrowser {
			if err := runCommand(*browser, []string{listenerURL}, *browserTimeout); err != nil {
				log.Fatalf("failed to open browser: %v", err)
			}
		}
	}
	log.Fatal(http.Serve(ln, nil))
}
