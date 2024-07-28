package handlers

import (
	"net/http"
	"os"

	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, login.Login("", "", "", ""))
}
