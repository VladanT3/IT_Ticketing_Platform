package models

import (
	"database/sql"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type TicketReopen struct {
	Reopen_ID     uuid.UUID
	Ticket_ID     uuid.UUID
	Reopened_by   uuid.UUID
	Reopen_Reason string
	Reopened_Date time.Time
}

func ReopenTicket(ticket_id string, reason string, analyst_id string) error {
	var db *sql.DB = database.DB_Connection
	query := `insert into ticket_reopen values(gen_random_uuid(), $1, $2, $3, current_timestamp);`

	_, err := db.Exec(query, ticket_id, analyst_id, reason)
	if err != nil {
		return err
	}

	old_ticket, err := GetTicket(ticket_id)
	if err != nil {
		return err
	}

	query = `update ticket set status = 'Open', updated_at = current_timestamp, closed_date = null, closed_by = null where ticket_id = $1;`
	_, err = db.Exec(query, ticket_id)
	if err != nil {
		return err
	}

	if old_ticket.Assigned_Analyst.UUID == old_ticket.Closed_by.UUID {
		query = `update analyst set number_of_closed_tickets = number_of_closed_tickets - 1, number_of_open_tickets = number_of_open_tickets + 1 where analyst_id = $1;`
		_, err = db.Exec(query, old_ticket.Assigned_Analyst.UUID)
		if err != nil {
			return err
		}

	} else {
		query = `update analyst set number_of_closed_tickets = number_of_closed_tickets - 1 where analyst_id = $1;`
		_, err = db.Exec(query, old_ticket.Closed_by.UUID)
		if err != nil {
			return err
		}

		query = `update analyst set number_of_open_tickets = number_of_open_tickets + 1 where analyst_id = $1;`
		_, err = db.Exec(query, old_ticket.Assigned_Analyst.UUID)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetTicketReopens(ticket_id string) ([]TicketReopen, error) {
	var db *sql.DB = database.DB_Connection
	query := `select * from ticket_reopen where ticket_id = $1 order by reopened_date desc;`
	reopens := []TicketReopen{}
	reopen := TicketReopen{}

	rows, err := db.Query(query, ticket_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&reopen.Reopen_ID,
			&reopen.Ticket_ID,
			&reopen.Reopened_by,
			&reopen.Reopen_Reason,
			&reopen.Reopened_Date,
		)
		if err != nil {
			return nil, err
		}
		reopens = append(reopens, reopen)
	}

	return reopens, nil
}
