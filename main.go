package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(
			// template.ParseFiles(filepath.Join("templates"), t.filename),
			template.ParseFiles(t.filename),
		)
	})
	t.templ.Execute(w, nil)
}

func Hoge() {
	fmt.Println("hoge")
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`
	// <html>
	// 	<head>
	// 		hi
	// 	</head>
	// </html>
	// 	`))
	// })

	if err := http.ListenAndServe(":18080", nil); err != nil {
		fmt.Println(err)
	}
}
