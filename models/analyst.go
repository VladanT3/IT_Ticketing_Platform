package models

import (
	"database/sql"
	"log/slog"
	"strings"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Analyst struct {
	Analyst_ID               uuid.UUID
	First_Name               string
	Last_Name                string
	Email                    string
	Password                 string
	Phone_Number             string
	Team_ID                  uuid.NullUUID
	Number_of_Open_Tickets   int
	Number_of_Opened_Tickets int
	Number_of_Closed_Tickets int
}

func CheckEmail(email string) (Analyst, bool, error) {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}
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
			return Analyst{}, false, nil
		}
		return Analyst{}, false, err
	}

	return analyst, true, nil
}

func CheckPassword(password string, email string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	var correctPassword bool
	query := `select (password = crypt($1, password)) as password from analyst where email = $2;`
	err := db.QueryRow(query, password, email).Scan(&correctPassword)
	if err != nil {
		return false, err
	}
	if !correctPassword {
		return false, nil
	} else {
		return true, nil
	}
}

func GetAnalyst(analystID string) Analyst {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}

	query := `select * from analyst where analyst_id = $1;`
	err := db.QueryRow(query, analystID).Scan(
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
			return Analyst{}
		}
		slog.Error("error getting analyst", "error message", err)
		return Analyst{}
	}

	return analyst
}

func UpdateLoggedInUser(loggedInUser Analyst) (Analyst, error) {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}
	query := `select * from analyst where analyst_id = $1;`

	err := db.QueryRow(query, loggedInUser.Analyst_ID).Scan(
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
		return Analyst{}, err
	}

	return analyst, nil
}

func GetAllAnalysts() []Analyst {
	var db *sql.DB = database.DB_Connection
	analysts := []Analyst{}

	query := `select * from analyst order by first_name, last_name;`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("error getting all analysts", "error message", err)
		return []Analyst{}
	}
	defer rows.Close()

	analyst := Analyst{}
	for rows.Next() {
		err = rows.Scan(
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
			slog.Error("error scanning all analysts", "error message", err)
			return []Analyst{}
		}

		analysts = append(analysts, analyst)
	}

	return analysts
}

func GetTeamsAnalysts(team_id string) []Analyst {
	var db *sql.DB = database.DB_Connection
	analysts := []Analyst{}

	query := `select * from analyst where team_id = $1 order by first_name, last_name;`
	rows, err := db.Query(query, team_id)
	if err != nil {
		slog.Error("error getting a teams analysts", "error message", err)
		return []Analyst{}
	}
	defer rows.Close()

	analyst := Analyst{}
	for rows.Next() {
		err = rows.Scan(
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
			slog.Error("error scanning a teams analysts", "error message", err)
			return []Analyst{}
		}

		analysts = append(analysts, analyst)
	}

	return analysts
}

func FilterUsers(search_term string, view_type string, team_id string, user_type string) []Analyst {
	var db *sql.DB = database.DB_Connection
	analysts := []Analyst{}
	analyst := Analyst{}
	search_term = strings.ToLower(search_term)
	search_term = "%" + search_term + "%"
	var queryArgs []any
	var query string

	if user_type == "Managers" {
		query = `
			select *
			from analyst an join manager m on a.Analyst_ID = m.Manager_ID
			where (lower(an.first_name) like $1 or
			lower(an.last_name) like $1)
		`
	} else if user_type == "Administrators" {
		query = `
			select *
			from analyst an join administrator ad on an.Analyst_ID = ad.Administrator_ID
			where (lower(an.first_name) like $1 or
			lower(an.last_name) like $1)
		`
	} else {
		query = `
			select * 
			from analyst an
			where (lower(an.first_name) like $1 or
			lower(an.last_name) like $1)
		`
	}

	queryArgs = append(queryArgs, search_term)

	if view_type == "Team View" {
		query += `
			and an.team_id = $2
		`
		queryArgs = append(queryArgs, team_id)
	}

	query += `
		order by an.first_name, an.last_name;
	`
	rows, err := db.Query(query, queryArgs...)
	if err != nil {
		slog.Error("error filtering analysts", "error message", err)
		return []Analyst{}
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
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
			slog.Error("error scanning filtered analysts", "error message", err)
			return []Analyst{}
		}
		analysts = append(analysts, analyst)
	}

	return analysts
}

func UpdateAnalyst(new_analyst Analyst) error {
	var db *sql.DB = database.DB_Connection
	query := `
		update analyst set
		first_name = $1,
		last_name = $2,
		email = $3,
		phone_number = $4,
		team_id = $5
		where analyst_id = $6;
	`
	_, err := db.Exec(query, new_analyst.First_Name, new_analyst.Last_Name, new_analyst.Email, new_analyst.Phone_Number, new_analyst.Team_ID.UUID, new_analyst.Analyst_ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAnalyst(analyst_id string) error {
	var db *sql.DB = database.DB_Connection
	query := `delete from analyst where analyst_id = $1;`

	_, err := db.Exec(query, analyst_id)
	if err != nil {
		return err
	}

	return nil
}
