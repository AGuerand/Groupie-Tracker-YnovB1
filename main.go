package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type PageData struct {
	var_Artists []ArtistsData
}

type RelationData struct {
	locations []string
	dates     []string
}

type ArtistsData struct {
	image        string
	name         string
	members      string
	creationDate string
	firstAlbum   string
	locations    string
	dates        string
	relation     RelationData
}

func main() {

	LoadData()

}

func LoadData() {
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
			fmt.Printf("Request failed %s\n", err)
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
	// for i := 0; i < 7; i++ {
	// 	fmt.Println(Artists[51].relation.locations[i], Artists[51].relation.dates[i])
	// }

}
func Get_ArtitsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_image = regexp.MustCompile(`https://.+.jpeg`)
	var regex_name = regexp.MustCompile(`"name":"(.+)","members`)
	var regex_members = regexp.MustCompile(`"members":\[(.+)\],"creationDate"`)
	var regex_creationDate = regexp.MustCompile(`\d{4}`)
	var regex_firstAlbum = regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)

	for _, element := range Data {

		Artists = append(Artists, ArtistsData{regex_image.FindString(element),
			regex_name.FindStringSubmatch(element)[1],
			regex_members.FindStringSubmatch(element)[1],
			regex_creationDate.FindString(element),
			regex_firstAlbum.FindString(element), "", "", RelationData{}})
	}
	return Artists
}
func Get_LocationsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"locations":\[(.+)\],"dates"`)

	for index, element := range Data {
		Artists[index].locations = regex_locations.FindStringSubmatch(element)[1]
	}
	return Artists
}
func Get_DatesData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_dates = regexp.MustCompile(`"dates":\[(.+)\]`)

	for index, element := range Data {
		Artists[index].dates = regex_dates.FindStringSubmatch(element)[1]
	}
	return Artists
}
func Get_RelationData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"([^\]][\D+]+)":\["\d{2}-\d{2}-\d{4}"`)
	var regex_dates = regexp.MustCompile(`"(\d{2}-\d{2}-\d{4})"`)
	LastIndex := 0
	locations := make([]string, 0)
	dates := make([]string, 0)
	datesToSlice := ""

	for Artists_index, element := range Data {
		for _, location := range regex_locations.FindAllStringSubmatch(element, -1) {
			locations = append(locations, location[1])
		}
		for index, indexFlag := range regex_locations.FindAllStringIndex(element, -1) {
			if index == 0 {
			} else if index == len(regex_locations.FindAllStringIndex(element, -1))-1 {
				for i, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {
					if i == 0 {
						datesToSlice += date[1]
					} else {
						datesToSlice += " " + date[1]
					}
				}
				dates = append(dates, datesToSlice)
				datesToSlice = ""
				for i, date := range regex_dates.FindAllStringSubmatch(element[indexFlag[0]:len(element)-1], -1) {
					if i == 0 {
						datesToSlice += date[1]
					} else {
						datesToSlice += " " + date[1]
					}
				}
				dates = append(dates, datesToSlice)
				datesToSlice = ""
			} else {
				for i, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {
					if i == 0 {
						datesToSlice += date[1]
					} else {
						datesToSlice += " " + date[1]
					}
				}
				dates = append(dates, datesToSlice)
				datesToSlice = ""
			}
			LastIndex = indexFlag[0]
		}
		Artists[Artists_index].relation = RelationData{locations, dates}
		LastIndex = 0
		locations = locations[:0]
		dates = dates[:0]
	}
	return Artists
}
