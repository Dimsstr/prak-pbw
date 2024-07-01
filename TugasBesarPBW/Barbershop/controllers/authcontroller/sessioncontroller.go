package authcontroller

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session")
}
