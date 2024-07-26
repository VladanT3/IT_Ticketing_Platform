package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	LoggedInUser = models.Analyst{}
	LoggedInUserType = ""

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
