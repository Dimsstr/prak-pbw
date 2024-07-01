package main

import (
	"go-web-native/config"
	"go-web-native/controllers/authcontroller"
	"go-web-native/controllers/capstercontroller"
	controllers "go-web-native/controllers/homecontroller"
	"go-web-native/controllers/ordercontroller"
	"go-web-native/controllers/servicecontroller"
	"log"
	"net/http"
)

// Middleware function to check authentication
func requireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := authcontroller.GetSession(r)
		userID, ok := session.Values["user_id"]
		if !ok || userID == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	// Database connection
	config.ConnectDB()

	// Routes
	// 1. Homepage
	http.HandleFunc("/", requireLogin(controllers.Index))

	// 2. Capster
	http.HandleFunc("/capsters", capstercontroller.Index)
	http.HandleFunc("/capsters/add", capstercontroller.Add)
	http.HandleFunc("/capsters/edit", capstercontroller.Edit)
	http.HandleFunc("/capsters/delete", capstercontroller.Delete)

	// 3. Service
	http.HandleFunc("/services", servicecontroller.Index)
	http.HandleFunc("/services/add", servicecontroller.Add)
	http.HandleFunc("/services/edit", servicecontroller.Edit)
	http.HandleFunc("/services/delete", servicecontroller.Delete)

	// 4. Order
	http.HandleFunc("/orders", ordercontroller.Index)
	http.HandleFunc("/orders/add", ordercontroller.Add)
	http.HandleFunc("/orders/edit", ordercontroller.Edit)
	http.HandleFunc("/orders/delete", ordercontroller.Delete)
	http.HandleFunc("/orders/detail", ordercontroller.OrderDetail)

	// 5. Auth
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)

	// Run server
	log.Println("Server running on port: 8909")
	log.Fatal(http.ListenAndServe(":8909", nil))
}
