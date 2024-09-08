package handlers

import (
	"net/http"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/team"
	"github.com/VladanT3/IT_Ticketing_Platform/views/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Profile(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, user.Profile(LoggedInUser, LoggedInUserType))
}

func GetTeamsAnalysts(w http.ResponseWriter, r *http.Request) error {
	team := r.FormValue("team")

	if team == "none" {
		return Render(w, r, user.SelectAnalyst(models.GetAllAnalysts()))
	} else {
		analysts := models.GetTeamsAnalysts(team)
		return Render(w, r, user.SelectAnalyst(analysts))
	}
}

func GetAnalystsTeam(w http.ResponseWriter, r *http.Request) error {
	analyst_id := r.FormValue("analyst")

	analyst := models.GetAnalyst(analyst_id)

	if analyst_id == "none" {
		return Render(w, r, team.SelectTeam(analyst))
	} else {
		return Render(w, r, team.SelectTeam(analyst))
	}
}

func ShowUserView(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:  "view_type",
		Value: "User View",
	})
	return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "User View"))
}

func ShowTeamView(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:  "view_type",
		Value: "Team View",
	})
	return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "Team View"))
}

func FilterUsers(w http.ResponseWriter, r *http.Request) error {
	search_term := r.FormValue("search")
	view_type := r.FormValue("view_type")
	user_type := r.FormValue("user_type")

	analysts := models.FilterUsers(search_term, view_type, LoggedInUser.Team_ID.UUID.String(), user_type)

	return Render(w, r, user.Users(analysts))
}

func ShowUserForm(w http.ResponseWriter, r *http.Request) error {
	analyst_id := chi.URLParam(r, "analyst_id")
	view_type_cookie, err := r.Cookie("view_type")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	view_type := view_type_cookie.Value

	analyst := models.GetAnalyst(analyst_id)

	return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, models.Analyst{}, [4]bool{false, false, false, false}, true))
}

func ShowNewUserForm(w http.ResponseWriter, r *http.Request) error {
	view_type_cookie, err := r.Cookie("view_type")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	view_type := view_type_cookie.Value

	return Render(w, r, user.UserForm(LoggedInUserType, models.Analyst{}, view_type, models.Analyst{}, [4]bool{false, false, false, false}, true))
}

func UserRedirect(w http.ResponseWriter, r *http.Request) error {
	mode := r.FormValue("mode")
	analyst_id := r.FormValue("analyst_id")

	http.SetCookie(w, &http.Cookie{
		Name:    "view_type",
		Value:   r.FormValue("view_type"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "analyst_id",
		Value:   r.FormValue("analyst_id"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	//TODO: continue this

	if mode == "create" {
		http.Redirect(w, r, "/user/create", http.StatusSeeOther)
		return nil
	} else if mode == "update" {
		http.Redirect(w, r, "/user/update"+analyst_id, http.StatusSeeOther)
		return nil
	}

	return nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	view_type := r.FormValue("view_type")
	analyst_id := r.FormValue("analyst_id")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	email := r.FormValue("email")
	phone_number := r.FormValue("phone_number")
	team := r.FormValue("team")

	analyst := models.GetAnalyst(analyst_id)

	var errs [4]bool = [4]bool{false, false, false, false}
	errCounter := 0
	var new_analyst models.Analyst = models.Analyst{
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Phone_Number: phone_number,
		Team_ID: uuid.NullUUID{
			UUID:  uuid.MustParse(team),
			Valid: true,
		},
	}

	if first_name == "" {
		errs[0] = true
		errCounter += 1
	}
	if last_name == "" {
		errs[1] = true
		errCounter += 1
	}
	if email == "" {
		errs[2] = true
		errCounter += 1
	}
	if phone_number == "" {
		errs[3] = true
		errCounter += 1
	}

	_, email_exists, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nerror checking email validity when updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if email_exists {
		return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, new_analyst, errs, false))
	}

	if errCounter > 0 {
		return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, new_analyst, errs, true))
	}

	err = models.UpdateAnalyst(new_analyst)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating logged in user details after updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if view_type == "User View" {
		w.Header().Add("HX-Redirect", "/users/view")
		return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "User View"))
	} else if view_type == "Team View" {
		w.Header().Add("HX-Redirect", "/users/team/view")
		return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "Team View"))
	}

	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	analyst_id := chi.URLParam(r, "analyst_id")

	err := models.DeleteAnalyst(analyst_id)
	if err != nil {
		err_msg := "Internal server error:\nerror deleting user: " + err.Error()
		w.Header().Add("ErrorMessage", err_msg)
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, user.DeletedUser())
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {

	return nil
}
