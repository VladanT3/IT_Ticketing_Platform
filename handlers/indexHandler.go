package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, login.Login("", ""))
}
