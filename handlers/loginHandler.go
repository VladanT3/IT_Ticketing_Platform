package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
)

var LoggedInUser models.Analyst = models.Analyst{}
var LoggedInUserType string

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	analyst := models.Analyst{}
	email := r.FormValue("email")
	password := r.FormValue("password")
	var dbConn *sql.DB = database.DB_Connection

	query := `select * from analyst where email = $1;`
	err := dbConn.QueryRow(query, email).Scan(
		&analyst.Analyst_id,
		&analyst.First_name,
		&analyst.Last_name,
		&analyst.Email,
		&analyst.Password,
		&analyst.Phone_number,
		&analyst.Team_id,
		&analyst.Number_of_open_tickets,
		&analyst.Number_of_opened_tickets,
		&analyst.Number_of_closed_tickets,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Render(w, r, login.Login("Incorrect email!", "", email, password))
		}
		log.Fatal("email error: ", err)
	}

	var correctPassword bool
	query = `select (password = crypt($1, password)) as password from analyst where email = $2;`
	err = dbConn.QueryRow(query, password, analyst.Email).Scan(&correctPassword)
	if err != nil {
		log.Fatal("password error: ", err)
	}
	if !correctPassword {
		return Render(w, r, login.Login("", "Incorrect password!", email, password))
	}

	LoggedInUser = analyst

	var isManager int
	query = `select count(*) as isManager from manager where manager_id = $1;`
	err = dbConn.QueryRow(query, analyst.Analyst_id).Scan(&isManager)
	if err != nil {
		log.Fatal("manager check error: ", err)
	}
	if isManager == 1 {
		LoggedInUserType = "manager"
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	var isAdmin int
	query = `select count(*) as isAdmin from administrator where administrator_id = $1;`
	err = dbConn.QueryRow(query, analyst.Analyst_id).Scan(&isAdmin)
	if err != nil {
		log.Fatal("admin check error: ", err)
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
