package ordercontroller

import (
	"go-web-native/entities"
	"go-web-native/models/capstermodel"
	"go-web-native/models/ordermodel"
	"go-web-native/models/servicemodel"
	"net/http"
	"strconv"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	orders := ordermodel.GetAll()
	data := map[string]interface{}{
		"orders": orders,
	}

	renderTemplate(w, "index", data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		capsters := capstermodel.GetAll()
		services := servicemodel.GetAll()

		data := map[string]interface{}{
			"capsters": capsters,
			"services": services,
		}

		renderTemplate(w, "create", data)
	} else if r.Method == "POST" {
		var order entities.Order

		order.Name = r.FormValue("name")
		capsterID, err := strconv.Atoi(r.FormValue("capster_id"))
		if err != nil {
			http.Error(w, "Invalid Capster ID", http.StatusBadRequest)
			return
		}
		order.CapsterID = capsterID

		serviceID, err := strconv.Atoi(r.FormValue("service_id"))
		if err != nil {
			http.Error(w, "Invalid Service ID", http.StatusBadRequest)
			return
		}
		order.ServiceID = serviceID

		order.Date = r.FormValue("date")
		order.Time = r.FormValue("time")
		order.Description = r.FormValue("description")

		ok := ordermodel.Create(order)
		if !ok {
			http.Error(w, "Failed to create order", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/orders", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idString := r.URL.Query().Get("id")
		if idString == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		order := ordermodel.Detail(id)
		if order.ID == 0 {
			http.NotFound(w, r)
			return
		}

		capsters := capstermodel.GetAll()
		services := servicemodel.GetAll()

		data := map[string]interface{}{
			"order":    order,
			"capsters": capsters,
			"services": services,
		}

		renderTemplate(w, "edit", data)
		return
	}

	if r.Method == "POST" {
		idString := r.FormValue("id")
		if idString == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var order entities.Order
		order.Name = r.FormValue("name")

		capsterID, err := strconv.Atoi(r.FormValue("capster_id"))
		if err != nil {
			http.Error(w, "Invalid Capster ID", http.StatusBadRequest)
			return
		}
		order.CapsterID = capsterID

		serviceID, err := strconv.Atoi(r.FormValue("service_id"))
		if err != nil {
			http.Error(w, "Invalid Service ID", http.StatusBadRequest)
			return
		}
		order.ServiceID = serviceID

		order.Date = r.FormValue("date")
		order.Time = r.FormValue("time")
		order.Description = r.FormValue("description")

		if ok := ordermodel.Update(id, order); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/orders", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := ordermodel.Delete(id); err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/orders", http.StatusSeeOther)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order := ordermodel.Detail(id)

	tmpl, err := template.ParseFiles("views/order/detail.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, order)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	tmplPath := "views/order/" + tmpl + ".html"
	temp, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

func OrderDetail(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order := ordermodel.Detail(id)

	tmpl, err := template.ParseFiles("views/order/detail.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Order entities.Order
	}{
		Order: order,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}
