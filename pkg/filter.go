package pkg

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func Filtering(creation_min string, creation_max string, album_min string, album_max string, members_min string, members_max string, location string) []models.Artists {
	var filtered []models.Artists

	for i := 0; i < len(Artist); i++ {
		check_location := false
		if location != "All" {
			for j := 0; j < len(Artist[i].Locations.Locations); j++ {
				if strings.Contains(Artist[i].Locations.Locations[j], strings.ToLower(location)) {
					check_location = true

					break
				}
			}
		} else {
			check_location = true
		}

		check_creation := false
		if strconv.Itoa(Artist[i].CreationDate) >= creation_min && strconv.Itoa(Artist[i].CreationDate) <= creation_max {
			check_creation = true
		}

		check_album := false
		if Artist[i].FirstAlbum[len(Artist[i].FirstAlbum)-4:] >= album_min && Artist[i].FirstAlbum[len(Artist[i].FirstAlbum)-4:] <= album_max {
			check_album = true
		}

		check_members := false
		if strconv.Itoa(len(Artist[i].Members)) >= members_min && strconv.Itoa(len(Artist[i].Members)) <= members_max {
			check_members = true
		}

		if check_location && check_creation && check_album && check_members {
			filtered = append(filtered, Artist[i])
		}
	}

	return filtered
}

func FindMinMax(a int, b int) bool {
	if a < b {
		return true
	}

	return false
}

func FindRange() error {
	var minCreationDate, maxCreationDate, minFirstAlbum, maxFirstAlbum, minMembers, maxMembers int

	for i := 0; i < len(Artist); i++ {
		firstAlbum, err := strconv.Atoi(Artist[i].FirstAlbum[len(Artist[i].FirstAlbum)-4:])
		if err != nil {
			return err
		}

		if i == 0 {
			minCreationDate = Artist[i].CreationDate
			maxCreationDate = Artist[i].CreationDate
			minFirstAlbum = firstAlbum
			maxFirstAlbum = firstAlbum
			minMembers = len(Artist[i].Members)
			maxMembers = len(Artist[i].Members)
		}

		if FindMinMax(maxCreationDate, Artist[i].CreationDate) {
			maxCreationDate = Artist[i].CreationDate
		}
		if !FindMinMax(minCreationDate, Artist[i].CreationDate) {
			minCreationDate = Artist[i].CreationDate
		}

		if FindMinMax(maxFirstAlbum, firstAlbum) {
			maxFirstAlbum = firstAlbum
		}
		if !FindMinMax(minFirstAlbum, firstAlbum) {
			minFirstAlbum = firstAlbum
		}

		if FindMinMax(maxMembers, len(Artist[i].Members)) {
			maxMembers = len(Artist[i].Members)
		}
		if !FindMinMax(minMembers, len(Artist[i].Members)) {
			minMembers = len(Artist[i].Members)
		}
	}

	return nil
}

func FindMAtch(finding string, array []string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == finding {
			return false
		}
	}

	return true
}

func FullListOfLocations() []string {
	arrayOfLocations := []string{}

	for i := 0; i < len(Artist); i++ {
		for j := 0; j < len(Artist[i].Locations.Locations); j++ {
			country := strings.Split(Artist[i].Locations.Locations[j], "-")

			if FindMAtch(country[1], arrayOfLocations) {
				arrayOfLocations = append(arrayOfLocations, country[1])
			}
		}
	}

	return arrayOfLocations
}
