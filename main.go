package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	test    []string
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	Letest := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	json.NewEncoder(w).Encode(articles)

	data := Article{
		test: Letest,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", HandleRequest)
	http.ListenAndServe(":8081", nil)
}
