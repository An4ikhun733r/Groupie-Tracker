package web

import (
	"encoding/json"
	pkg "groupie-tracker/pkg"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)

		return
	}

	temp, err := template.ParseFiles("ui/templates/main.page.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}

	if pkg.Preloading {
		err = pkg.GetDataFromJSON(&pkg.Artist, "https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)

			return
		}

		err = pkg.BasicData()
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)

			return
		}

		pkg.Preloading = false
	}

	err = temp.Execute(w, pkg.Artist)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)

		return
	}

	id, err := strconv.Atoi((path.Base(r.URL.Path)))
	if err != nil || !pkg.IsValid(path.Base(r.URL.Path)) {
		ErrorHandler(w, r, http.StatusBadRequest)

		return
	}

	if !pkg.IsRange(id) {
		ErrorHandler(w, r, http.StatusNotFound)

		return
	}

	temp, err := template.ParseFiles("ui/templates/artist.page.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}

	err = temp.Execute(w, pkg.Artist[id-1])
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)

		return
	}

	searched := pkg.Searching(r.URL.Query().Get("searching"))

	temp, err := template.ParseFiles("ui/templates/main.page.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}

	err = temp.Execute(w, searched)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}
}

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)

		return
	}

	searched := pkg.Filtering(r.URL.Query().Get("creation_min"), r.URL.Query().Get("creation_max"), r.URL.Query().Get("album_min"), r.URL.Query().Get("album_max"), r.URL.Query().Get("members_min"), r.URL.Query().Get("members_max"), r.FormValue("location"))

	temp, err := template.ParseFiles("ui/templates/main.page.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}

	err = temp.Execute(w, searched)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}
}

func Autocomplete(w http.ResponseWriter, r *http.Request) {
	searched := pkg.AutoSearching(r.URL.Query().Get("search"))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searched)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	errMessage := ""

	switch status {
	case 400:
		errMessage = "400\nBad Request"
		w.WriteHeader(http.StatusNotFound)
	case 404:
		errMessage = "404\nPage Not Found"
		w.WriteHeader(http.StatusBadRequest)
	case 405:
		errMessage = "405\nMethod Not Allowed"
		w.WriteHeader(http.StatusMethodNotAllowed)
	case 500:
		errMessage = "500\nInternal Server Error"
		w.WriteHeader(http.StatusInternalServerError)
	}

	ts, err := template.ParseFiles("ui/templates/err.page.html")
	if err != nil {
		http.Error(w, "500\nInternal Server Error", 500)
		return
	}

	err = ts.Execute(w, errMessage)
	if err != nil {
		http.Error(w, "500\nInternal Server Error", 500)
		return
	}
}
