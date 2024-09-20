package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
)

var LoggedInUser models.Analyst = models.Analyst{}
var LoggedInUserType string = ""

func Login(w http.ResponseWriter, r *http.Request) error {
	analyst := models.Analyst{}
	email := r.FormValue("email")
	password := r.FormValue("password")

	analyst, valid_email, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nemail error: " + err.Error()
		return Render(w, r, layouts.ErrorMessage("", err_msg))
	}
	if !valid_email {
		return Render(w, r, login.Login(true, false, email, password))
	}

	valid_pass, err := models.CheckPassword(password, email)
	if err != nil {
		err_msg := "Internal server error:\npassword error: " + err.Error()
		return Render(w, r, layouts.ErrorMessage("", err_msg))
	}
	if !valid_pass {
		return Render(w, r, login.Login(false, true, email, password))
	}

	LoggedInUser = analyst

	is_manager, err := models.IsUserManager(analyst.Analyst_ID.String())
	if err != nil {
		err_msg := "Internal server error:\nerror checking if user is manager: " + err.Error()
		return Render(w, r, layouts.ErrorMessage("", err_msg))
	}
	if is_manager {
		LoggedInUserType = "manager"
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	is_admin, err := models.IsUserAdmin(analyst.Analyst_ID.String())
	if err != nil {
		err_msg := "Internal server error:\nerror checking if user is admin: " + err.Error()
		return Render(w, r, layouts.ErrorMessage("", err_msg))
	}
	if is_admin {
		LoggedInUserType = "admin"
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	LoggedInUserType = "analyst"
	http.Redirect(w, r, "/profile", http.StatusSeeOther)

	return nil
}
