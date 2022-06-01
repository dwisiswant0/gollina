package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"math/rand"
	"net/http"
	"text/template"

	"gopkg.in/loremipsum.v1"
)

func (opt *options) serve() {
	http.HandleFunc("/", opt.page)

	http.FileServer(http.FS(res))
	log.Printf("Server listening at %s:%d...\n", opt.host, opt.port)

	if err := http.ListenAndServe(fmt.Sprint(opt.host, ":", strconv.Itoa(opt.port)), nil); err != nil {
		log.Fatal(err)
	}
}

func (opt *options) page(w http.ResponseWriter, r *http.Request) {
	li := loremipsum.NewWithSeed(rand.Int63())

	page, ok := pages[r.URL.Path]
	if !ok {
		log.Printf(Err404, r.RequestURI, "")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFS(res, page)
	if err != nil {
		log.Printf(Err404, r.RequestURI, "in pages cache")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := map[string]interface{}{
		"lorem_ipsum": strings.Replace(li.Paragraphs(10), `\n`, `<br><br>`, -1),
		"PAYLOAD":     opt.payload,
	}

	if err := tpl.Execute(w, data); err != nil {
		return
	}
}
