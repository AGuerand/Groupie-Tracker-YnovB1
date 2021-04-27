package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.HandleFunc("/", hom)
	//http.HandleFunc("/about", abo)
	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("assets"))))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func hom(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "default.html", Ascii)
	arg := r.FormValue("Arg")
	template := r.FormValue("Template")
	asciiLetters, err := mapping(template) // recover the map
	if err != nil {
		//Do Something
	}
	phrase := []rune(arg) // create a slice of rune who have every character
	Ascii(phrase, asciiLetters, w)
}

// func abo(w http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(w, "about.gohtml", nil)
// }
