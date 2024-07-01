package authcontroller

import (
	"go-web-native/models"
	"html/template"
	"net/http"
	"strings"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}

	usernameOrEmail := r.FormValue("username")
	password := r.FormValue("password")

	// Fetch user based on username or email
	user, err := models.GetUserByUsernameOrEmail(usernameOrEmail, usernameOrEmail)
	if err != nil {
		// Handle error, redirect to login page with error message
		http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
		return
	}

	// Verify password
	err = models.VerifyPassword(user.Password, password)
	if err != nil {
		// Handle password mismatch error, redirect to login page with error message
		http.Redirect(w, r, "/login?error=2", http.StatusSeeOther)
		return
	}

	// If login successful, store user_id in session
	session, _ := GetSession(r)
	session.Values["user_id"] = user.Id
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "register.html", nil)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Check if username already exists
	_, err := models.GetUserByUsername(username)
	if err == nil {
		http.Error(w, "Username sudah terdaftar", http.StatusBadRequest)
		return
	}

	// Check if email already exists
	_, err = models.GetUserByEmail(email)
	if err == nil {
		http.Error(w, "Email sudah terdaftar", http.StatusBadRequest)
		return
	}

	// Validate email domain
	if !strings.HasSuffix(email, "@barbermasbro.com") {
		http.Error(w, "Email harus menggunakan domain @barbermasbro.com", http.StatusBadRequest)
		return
	}

	// Create new user
	err = models.CreateUser(username, email, password)
	if err != nil {
		http.Error(w, "Gagal mendaftarkan pengguna", http.StatusInternalServerError)
		return
	}

	// Redirect to login page after successful registration
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear user_id session on logout
	session, _ := GetSession(r)
	delete(session.Values, "user_id")
	session.Save(r, w)

	// Redirect to login page after logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
