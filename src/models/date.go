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

func Get_DatesData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_dates = regexp.MustCompile(`"dates":\[(.+)\]`)

	for index, element := range Data {
		Artists[index].dates = regex_dates.FindStringSubmatch(element)[1]
	}

	return Artists
}
