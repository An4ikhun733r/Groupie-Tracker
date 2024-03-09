package pkg

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func Searching(searchingText string) []models.Artists {
	if searchingText == "" {
		return Artist
	}

	searchingText = strings.ToLower(searchingText)
	var searched []models.Artists
	var check bool

	for i := 0; i < len(Artist); i++ {
		if strings.ToLower(Artist[i].Name) != strings.ReplaceAll(strings.ToLower(Artist[i].Name), searchingText, "") {
			searched = append(searched, Artist[i])

			continue
		}

		if strings.ToLower(Artist[i].FirstAlbum) != strings.ReplaceAll(strings.ToLower(Artist[i].FirstAlbum), searchingText, "") {
			searched = append(searched, Artist[i])

			continue
		}

		if strings.ToLower(strconv.Itoa(Artist[i].CreationDate)) != strings.ReplaceAll(strings.ToLower(strconv.Itoa(Artist[i].CreationDate)), searchingText, "") {
			searched = append(searched, Artist[i])

			continue
		}

		for k := 0; k < len(Artist[i].Members); k++ {
			check = true

			if strings.ToLower(Artist[i].Members[k]) != strings.ReplaceAll(strings.ToLower(Artist[i].Members[k]), searchingText, "") {
				searched = append(searched, Artist[i])

				check = false

				break
			}
		}

		if check {
			for j := 0; j < len(Artist[i].Locations.Locations); j++ {
				if strings.ToLower(Artist[i].Locations.Locations[j]) != strings.ReplaceAll(strings.ToLower(Artist[i].Locations.Locations[j]), searchingText, "") {
					searched = append(searched, Artist[i])

					break
				}
			}
		}
	}

	return searched
}

func AutoSearching(searchingText string) []string {
	searchingText = strings.ToLower(searchingText)
	searched := []string{}
	var check bool

	for i := 0; i < len(Artist); i++ {
		if strings.ToLower(Artist[i].Name) != strings.ReplaceAll(strings.ToLower(Artist[i].Name), searchingText, "") {
			searched = append(searched, Artist[i].Name)
		}

		if strings.ToLower(Artist[i].FirstAlbum) != strings.ReplaceAll(strings.ToLower(Artist[i].FirstAlbum), searchingText, "") {
			searched = append(searched, Artist[i].FirstAlbum+" - album")
		}

		if strings.ToLower(strconv.Itoa(Artist[i].CreationDate)) != strings.ReplaceAll(strings.ToLower(strconv.Itoa(Artist[i].CreationDate)), searchingText, "") {
			searched = append(searched, strconv.Itoa(Artist[i].CreationDate)+" - creation")
		}

		for k := 0; k < len(Artist[i].Members); k++ {
			check = true

			if strings.ToLower(Artist[i].Members[k]) != strings.ReplaceAll(strings.ToLower(Artist[i].Members[k]), searchingText, "") {
				searched = append(searched, Artist[i].Members[k]+" - member")

				check = false
			}
		}

		if check {
			for j := 0; j < len(Artist[i].Locations.Locations); j++ {
				if strings.ToLower(Artist[i].Locations.Locations[j]) != strings.ReplaceAll(strings.ToLower(Artist[i].Locations.Locations[j]), searchingText, "") {
					searched = append(searched, Artist[i].Locations.Locations[j]+" - location")
				}
			}
		}
	}

	return searched
}
