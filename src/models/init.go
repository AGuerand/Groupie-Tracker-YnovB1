package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type RelationData struct {
	Locations []string
	Dates     []string
}

type ArtistsData struct {
	Relation     []string //RelationData
	Image        string
	Name         string
	Members      []string
	CreationDate string
	FirstAlbum   string
	Locations    string
	Dates        string
}

// Fonction principale appeler dans le main.go
func LoadData() []ArtistsData {
	Artists := make([]ArtistsData, 0)
	//Get data from API
	APIData := [4]string{
		"https://groupietrackers.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/locations",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relation",
	}
	var regex_bracket = regexp.MustCompile(`{([^{}]*)}`)

	for index, link := range APIData {
		response, err := http.Get(link)
		if err != nil {
			fmt.Printf("Reest failed %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			result := regex_bracket.FindAllString(string(data), -1)
			switch index {
			case 0:
				Artists = Get_ArtitsData(result, Artists) //recupere Artist Data => nom, images, membres, date de création et first album
			case 1:
				Artists = Get_LocationsData(result, Artists) // recupere les lieus des concerts
			case 2:
				Artists = Get_DatesData(result, Artists) //recupere les dates des concerts
			case 3:
				Artists = Get_RelationData(result, Artists) //recupere les relations =>date + location
			}
		}
	}

	return Artists
}
