package model

import "regexp"

func Get_ArtitsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_image = regexp.MustCompile(`https://.+.jpeg`)
	var regex_name = regexp.MustCompile(`"name":"(.+)","members`)
	var regex_members = regexp.MustCompile(`"members":\[(.+)\],"creationDate"`)
	var regex_creationDate = regexp.MustCompile(`\d{4}`)
	var regex_firstAlbum = regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)

	for _, element := range Data {
		Artists = append(Artists, ArtistsData{RelationData{},
			regex_image.FindString(element),
			regex_name.FindStringSubmatch(element)[1],
			regex_members.FindStringSubmatch(element)[1],
			regex_creationDate.FindString(element),
			regex_firstAlbum.FindString(element),
			"", ""})
	}

	return Artists
}
