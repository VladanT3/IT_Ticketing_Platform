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

func GetAllTeams() []Team {
	var db *sql.DB = database.DB_Connection
	teams := []Team{}

	query := `select * from team order by team_name;`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("error getting all teams", "error message", err)
		return []Team{}
	}
	defer rows.Close()

	team := Team{}
	for rows.Next() {
		err = rows.Scan(&team.Team_ID, &team.Team_Name, &team.Manager_ID)
		if err != nil {
			slog.Error("error getting all teams", "error message", err)
			return []Team{}
		}

		teams = append(teams, team)
	}

	return teams
}

func GetTeamIDByName(name string) (uuid.UUID, error) {
	var db *sql.DB = database.DB_Connection
	query := `select team_id from team where team_name = $1;`
	var team_id uuid.UUID

	err := db.QueryRow(query, name).Scan(&team_id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return team_id, nil
}
