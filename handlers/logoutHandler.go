package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
)

func Logout(w http.ResponseWriter, r *http.Request) error {
	LoggedInUser = models.Analyst{}
	LoggedInUserType = ""

	w.Header().Add("HX-Redirect", "/")
	return Render(w, r, login.Login("", "", "", ""))
}
