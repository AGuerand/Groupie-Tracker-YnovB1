package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type PageData struct {
	Test []ArtistsData
}

type ArtistsData struct {
	image        string
	name         string
	menbers      string
	creationDate string
	locations    string
	dates        string
	relation     string
}

func main() {

	TooPrint := make([]ArtistsData, 0)

	TooPrint = append(TooPrint, ArtistsData{"je", "suis", "un", "test", "de", "projet", "x", "8"})
	TooPrint = append(TooPrint, ArtistsData{"je2", "suis2", "un2", "test2", "de2", "projet2", "x2", "8_2"})
	fmt.Println(TooPrint[0].relation, TooPrint[1].menbers)
	// m := map[string]Artists{
	// 	"cow": Artists{Famous_Artists{"walk"}, "grass", "moo"},
	// }
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Printf("Request failed %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		// for _, file := range string(data) {
		// 	z01.PrintRune(file)
		// }

		// var lookforHTTP = regexp.MustCompile(`^".+"$`)
		// fmt.Println(lookforHTTP.MatchString(string(data)))
	}

}

// func LoadData() {

// 	APIData := [4]string{
// 		"https://groupietrackers.herokuapp.com/api/artists",
// 		"https://groupietrackers.herokuapp.com/api/locations",
// 		"https://groupietrackers.herokuapp.com/api/dates",
// 		"https://groupietrackers.herokuapp.com/api/relation",
// 	}

// 	for _, link := range APIData {

// 		response, err := http.Get(link)
// 		if err != nil {
// 			fmt.Printf("Request failed %s\n", err)
// 		} else {
// 			data, _ := ioutil.ReadAll(response.Body)
// 			fmt.Println(string(data))

// 			// for _, file := range string(data) {
// 			// 	z01.PrintRune(file)
// 			// }

// 			var lookforHTTP = regexp.MustCompile(`^".+"$`)
// 			fmt.Println(lookforHTTP.MatchString(string(data)))

// 		}

// 	}

// }
