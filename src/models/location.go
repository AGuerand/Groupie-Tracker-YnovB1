package models

import "regexp"

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

func Get_LocationsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_locations = regexp.MustCompile(`"locations":\[(.+)\],"dates"`)

	for index, element := range Data {
		Artists[index].locations = regex_locations.FindStringSubmatch(element)[1]
	}

	return Artists
}
