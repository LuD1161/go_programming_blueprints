package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templ represents a single template
type templateHandler struct {
	fileName string
	once     sync.Once          // to compile the template only once
	templ    *template.Template // compiled template reference
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.fileName)))
	})
	t.templ.Execute(w, r)
}

func main() {
	// root
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags
	http.Handle("/", &templateHandler{fileName: "chat.html"})

	r := newRoom()
	http.Handle("/room", r)
	// get the room going
	go r.run()

	// start the webserver
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
