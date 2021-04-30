package model

import "regexp"

func Get_DatesData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_dates = regexp.MustCompile(`"dates":\[(.+)\]`)

	for index, element := range Data {
		Artists[index].dates = regex_dates.FindStringSubmatch(element)[1]
	}

	return Artists
}
