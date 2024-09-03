package models

import (
	"database/sql"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type AssignmentHistory struct {
	Assignment_ID       uuid.UUID
	Ticket_ID           uuid.UUID
	Assigned_From       uuid.NullUUID
	Assigned_to_Analyst uuid.NullUUID
	Assigned_to_Team    uuid.NullUUID
	Assignment_Message  string
	Assignment_Date     time.Time
}

func AssignTicket(ticket_id string, assigned_from string, assigned_to_analyst string, assigned_to_team string, message string) error {
	var db *sql.DB = database.DB_Connection
	query := `insert into assignment_history values(gen_random_uuid(), $1, $2, $3, $4, $5, current_timestamp);`
	_, err := db.Exec(query, ticket_id, assigned_from, assigned_to_analyst, assigned_to_team, message)
	if err != nil {
		return err
	}

	query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1 where analyst_id = $1;`
	_, err = db.Exec(query, assigned_from)
	if err != nil {
		return err
	}

	query = `update analyst set number_of_open_tickets = number_of_open_tickets + 1 where analyst_id = $1;`
	_, err = db.Exec(query, assigned_to_analyst)
	if err != nil {
		return err
	}

	return nil
}
