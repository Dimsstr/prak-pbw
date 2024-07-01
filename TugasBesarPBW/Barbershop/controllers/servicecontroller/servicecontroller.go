package servicecontroller

import (
	"go-web-native/entities"
	"go-web-native/models/servicemodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	services := servicemodel.GetAll()

	data := map[string]interface{}{
		"services": services,
	}

	temp, err := template.ParseFiles("views/service/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/service/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var service entities.Service

		service.Name = r.FormValue("name")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			panic(err)
		}
		service.Price = price
		service.CreatedAt = time.Now()
		service.UpdatedAt = time.Now()

		ok := servicemodel.Create(service)
		if !ok {
			temp, _ := template.ParseFiles("views/service/create.html")
			temp.Execute(w, nil)
			return
		}

		http.Redirect(w, r, "/services", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		service := servicemodel.Detail(id)
		data := map[string]interface{}{
			"service": service,
		}

		temp, err := template.ParseFiles("views/service/edit.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var service entities.Service

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		service.Name = r.FormValue("name")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			panic(err)
		}
		service.Price = price
		service.UpdatedAt = time.Now()

		if ok := servicemodel.Update(id, service); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/services", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := servicemodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/services", http.StatusSeeOther)
}
