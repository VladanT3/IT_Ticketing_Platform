package models

import (
	"database/sql"
	"log"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Analyst struct {
	Analyst_id               uuid.UUID
	First_name               string
	Last_name                string
	Email                    string
	Password                 string
	Phone_number             string
	Team_id                  uuid.NullUUID
	Number_of_open_tickets   int
	Number_of_opened_tickets int
	Number_of_closed_tickets int
}

var dbConn *sql.DB = database.DB_Connection

func GetAnalystsTeam(analyst Analyst) string {
	var teamName string
	query := `select t.team_name from team t join analyst a on t.team_id = a.team_id where a.analyst_id = $1;`
	err := dbConn.QueryRow(query, analyst.Analyst_id).Scan(&teamName)
	if err != nil {
		if err == sql.ErrNoRows {
			teamName = ""
		}
		log.Fatal("error getting team name: ", err)
	}

	return teamName
}
