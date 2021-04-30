package model

import "regexp"

func Get_LocationsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"locations":\[(.+)\],"dates"`)

	for index, element := range Data {
		Artists[index].locations = regex_locations.FindStringSubmatch(element)[1]
	}

	return Artists
}
