package capstercontroller

import (
	"go-web-native/entities"
	"go-web-native/models/capstermodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	capsters := capstermodel.GetAll()
	data := map[string]any{
		"capsters": capsters,
	}

	temp, err := template.ParseFiles("views/capster/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/capster/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var capster entities.Capster

		capster.Name = r.FormValue("name")
		capster.CreatedAt = time.Now()
		capster.UpdatedAt = time.Now()

		ok := capstermodel.Create(capster)
		if !ok {
			temp, _ := template.ParseFiles("views/capster/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/capsters", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/capster/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		capster := capstermodel.Detail(id)
		data := map[string]any{
			"capster": capster,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var capster entities.Capster

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		capster.Name = r.FormValue("name")
		capster.UpdatedAt = time.Now()

		if ok := capstermodel.Update(id, capster); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/capsters", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := capstermodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/capsters", http.StatusSeeOther)
}
