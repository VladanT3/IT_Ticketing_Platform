package handlers

import (
	"net/http"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
	"github.com/VladanT3/IT_Ticketing_Platform/views/team"
	"github.com/VladanT3/IT_Ticketing_Platform/views/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Profile(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	pass_change_check_cookie, err := r.Cookie("password_changed")
	if err != nil {
		return Render(w, r, user.Profile(LoggedInUser, LoggedInUserType, false))
	}

	pass_change_check := pass_change_check_cookie.Value

	if pass_change_check == "yes" {
		return Render(w, r, user.Profile(LoggedInUser, LoggedInUserType, true))
	} else {
		return Render(w, r, user.Profile(LoggedInUser, LoggedInUserType, false))
	}
}

func GetTeamsAnalysts(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	team := r.FormValue("team")

	if team == "none" {
		return Render(w, r, user.SelectAnalyst(models.GetAllAnalysts()))
	} else {
		analysts := models.GetTeamsAnalysts(team)
		return Render(w, r, user.SelectAnalyst(analysts))
	}
}

func GetAnalystsTeam(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	analyst_id := r.FormValue("analyst")

	analyst := models.GetAnalyst(analyst_id)

	if analyst_id == "none" {
		return Render(w, r, team.SelectTeam(analyst))
	} else {
		return Render(w, r, team.SelectTeam(analyst))
	}
}

func ShowUserView(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "User View"))
}

func ShowTeamView(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "manager" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of managerial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, user.UserView(LoggedInUserType, LoggedInUser, "Team View"))
}

func FilterUsers(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "analyst" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of higher credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	search_term := r.FormValue("search")
	view_type := r.FormValue("view_type")
	user_type := r.FormValue("user_type")

	analysts := models.FilterUsers(search_term, view_type, LoggedInUser.Team_ID.UUID.String(), user_type)

	return Render(w, r, user.Users(analysts, view_type))
}

func ShowUserForm(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "analyst" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of higher credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	analyst_id := chi.URLParam(r, "analyst_id")
	view_type := r.FormValue("view_type")

	analyst_exists, err := models.UserExists(analyst_id)
	if err != nil {
		err_msg := "Internal server error:\nerror checking if user exists: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if !analyst_exists {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "No user found: user doesn't exist or it was removed!"))
	}

	analyst := models.GetAnalyst(analyst_id)

	w.Header().Add("HX-Redirect", "/user/"+analyst_id)

	return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, models.Analyst{}, [5]bool{false, false, false, false, false}, true, "update", "", false))
}

func ShowNewUserForm(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	view_type := r.FormValue("view_type")

	return Render(w, r, user.UserForm(LoggedInUserType, models.Analyst{}, view_type, models.Analyst{}, [5]bool{false, false, false, false, false}, true, "create", "", false))
}

func UserRedirect(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "analyst" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of higher credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	mode := r.FormValue("mode")
	analyst_id := r.FormValue("analyst_id")
	user_type := r.FormValue("user_type")
	team_id := r.FormValue("team")
	team_name := models.GetTeam(team_id).Team_Name
	view_type := r.FormValue("view_type")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	phone_number := r.FormValue("phone_number")

	var new_analyst models.Analyst = models.Analyst{
		Analyst_ID:   uuid.MustParse(analyst_id),
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Password:     password,
		Phone_Number: phone_number,
		Team_ID: uuid.NullUUID{
			UUID:  uuid.MustParse(team_id),
			Valid: true,
		},
	}

	var errs [5]bool = [5]bool{false, false, false, false, false}
	var errCounter int
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
	if mode == "create" {
		if password == "" {
			errs[3] = true
			errCounter += 1
		}
	}
	if phone_number == "" {
		errs[4] = true
		errCounter += 1
	}

	if errCounter > 0 {
		return Render(w, r, user.UserForm(LoggedInUserType, models.GetAnalyst(analyst_id), view_type, new_analyst, errs, true, mode, user_type, false))
	}

	if user_type == "admin" && team_name != "Administrators" {
		return Render(w, r, user.UserForm(LoggedInUserType, models.GetAnalyst(analyst_id), view_type, new_analyst, errs, true, mode, user_type, true))
	} else if user_type != "admin" && team_name == "Administrators" {
		return Render(w, r, user.UserForm(LoggedInUserType, models.GetAnalyst(analyst_id), view_type, new_analyst, errs, true, mode, user_type, true))
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "view_type",
		Value:   view_type,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "analyst_id",
		Value:   analyst_id,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "first_name",
		Value:   first_name,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "last_name",
		Value:   last_name,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "email",
		Value:   email,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "password",
		Value:   password,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "phone_number",
		Value:   phone_number,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "team",
		Value:   team_id,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "user_type",
		Value:   user_type,
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})

	if view_type == "Team View" {
		http.Redirect(w, r, "/user/request/update/"+analyst_id, http.StatusSeeOther)
		return nil
	} else if view_type == "User View" {
		if mode == "create" {
			http.Redirect(w, r, "/user/create", http.StatusSeeOther)
			return nil
		} else if mode == "update" {
			http.Redirect(w, r, "/user/update/"+analyst_id, http.StatusSeeOther)
			return nil
		}
	}

	return nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

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
	var new_analyst models.Analyst = models.Analyst{
		Analyst_ID:   uuid.MustParse(analyst_id),
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Phone_Number: phone_number,
		Team_ID: uuid.NullUUID{
			UUID:  uuid.MustParse(team),
			Valid: true,
		},
	}

	_, email_exists, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nerror checking email validity when updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if email_exists {
		same_email, err := models.IsEmailSame(analyst_id, email)
		if err != nil {
			err_msg := "Internal server error:\nerror checking if email is new: " + err.Error()
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}
		if !same_email {
			return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, new_analyst, errs, false, "update", user_type, false))
		}
	}

	err = models.UpdateAnalyst(new_analyst, user_type)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating logged in user details after updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	http.Redirect(w, r, "/users/view", http.StatusSeeOther)

	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	analyst_id := chi.URLParam(r, "analyst_id")

	err := models.DeleteAnalyst(analyst_id)
	if err != nil {
		err_msg := "Internal server error:\nerror deleting user: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, user.DeletedUser())
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

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
	_, email_exists, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nerror checking email validity when updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if email_exists {
		return Render(w, r, user.UserForm(LoggedInUserType, models.Analyst{}, view_type, new_analyst, errs, false, "create", user_type, false))
	}

	err = models.CreateAnalyst(new_analyst, user_type)
	if err != nil {
		err_msg := "Internal server error:\nerror creating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	http.Redirect(w, r, "/users/view", http.StatusSeeOther)

	return nil
}

func RequestUserInfoChange(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "manager" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of managerial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

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
	var new_analyst models.Analyst = models.Analyst{
		Analyst_ID:   uuid.MustParse(analyst_id),
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Phone_Number: phone_number,
		Team_ID: uuid.NullUUID{
			UUID:  uuid.MustParse(team),
			Valid: true,
		},
	}

	_, email_exists, err := models.CheckEmail(email)
	if err != nil {
		err_msg := "Internal server error:\nerror checking email validity when updating user: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if email_exists {
		same_email, err := models.IsEmailSame(analyst_id, email)
		if err != nil {
			err_msg := "Internal server error:\nerror checking if email is new: " + err.Error()
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}
		if !same_email {
			return Render(w, r, user.UserForm(LoggedInUserType, analyst, view_type, new_analyst, errs, false, "update", user_type, false))
		}
	}

	ticket_id, err := models.RequestUserInfoChange(new_analyst, user_type, LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror creating a user info change request: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	http.Redirect(w, r, "/ticket/"+ticket_id, http.StatusSeeOther)
	return nil
}

func ShowChangePasswordForm(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, user.ChangePasswordForm(LoggedInUserType, "", "", false))
}

func ChangePassword(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	password := r.FormValue("password")
	repeat_password := r.FormValue("repeat_password")

	if password != repeat_password {
		return Render(w, r, user.ChangePasswordForm(LoggedInUserType, password, repeat_password, true))
	}

	err := models.ChangePassword(LoggedInUser.Analyst_ID.String(), password)
	if err != nil {
		err_msg := "Internal server error:\nerror changing password: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "password_changed",
		Value:   "yes",
		Expires: time.Time.Add(time.Now(), time.Second*5),
		Path:    "/",
	})

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
	return nil
}

func ShowPasswordChangeSuccess(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	return Render(w, r, user.PasswordChangeSuccess())
}
