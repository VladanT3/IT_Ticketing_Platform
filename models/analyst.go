package models

import (
	"database/sql"
	"log/slog"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/views/analyst"
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

	query := `select * from analyst;`
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

	query := `select * from analyst where team_id = $1;`
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
