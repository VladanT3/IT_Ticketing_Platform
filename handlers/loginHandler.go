package handlers

import (
	"database/sql"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
)

var LoggedInUser models.Analyst = models.Analyst{}
var LoggedInUserType string

func Login(w http.ResponseWriter, r *http.Request) error {
	analyst := models.Analyst{}
	email := r.FormValue("email")
	password := r.FormValue("password")
	var db *sql.DB = database.DB_Connection

	query := `select * from analyst where email = $1;`
	err := db.QueryRow(query, email).Scan(
		&analyst.Analyst_ID,
		&analyst.First_Name,
		&analyst.Last_Name,
		&analyst.Email,
		&analyst.Password,
		&analyst.Phone_Number,
		&analyst.Team_ID,
		&analyst.Number_of_Open_Tickets,
		&analyst.Number_of_Opened_Tickets,
		&analyst.Number_of_Closed_Tickets,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Render(w, r, login.Login("Incorrect email!", "", email, password))
		}
		errMsg := "email error: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	var correctPassword bool
	query = `select (password = crypt($1, password)) as password from analyst where email = $2;`
	err = db.QueryRow(query, password, analyst.Email).Scan(&correctPassword)
	if err != nil {
		errMsg := "password error: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	if !correctPassword {
		return Render(w, r, login.Login("", "Incorrect password!", email, password))
	}

	LoggedInUser = analyst

	var isManager int
	query = `select count(*) as isManager from manager where manager_id = $1;`
	err = db.QueryRow(query, analyst.Analyst_ID).Scan(&isManager)
	if err != nil {
		errMsg := "error checking if user is manager: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	if isManager == 1 {
		LoggedInUserType = "manager"
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	var isAdmin int
	query = `select count(*) as isAdmin from administrator where administrator_id = $1;`
	err = db.QueryRow(query, analyst.Analyst_ID).Scan(&isAdmin)
	if err != nil {
		errMsg := "error checking if user is admin: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	if isAdmin == 1 {
		LoggedInUserType = "admin"
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	LoggedInUserType = "analyst"
	http.Redirect(w, r, "/profile", http.StatusSeeOther)

	return nil
}
