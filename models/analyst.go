package models

import "github.com/google/uuid"

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
