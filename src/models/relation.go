package model

import (
	"regexp"
)

//recupere les relations =>date + location
func Get_RelationData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"([^\]][\D+]+)":\["\d{2}-\d{2}-\d{4}"`)
	var regex_dates = regexp.MustCompile(`"(\d{2}-\d{2}-\d{4})"`)
	LastIndex := 0
	// locations := make([]string, 0)
	// dates := make([]string, 0)
	data := ""

	// result := RelationData{}

	for Artists_index, element := range Data {
		// for _, location := range regex_locations.FindAllStringSubmatch(element, -1) {
		// 	locations = append(locations, location[1])
		// }
		for index, indexFlag := range regex_locations.FindAllStringIndex(element, -1) {

			if index == 0 {
			} else if index == len(regex_locations.FindAllStringIndex(element, -1))-1 {
				data += "," + regex_locations.FindStringSubmatch(element[LastIndex:indexFlag[0]])[1] + ","
				for _, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {

					data += " " + date[1]
				}
				// dates = append(dates, datesToSlice)
				data += "," + regex_locations.FindStringSubmatch(element[indexFlag[0] : len(element)-1])[1] + ","
				for _, date := range regex_dates.FindAllStringSubmatch(element[indexFlag[0]:len(element)-1], -1) {

					data += " " + date[1]
				}
				// dates = append(dates, datesToSlice)
				// datesToSlice = ""
			} else {
				data += "," + regex_locations.FindStringSubmatch(element[LastIndex:indexFlag[0]])[1] + ","
				for _, date := range regex_dates.FindAllStringSubmatch(element[LastIndex:indexFlag[0]], -1) {

					data += " " + date[1]
				}
				// dates = append(dates, datesToSlice)
			}
			LastIndex = indexFlag[0]
		}

		Artists[Artists_index].Relation = append(Artists[Artists_index].Relation, data)
		// fmt.Println(data)
		// fmt.Println(locations, dates) // valide
		LastIndex = 0
		data = ""
		// locations = locations[:0]
		// dates = dates[:0]
		// fmt.Println(Artists[Artists_index].Relation) // valide ???
	}
	// for i := range Artists {
	// 	fmt.Printf("%v\n", Artists[i].Relation)
	// 	fmt.Println()
	// }
	return Artists
}
