package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/LuD1161/go_programming_blueprints/trace"
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
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags
	http.Handle("/chat", MustAuth(&templateHandler{fileName: "chat.html"}))
	http.Handle("/login", &templateHandler{fileName: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/room", r)
	// get the room going
	go r.run()

	// start the webserver
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
