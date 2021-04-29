package main

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

func main() {

	LoadData()
}

func LoadData() {
	Artists := make([]ArtistsData, 0)
	APIData := [4]string{
		"https://groupietracrs.herokuapp.com/api/artists",
		"https://groupietrackers.herokuapp.com/api/location",
		"https://groupietrackers.herokuapp.com/api/dates",
		"https://groupietrackers.herokuapp.com/api/relatio",
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
				// Artists = Get_RelationData(result, Artists)
				fmt.Println(Artists[0], Artists[51])
			}
		}
	}

}

// func Get_ArtitsData(Data []string, Artists []ArtistsData) []ArtistsData {
// // // 	var regex_image = regexp.MustCompile(`https://.+.jpeg`)
// // // 	var regex_name = regexp.MustCompile(`"name":"(.+)","memrs`)
// // // 	var regex_members = regexp.MustCompile(`"members":\[(.+)\],"eationDate"`)
// // // 	var regex_creationDate = regexp.MustCompile(`\d{4}`)
// // // 	var regex_firstAlbum = regexp.MustCompile(`\d{2}-\d{-\d{4}`)

// 	for _, element := range Data {

// 	Artists = append(Artists, ArtistsData{RelationData{},
// // 			regex_image.FindString(element),
// // 			regex_name.FindStringSubmatch(elent)[1],
// // 			regex_members.FindStringSubmatch(element)],
// // 			regex_creationDate.FindString(element),
// // 			regex_firstAlbum.FindString(element),
// // 			"", ""})
// // 	}
// // 	rurn Artists
// // }

// funcGet_LocationsData(Data []string, Artists []ArtistsData) []ArtistsData {
// 	var regex_locations = regexp.MustCompile(`"locations":\[(.+)\],"dates"`)

// for index, element := range Data {
// 		Artists[index].locations = regexocations.FindStringSubmatch(element)[1]
// 	}
// 	rurn Artists
// }
// fc Get_DatesData(Data []string, Artists []ArtistsData) []ArtistsData {
// 	var regex_dates = regexp.MustCompile(`"dates":\[(.+)\]`)

// for index, element := range Data {
// 		Artists[index].dates = regex_dat.FindStringSubmatch(element)[1]
// 	}
// 	rurn Artists
// }

// func Get_RelationData(Data []string, Artists []ArtistsData) []ArtistsData {
// 	var regex_locations = regexp.MustCompile(`"([^\]][\D+]+)":\["\d{2}-\d{2}-{4}"`)
// 	var regex_dates = regexp.MustCompile(`"(\d{2}-\d{2}-\d{4})"`)
// 	LastIndex := 0
// 	locations := me([]string, 0)
// 	dates := make([]string, 0)
// 	datesToSlice := ""
// 	result := Relationta{}

// for Artists_index, element := range Data {
// 		for _, location := range regex_locationsindAllStringSubmatch(element, -1) {
// 			locations = append(locations, location[1])
// 		}
// 		f index, indexFlag := range regex_locations.FindAllStringIndex(element, -1) {
// 			if index == 0 {
// 			} else if index= len(regex_locations.FindAllStringIndex(element, -1))-1 {
// 				for i, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:dexFlag[0]], -1) {
// 					if i == 0 {
// 						datesToSli += date[1]
// 					} else {
// 						datesToice += " " + date[1]
// 					}
// 				}
// 				des = append(dates, datesToSlice)
// 				datesToSlice = ""
// 				for i, date := rae regex_dates.FindAllStringSubmatch(element[indexFlag[0]:len(element)-1], -1) {
// 					if i == 0 {
// 						datesToSli += date[1]
// 					} else {
// 						datesToice += " " + date[1]
// 					}
// 				}
// 				des = append(dates, datesToSlice)
// 				datesToSlice = ""
// 			} else {
// 				for i, te := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {
// 					if i == 0 {
// 						datesToSli += date[1]
// 					} else {
// 						datesToice += " " + date[1]
// 					}
// 				}
// 				des = append(dates, datesToSlice)
// 				datesToSlice = ""
// 			}
// 			LtIndex = indexFlag[0]
// 		}
// 		rult = RelationData{locations, dates}
// 		Artists[Artists_index].relation = resu
// 		LastIndex = 0
// 		locations = lations[:0]
// 		dates = dates[:0]
// 	}
// 	rurn Artists
// }
