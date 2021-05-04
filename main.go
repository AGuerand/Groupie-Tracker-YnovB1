package main

import (
	model "Groupie-Tracker/src/models"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", hom)  
	http.HandleFunc("/test", test)                                                                 // the func redirect to the hom function when we go to localhost:6082/
	http.HandleFunc("/Select", choose)                                                            // the func redirect to the read function when	we go to localhost:6082/read
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) // we handle /assets/ in our localhost
	fmt.Printf("Starting server at port 6085\n")
	if err := http.ListenAndServe(":6085", nil); err != nil { // we open the serve and we have and we have an error handling
		log.Fatal(err)
	}
}

type PageData struct {
	Image []string
}

func hom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOM")
	// fmt.Println("Chargement des struct : (200) ")
	Artists := model.LoadData()
	// i:= 0
	// for index, data := range Artists {
		// 	fmt.Println("Artist n:", index+1, " Nom:", data.Name, " Membre:", data.Members)
	// 	fmt.Println("Date de création:", data.CreationDate, " Date du premier album:", data.FirstAlbum)
	// 	fmt.Println("image du groupe:", data.Image)
	// 	fmt.Println("Lieu des concerts:", data.Locations)
	// 	fmt.Println("Dates des concerts:", data.Dates)
	// 	fmt.Println()
	// }
	TooPrint := make([]string,0)
	tpl := template.Must(template.ParseFiles("mygptrack/index.html"))

	for i := range Artists{
		
		TooPrint = append(TooPrint,Artists[i].Image)
	}
	
	
	data := PageData{
		Image: TooPrint,
	}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}

func choose(w http.ResponseWriter, r *http.Request) {
	Artists := model.LoadData()
	i:= 0
	test := Artists[i].Image
	fmt.Println("zdfguilqszdfuifazGZZFGUZUK")
	tpl := template.Must(template.ParseFiles("mygptrack/Select.html"))

	err := tpl.Execute(w, test)
	if err != nil {
		log.Fatal(err)
	}

}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOM")
	// fmt.Println("Chargement des struct : (200) ")
	Artists := model.LoadData()
	// i:= 0
	// for index, data := range Artists {
		// 	fmt.Println("Artist n:", index+1, " Nom:", data.Name, " Membre:", data.Members)
	// 	fmt.Println("Date de création:", data.CreationDate, " Date du premier album:", data.FirstAlbum)
	// 	fmt.Println("image du groupe:", data.Image)
	// 	fmt.Println("Lieu des concerts:", data.Locations)
	// 	fmt.Println("Dates des concerts:", data.Dates)
	// 	fmt.Println()
	// }
	TooPrint := make([]string,0)
	tpl := template.Must(template.ParseFiles("mygptrack/test.html"))

	for i := range Artists{
		
		TooPrint = append(TooPrint,Artists[i].Image)
	}
	
	
	data := PageData{
		Image: TooPrint,
	}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
	}
