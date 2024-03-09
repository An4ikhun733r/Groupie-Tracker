package pkg

import (
	"encoding/json"
	"groupie-tracker/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	Artist     []models.Artists
	Preloading = true
)

func GetDataFromJSON(data interface{}, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

func BasicData() error {
	for i := 0; i < len(Artist); i++ {
		err := AdditionalData(i + 1)
		if err != nil {
			return err
		}
	}

	return nil
}

func AdditionalData(id int) error {
	location := models.StructLocations{}
	err := GetDataFromJSON(&location, "https://groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(id))
	if err != nil {
		return err
	}
	Artist[id-1].Locations = location

	date := models.StructConcertDates{}
	err = GetDataFromJSON(&date, "https://groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(id))
	if err != nil {
		return err
	}
	Artist[id-1].ConcertDates = date

	relation := models.StructRelations{}
	err = GetDataFromJSON(&relation, "https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id))
	if err != nil {
		return err
	}
	Artist[id-1].Relations = relation

	// Artist[id-1].Relations = ChangingViewOfRelation(relation)
	// Artist[id-1].Members = ChangingViewOfMembers(Artist[id-1].Members)

	return nil
}
