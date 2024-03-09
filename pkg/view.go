package pkg

import (
	"groupie-tracker/models"
	"strings"
)

func ChangingViewOfRelation(relations models.StructRelations) models.StructRelations {
	maprelations := relations.DatesLocations

	newMapRelations := make(map[string][]string)

	for key, value := range maprelations {
		arrayKey := strings.Split(key, "-")

		newMapRelations[ChangingViewOfLocations(arrayKey)] = ChangingViewOfMembers(value)
	}

	relations.DatesLocations = newMapRelations

	return relations
}

func ChangingViewOfLocations(locations []string) string {
	newKey := ""

	for i := 0; i < len(locations); i++ {
		newArrayKey := strings.Split(locations[i], "_")

		for j := 0; j < len(newArrayKey); j++ {
			if j+1 != len(newArrayKey) {
				newKey += strings.ToUpper(newArrayKey[j]) + " "
			} else {
				newKey += strings.ToUpper(newArrayKey[j])
			}
		}

		if i != len(locations)-1 {
			newKey += ", "
		}
	}

	return newKey
}

func ChangingViewOfMembers(members []string) []string {
	for i := 0; i < len(members); i++ {
		if i+1 != len(members) {
			members[i] += ","
		}
	}

	return members
}
