package models

import (
	"database/sql"
	"log"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Team struct {
	Team_ID    uuid.UUID
	Team_Name  string
	Manager_ID uuid.NullUUID
}

func GetAnalystsTeam(analystID string) Team {
	var db *sql.DB = database.DB_Connection
	team := Team{}
	query := `select t.* from team t join analyst a on t.team_id = a.team_id where a.analyst_id = $1;`
	err := db.QueryRow(query, analystID).Scan(&team.Team_ID, &team.Team_Name, &team.Manager_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Team{}
		}
		log.Fatal("error getting team name: ", err)
	}

	return team
}
