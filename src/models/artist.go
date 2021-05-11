package model

import "regexp"

func Get_ArtitsData(Data []string, Artists []ArtistsData) []ArtistsData {
	var regex_image = regexp.MustCompile(`https://.+.jpeg`)
	var regex_name = regexp.MustCompile(`"name":"(.+)","members`)
	var regex_members_list = regexp.MustCompile(`"members":\[(.+)\],"creationDate"`)
	var regex_creationDate = regexp.MustCompile(`\d{4}`)
	var regex_firstAlbum = regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)

	for _, element := range Data {
		Artists = append(Artists, ArtistsData{RelationData{},
			regex_image.FindString(element),           // + "," pour les images !!!!
			regex_name.FindStringSubmatch(element)[1], // + "," pour les noms !!!!
			getMembers(regex_members_list.FindStringSubmatch(element)[1]),
			regex_creationDate.FindString(element),
			regex_firstAlbum.FindString(element),
			"", ""})
	}

	return Artists
}

func getMembers(s string) []string {
	var regex_members = regexp.MustCompile(`([^,].+)`)
	last_index := 1
	result := make([]string, 0)
	for index := range s {
		if s[index] == ',' {
			result = append(result, regex_members.FindString(s[last_index:index-1])+",")
			last_index = index + 2
		} else if index == len(s)-1 {
			result = append(result, regex_members.FindString(s[last_index:len(s)-1])+",")
		}
	}
	return result
}
