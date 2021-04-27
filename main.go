package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", hom)
	//http.HandleFunc("/about", abo)
	// http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("assets"))))

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
	tpl.ExecuteTemplate(w, "index.html", Ascii)
}

// func abo(w http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(w, "about.gohtml", nil)
// }
