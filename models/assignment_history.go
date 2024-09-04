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
	old_ticket, err := GetTicket(ticket_id)
	if err != nil {
		return err
	}

	if assigned_to_analyst != "none" {
		query := `insert into assignment_history values(gen_random_uuid(), $1, $2, $3, $4, $5, current_timestamp);`
		_, err := db.Exec(query, ticket_id, assigned_from, assigned_to_analyst, assigned_to_team, message)
		if err != nil {
			return err
		}

		if old_ticket.Assigned_Analyst.UUID.String() != assigned_from {
			if old_ticket.Assigned_Analyst.Valid {
				query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1 where analyst_id = $1;`
				_, err = db.Exec(query, old_ticket.Assigned_Analyst.UUID.String())
				if err != nil {
					return err
				}
			}

			query = `update analyst set number_of_open_tickets = number_of_open_tickets + 1 where analyst_id = $1;`
			_, err = db.Exec(query, assigned_to_analyst)
			if err != nil {
				return err
			}
		} else {
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
		}

		query = `update ticket set assigned_analyst = $1, assigned_team = $2 where ticket_id = $3;`
		_, err = db.Exec(query, assigned_to_analyst, assigned_to_team, ticket_id)
		if err != nil {
			return err
		}
	} else {
		query := `insert into assignment_history values(gen_random_uuid(), $1, $2, null, $3, $4, current_timestamp);`
		_, err := db.Exec(query, ticket_id, assigned_from, assigned_to_team, message)
		if err != nil {
			return err
		}

		if old_ticket.Assigned_Analyst.UUID.String() != assigned_from {
			if old_ticket.Assigned_Analyst.Valid {
				query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1 where analyst_id = $1;`
				_, err = db.Exec(query, old_ticket.Assigned_Analyst.UUID.String())
				if err != nil {
					return err
				}
			}
		} else {
			query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1 where analyst_id = $1;`
			_, err = db.Exec(query, assigned_from)
			if err != nil {
				return err
			}
		}

		query = `update ticket set assigned_analyst = null, assigned_team = $1 where ticket_id = $2;`
		_, err = db.Exec(query, assigned_to_team, ticket_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAllTicketsAssignments(ticket_id string) ([]AssignmentHistory, error) {
	var db *sql.DB = database.DB_Connection
	assignments := []AssignmentHistory{}
	assignment := AssignmentHistory{}

	query := `select * from assignment_history where ticket_id = $1;`

	rows, err := db.Query(query, ticket_id)
	if err != nil {
		return []AssignmentHistory{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&assignment.Assignment_ID,
			&assignment.Ticket_ID,
			&assignment.Assigned_From,
			&assignment.Assigned_to_Analyst,
			&assignment.Assigned_to_Team,
			&assignment.Assignment_Message,
			&assignment.Assignment_Date,
		)
		if err != nil {
			return []AssignmentHistory{}, err
		}

		assignments = append(assignments, assignment)
	}

	return assignments, nil
}
