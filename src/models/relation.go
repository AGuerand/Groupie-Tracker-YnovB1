package model

import (
	"regexp"
)

//get relation =>date + location
func Get_RelationData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"([^\]][\D+]+)":\["\d{2}-\d{2}-\d{4}"`)
	var regex_dates = regexp.MustCompile(`"(\d{2}-\d{2}-\d{4})"`)
	LastIndex := 0
	data := ""

	for Artists_index, element := range Data {
		for index, indexFlag := range regex_locations.FindAllStringIndex(element, -1) {
			// add space and "," for visibility
			if index == 0 {
			} else if index == len(regex_locations.FindAllStringIndex(element, -1))-1 {
				data += "," + regex_locations.FindStringSubmatch(element[LastIndex:indexFlag[0]])[1] + ","
				for _, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {

					data += " " + date[1]
				}

				data += "," + regex_locations.FindStringSubmatch(element[indexFlag[0] : len(element)-1])[1] + ","
				for _, date := range regex_dates.FindAllStringSubmatch(element[indexFlag[0]:len(element)-1], -1) {

					data += " " + date[1]
				}
			} else {
				data += "," + regex_locations.FindStringSubmatch(element[LastIndex:indexFlag[0]])[1] + ","
				for _, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {

					data += " " + date[1]
				}
			}
			LastIndex = indexFlag[0]
		}
		// add date and location in a array
		Artists[Artists_index].Relation = append(Artists[Artists_index].Relation, data)
		LastIndex = 0
		data = ""
	}

	return Artists
}
