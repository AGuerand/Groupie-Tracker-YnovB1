package model

import "regexp"

// get concert location
func Get_LocationsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"locations":\[(.+)\],"dates"`)

	for index, element := range Data {
		Artists[index].Locations = regex_locations.FindStringSubmatch(element)[1]
	}

	return Artists
}
