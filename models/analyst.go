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

func GetAnalyst(analystID string) Analyst {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}

	query := `select * from analyst where analyst_id = $1;`
	err := db.QueryRow(query, analystID).Scan(
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
			return Analyst{}
		}
		log.Fatal("error getting analyst: ", err)
	}

	return analyst
}
