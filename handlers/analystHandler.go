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
		Name:    "view_type",
		Value:   "User View",
		Expires: time.Time.Add(time.Now(), time.Hour*1),
	})
	return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "User View"))
}

func ShowTeamView(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:    "view_type",
		Value:   "Team View",
		Expires: time.Time.Add(time.Now(), time.Hour*1),
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

	return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, models.Analyst{}, [5]bool{false, false, false, false, false}, true, "update", ""))
}

func ShowNewUserForm(w http.ResponseWriter, r *http.Request) error {
	view_type_cookie, err := r.Cookie("view_type")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	view_type := view_type_cookie.Value

	return Render(w, r, user.UserForm(LoggedInUserType, models.Analyst{}, view_type, models.Analyst{}, [5]bool{false, false, false, false, false}, true, "create", ""))
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
	http.SetCookie(w, &http.Cookie{
		Name:    "first_name",
		Value:   r.FormValue("first_name"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "last_name",
		Value:   r.FormValue("last_name"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "email",
		Value:   r.FormValue("email"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "password",
		Value:   r.FormValue("password"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "phone_number",
		Value:   r.FormValue("phone_number"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "team",
		Value:   r.FormValue("team"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "user_type",
		Value:   r.FormValue("user_type"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})

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
	view_type_cookie, err := r.Cookie("view_type")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'view_type' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	view_type := view_type_cookie.Value
	analyst_id_cookie, err := r.Cookie("analyst_id")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'analyst_id' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	analyst_id := analyst_id_cookie.Value
	first_name_cookie, err := r.Cookie("first_name")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'first_name' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	first_name := first_name_cookie.Value
	last_name_cookie, err := r.Cookie("last_name")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'last_name' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	last_name := last_name_cookie.Value
	email_cookie, err := r.Cookie("email")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'email' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	email := email_cookie.Value
	phone_number_cookie, err := r.Cookie("phone_number")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'phone_number' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	phone_number := phone_number_cookie.Value
	team_cookie, err := r.Cookie("team")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'team' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	team := team_cookie.Value
	user_type_cookie, err := r.Cookie("user_type")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'user_type' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	user_type := user_type_cookie.Value

	analyst := models.GetAnalyst(analyst_id)

	var errs [5]bool = [5]bool{false, false, false, false, false}
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
		errs[4] = true
		errCounter += 1
	}

	_, email_exists, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nerror checking email validity when updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if email_exists {
		return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, new_analyst, errs, false, "update", user_type))
	}

	if errCounter > 0 {
		return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, new_analyst, errs, true, "update", user_type))
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
	view_type_cookie, err := r.Cookie("view_type")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'view_type' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	view_type := view_type_cookie.Value
	first_name_cookie, err := r.Cookie("first_name")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'first_name' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	first_name := first_name_cookie.Value
	last_name_cookie, err := r.Cookie("last_name")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'last_name' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	last_name := last_name_cookie.Value
	email_cookie, err := r.Cookie("email")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'email' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	email := email_cookie.Value
	password_cookie, err := r.Cookie("password")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'password' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	password := password_cookie.Value
	phone_number_cookie, err := r.Cookie("phone_number")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'phone_number' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	phone_number := phone_number_cookie.Value
	team_cookie, err := r.Cookie("team")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'team' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	team := team_cookie.Value
	user_type_cookie, err := r.Cookie("user_type")
	if err != nil {
		err_msg := "Internal server error:\ncookie with name 'user_type' doesn't exist!"
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	user_type := user_type_cookie.Value

	var errs [5]bool = [5]bool{false, false, false, false, false}
	errCounter := 0
	var new_analyst models.Analyst = models.Analyst{
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Password:     password,
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
	if password == "" {
		errs[3] = true
		errCounter += 1
	}
	if phone_number == "" {
		errs[4] = true
		errCounter += 1
	}

	_, email_exists, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nerror checking email validity when updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if email_exists {
		return Render(w, r, user.UserForm(LoggedInUserType, models.Analyst{}, view_type, new_analyst, errs, false, "update", user_type))
	}

	if errCounter > 0 {
		return Render(w, r, user.UserForm(LoggedInUserType, models.Analyst{}, view_type, new_analyst, errs, true, "update", user_type))
	}

	err = models.CreateAnalyst(new_analyst, user_type)
	if err != nil {
		err_msg := "Internal server error:\nerror creating user: " + err.Error()
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

func ChangePassword(w http.ResponseWriter, r *http.Request) error {
	return nil
}
