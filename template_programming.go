package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Head    string
	Content string
	SairAdi string
}

func handler(w http.ResponseWriter, r *http.Request) {

	p := Page{
		Title:   "Regaip Kurt",
		Head:    "GoLang WEB Programlama Örnekleri",
		Content: "WEB Template ile Çalışma Örneği",
		SairAdi: "Orhan Veli",
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
