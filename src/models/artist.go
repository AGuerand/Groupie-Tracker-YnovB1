package model

import "regexp"

//recupere Artist Data => nom, images, membres, date de création et first album
func Get_ArtitsData(Data []string, Artists []ArtistsData) []ArtistsData { //Récupere :
	var regex_image = regexp.MustCompile(`https://.+.jpeg`)                          // image
	var regex_name = regexp.MustCompile(`"name":"(.+)","members`)                    // nom du groupe
	var regex_members_list = regexp.MustCompile(`"members":\[(.+)\],"creationDate"`) // membredu groupes
	var regex_creationDate = regexp.MustCompile(`\d{4}`)                             // date de création
	var regex_firstAlbum = regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)                   //first album date

	empty := make([]string, 0)

	for _, element := range Data {
		Artists = append(Artists, ArtistsData{empty, // Append a Artists:
			regex_image.FindString(element),                               // load image
			regex_name.FindStringSubmatch(element)[1] + ",",               // load name
			getMembers(regex_members_list.FindStringSubmatch(element)[1]), // load members
			regex_creationDate.FindString(element),                        // load creation date
			regex_firstAlbum.FindString(element),                          // load first album date
			"", ""})
	}

	return Artists
}

// recupere le nom des membres du groupe
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
