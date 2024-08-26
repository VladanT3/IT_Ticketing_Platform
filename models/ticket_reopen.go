package models

import (
	"database/sql"
	"log"
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

func GetTicketReopens(ticket_id string) []TicketReopen {
	var db *sql.DB = database.DB_Connection
	query := `select * from ticket_reopen where ticket_id = $1 order by reopened_date desc;`
	reopens := []TicketReopen{}
	reopen := TicketReopen{}

	rows, err := db.Query(query, ticket_id)
	if err != nil {
		log.Fatal("error getting ticket reopens: ", err)
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
			log.Fatal("error scanning ticket reopens: ", err)
		}

		reopens = append(reopens, reopen)
	}

	return reopens
}
