package main

import (
	model "Groupie-Tracker/src/models"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", hom)                                                                 // the func redirect to the hom function when we go to localhost:6082/
	// http.HandleFunc("/read", read)                                                            // the func redirect to the read function when	we go to localhost:6082/read
	http.Handle("../assets/", http.StripPrefix("../assets", http.FileServer(http.Dir("assets")))) // we handle /assets/ in our localhost
	fmt.Printf("Starting server at port 6082\n")
	if err := http.ListenAndServe(":6082", nil); err != nil { // we open the serve and we have and we have an error handling
		log.Fatal(err)
	}
}

func hom(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Chargement des struct : (200) ")
	Artists := model.LoadData()
	for index, data := range Artists {
		fmt.Println("Artist n:", index+1, " Nom:", data.Name, " Membre:", data.Members)
		fmt.Println("Date de création:", data.CreationDate, " Date du premier album:", data.FirstAlbum)
		fmt.Println("image du groupe:", data.Image)
		fmt.Println("Lieu des concerts:", data.Locations)
		fmt.Println("Dates des concerts:", data.Dates)
		fmt.Println()
	}

	tpl := template.Must(template.ParseFiles("../mygptrack/index.html"))
	// asciiLetters, err := mapping(data.Template) // recover the map
	// if err != nil {
	// 	//Do Something
	// }
	// phrase := []rune(data.Arg) // create a slice of rune who have every character

	// if data.Format != "" {
	// 	err = os.Remove("assets/ascii." + data.Format)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("Removing assets/ascii." + data.Format)

	// 	file, err := os.OpenFile("assets/ascii."+data.Format, os.O_CREATE|os.O_WRONLY, 0666) //create a file with a name and permissions
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("Opening the file assets/ascii." + data.Format)
	// 	AsciiInFile(phrase, asciiLetters, file)
	// 	file.Close()
	// }

	// data.Res = Ascii(phrase, asciiLetters, w)
	err := tpl.Execute(w, Artists)
	if err != nil {
		log.Fatal(err)
	}

}
