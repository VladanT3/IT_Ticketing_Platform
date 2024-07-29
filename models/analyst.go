package models

import (
	"database/sql"
	"log"

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
		log.Fatal("error getting analyst: ", err)
	}

	return analyst
}
