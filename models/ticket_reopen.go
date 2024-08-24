package models

import (
	"time"

	"github.com/google/uuid"
)

type TicketReopen struct {
	Reopen_ID     uuid.UUID
	Ticket_ID     uuid.UUID
	Reopened_by   uuid.UUID
	Reopen_Reason string
	Reopen_Date   time.Time
}
