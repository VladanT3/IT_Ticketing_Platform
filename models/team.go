package models

import (
	"database/sql"
	"log/slog"

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
		slog.Error("error getting analysts team", "error message", err)
		return Team{}
	}

	return team
}

func GetTeam(teamID string) Team {
	var db *sql.DB = database.DB_Connection
	team := Team{}

	query := `select * from team where team_id = $1;`
	err := db.QueryRow(query, teamID).Scan(&team.Team_ID, &team.Team_Name, &team.Manager_ID)
	if err != nil {
		slog.Error("error getting team", "error message", err)
		return Team{}
	}

	return team
}
