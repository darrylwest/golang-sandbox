package main

import (
	"../trace"
	"./room"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("chat/templates", t.filename)))
	})

	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "the application address")
	flag.Parse()

	r := NewRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	log.Println("start the server in background: version 00.90.14 on ", *addr)
	go r.run()

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen and serve error: ", err)
	} else {
		log.Println("running on: ", *addr)
	}
}
