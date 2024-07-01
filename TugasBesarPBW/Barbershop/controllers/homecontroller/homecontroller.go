// controllers/homecontroller.go

package controllers

import (
	"html/template"
	"net/http"

	"go-web-native/models/capstermodel"
	"go-web-native/models/ordermodel"
	"go-web-native/models/servicemodel"
)

// Index function to handle GET request for home page.
func Index(w http.ResponseWriter, r *http.Request) {
	// Retrieve all services
	services := servicemodel.GetAll()

	// Retrieve all capsters
	capsters := capstermodel.GetAll()

	// Retrieve all orders
	orders := ordermodel.GetAll()

	// Prepare data to be passed to the template
	data := map[string]interface{}{
		"services": services,
		"capsters": capsters,
		"orders":   orders,
	}

	// Load your home/index.html template and execute it
	temp, err := template.ParseFiles("views/home/index.html") // Adjust the path according to your file structure
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
