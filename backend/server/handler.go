package server

import (
	backend "groupie-tracker/backend/database"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type ErrorPage struct {
	StatusCode int
	StatusText string
}

func errHandler(w http.ResponseWriter, r *http.Request, statusCode int, statusText string) {
	data := ErrorPage{
		StatusCode: statusCode,
		StatusText: statusText,
	}
	t, err := template.ParseFiles("./frontend/html/errpage.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.WriteHeader(statusCode)
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method != http.MethodGet {
		errHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	t, err := template.ParseFiles("frontend/html/index.html")
	if err != nil {
		log.Println(err)
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = backend.GetData(&backend.Artist, "https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = t.Execute(w, backend.Artist)
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if r.URL.Path != "/artist/" {
		errHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	id := r.URL.Query().Get("id")

	if !backend.IsValid(id) {
		errHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	idd, err := strconv.Atoi(id)

	err = backend.GetData(&backend.Artist, "https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if !backend.IsRange(idd) {
		errHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	err = backend.AdditionalData(idd)
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	t, err := template.ParseFiles("frontend/html/artist.html")
	if err != nil {
		log.Println(err)
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = t.Execute(w, backend.Artist[idd-1])
	if err != nil {
		http.Error(w, "Error executin file", http.StatusInternalServerError)
		return
	}
}

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))

		return
	}

	searched := backend.Filtering(r.URL.Query().Get("creation_min"), r.URL.Query().Get("creation_max"), r.URL.Query().Get("album_min"), r.URL.Query().Get("album_max"), r.URL.Query().Get("members_min"), r.FormValue("location"))

	temp, err := template.ParseFiles("frontend/html/index.html")
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))

		return
	}

	err = temp.Execute(w, searched)
	if err != nil {
		errHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))

		return
	}
}