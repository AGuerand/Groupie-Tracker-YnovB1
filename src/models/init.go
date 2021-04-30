package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type RelationData struct {
	locations []string
	dates     []string
}

type ArtistsData struct {
	relation     RelationData
	image        string
	name         string
	members      string
	creationDate string
	firstAlbum   string
	locations    string
	dates        string
}

func LoadData() []ArtistsData {
	Artists := make([]ArtistsData, 0)
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
				Artists = Get_ArtitsData(result, Artists)
			case 1:
				Artists = Get_LocationsData(result, Artists)
			case 2:
				Artists = Get_DatesData(result, Artists)
			case 3:
				Artists = Get_RelationData(result, Artists)
			}
		}
	}

	return Artists
}
