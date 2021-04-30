package model

import (
	"regexp"
)

func Get_RelationData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"([^\]][\D+]+)":\["\d{2}-\d{2}-\d{4}"`)
	var regex_dates = regexp.MustCompile(`"(\d{2}-\d{2}-\d{4})"`)
	LastIndex := 0
	locations := make([]string, 0)
	dates := make([]string, 0)
	datesToSlice := ""
	// result := RelationData{}

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
		Artists[Artists_index].Relation.Locations = locations
		Artists[Artists_index].Relation.Dates = dates
		// fmt.Println(locations, dates)
		LastIndex = 0
		locations = locations[:0]
		dates = dates[:0]
	}
	// for i := 0; i < 51; i++ {
	// 	fmt.Println(Artists[i].relation)
	// }
	return Artists
}
