package models

import (
	"database/sql"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Ticket struct {
	Ticket_ID        uuid.UUID
	Ticket_Number    int
	Type             string
	Status           string
	Category         uuid.NullUUID
	Subcategory      uuid.NullUUID
	Title            string
	Description      string
	Customer_Contact string
	Opened_Date      time.Time
	Updated_at       time.Time
	Closed_Date      sql.NullTime
	Assigned_Team    uuid.NullUUID
	Assigned_Analyst uuid.NullUUID
	Opened_by        uuid.NullUUID
	Closed_by        uuid.NullUUID
}

func GetAnalystsTickets(analystID string) []Ticket {
	var db *sql.DB = database.DB_Connection
	var tickets []Ticket

	query := `select t.* from ticket t join analyst a on t.assigned_analyst = a.analyst_id where a.analyst_id = $1 order by t.opened_date;`
	rows, err := db.Query(query, analystID)
	if err != nil {
		slog.Error("error getting analysts tickets", "error message", err)
		return []Ticket{}
	}
	defer rows.Close()

	ticket := Ticket{}
	for rows.Next() {
		err = rows.Scan(
			&ticket.Ticket_ID,
			&ticket.Ticket_Number,
			&ticket.Type,
			&ticket.Status,
			&ticket.Category,
			&ticket.Subcategory,
			&ticket.Title,
			&ticket.Description,
			&ticket.Customer_Contact,
			&ticket.Opened_Date,
			&ticket.Updated_at,
			&ticket.Closed_Date,
			&ticket.Assigned_Team,
			&ticket.Assigned_Analyst,
			&ticket.Opened_by,
			&ticket.Closed_by,
		)
		if err != nil {
			slog.Error("error scanning analysts tickets", "error message", err)
			return []Ticket{}
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}

func CreateTicket(ticket Ticket, team_id uuid.UUID, analyst_id uuid.UUID) (string, error) {
	var db *sql.DB = database.DB_Connection
	query := `insert into ticket values(gen_random_uuid(), default, $1, 'Open', $2, $3, $4, $5, $6, default, default, null, $7, $8, $9, null) returning ticket_id;`
	new_ticket_id := ""
	err := db.QueryRow(query,
		ticket.Type,
		ticket.Category,
		ticket.Subcategory,
		ticket.Title,
		ticket.Description,
		ticket.Customer_Contact,
		team_id,
		analyst_id,
		analyst_id,
	).Scan(&new_ticket_id)
	if err != nil {
		return "", err
	}

	query = `update analyst set number_of_open_tickets = number_of_open_tickets + 1, number_of_opened_tickets = number_of_opened_tickets + 1 where analyst_id = $1;`
	_, err = db.Exec(query, analyst_id)
	if err != nil {
		return "", err
	}

	return new_ticket_id, nil
}

func GetTicket(ticketID string) (Ticket, error) {
	var db *sql.DB = database.DB_Connection
	ticket := Ticket{}

	query := `select * from ticket where ticket_id = $1;`
	err := db.QueryRow(query, ticketID).Scan(
		&ticket.Ticket_ID,
		&ticket.Ticket_Number,
		&ticket.Type,
		&ticket.Status,
		&ticket.Category,
		&ticket.Subcategory,
		&ticket.Title,
		&ticket.Description,
		&ticket.Customer_Contact,
		&ticket.Opened_Date,
		&ticket.Updated_at,
		&ticket.Closed_Date,
		&ticket.Assigned_Team,
		&ticket.Assigned_Analyst,
		&ticket.Opened_by,
		&ticket.Closed_by,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Ticket{}, nil
		}
		return Ticket{}, err
	}

	return ticket, nil
}

func UpdateTicket(ticket Ticket) (string, error) {
	var db *sql.DB = database.DB_Connection
	query := `
		update ticket set
		type = $1,
		category_id = $2,
		subcategory_id = $3,
		title = $4,
		description = $5,
		customer_contact = $6,
		updated_at = current_timestamp
		where ticket_id = $7
		returning ticket_id;
	`
	new_ticket_id := ""
	err := db.QueryRow(query, ticket.Type, ticket.Category, ticket.Subcategory, ticket.Title, ticket.Description, ticket.Customer_Contact, ticket.Ticket_ID).Scan(&new_ticket_id)
	if err != nil {
		return "", err
	}

	return new_ticket_id, nil
}

func FilterTickets(search string, customer string, ticketType string, status string, category string, subcategory string, searchType string, teamID string) ([]Ticket, error) {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}
	var queryArgs []any

	search = strings.ToLower(search)
	customer = strings.ToLower(customer)
	search = "%" + search + "%"
	customer = "%" + customer + "%"

	queryArgs = append(queryArgs, search)
	queryArgs = append(queryArgs, customer)
	var queryArgIter int = 2

	query := `
		select *
		from ticket
		where (ticket_number::varchar like $1 or
		lower(title) like $1 or
		lower(description) like $1) and
		lower(customer_contact) like $2
	`

	if searchType == "Team Tickets" {
		query += `
			and assigned_team = $3
		`
		queryArgs = append(queryArgs, teamID)
		queryArgIter++
	} else if searchType == "Unassigned Tickets" {
		query += `
			and assigned_team = $3 and
			assigned_analyst is null
		`
		queryArgs = append(queryArgs, teamID)
		queryArgIter++
	}

	if category != "none" {
		query += `
			and category_id = $` + strconv.Itoa(queryArgIter+1) + `
		`
		queryArgs = append(queryArgs, category)
		queryArgIter++

		query += `
			and subcategory_id = $` + strconv.Itoa(queryArgIter+1) + `
		`
		queryArgs = append(queryArgs, subcategory)
		queryArgIter++
	}

	if ticketType != "Both" {
		query += `
			and type = $` + strconv.Itoa(queryArgIter+1) + `
		`
		queryArgs = append(queryArgs, ticketType)
		queryArgIter++
	}

	if status != "Both" {
		query += `
			and status = $` + strconv.Itoa(queryArgIter+1) + `
		`
		queryArgs = append(queryArgs, status)
		queryArgIter++
	}

	query += `
		order by opened_date;
	`

	rows, err := db.Query(query, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&ticket.Ticket_ID,
			&ticket.Ticket_Number,
			&ticket.Type,
			&ticket.Status,
			&ticket.Category,
			&ticket.Subcategory,
			&ticket.Title,
			&ticket.Description,
			&ticket.Customer_Contact,
			&ticket.Opened_Date,
			&ticket.Updated_at,
			&ticket.Closed_Date,
			&ticket.Assigned_Team,
			&ticket.Assigned_Analyst,
			&ticket.Opened_by,
			&ticket.Closed_by,
		)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func GetAllTickets() []Ticket {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}

	query := `select * from ticket order by opened_date;`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("error getting all tickets", "error message", err)
		return []Ticket{}
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&ticket.Ticket_ID,
			&ticket.Ticket_Number,
			&ticket.Type,
			&ticket.Status,
			&ticket.Category,
			&ticket.Subcategory,
			&ticket.Title,
			&ticket.Description,
			&ticket.Customer_Contact,
			&ticket.Opened_Date,
			&ticket.Updated_at,
			&ticket.Closed_Date,
			&ticket.Assigned_Team,
			&ticket.Assigned_Analyst,
			&ticket.Opened_by,
			&ticket.Closed_by,
		)
		if err != nil {
			slog.Error("error scanning all tickets", "error message", err)
			return []Ticket{}
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

func GetTeamsUnassignedTickets(teamID string) []Ticket {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}

	query := `select * from ticket where assigned_analyst is null and assigned_team = $1 order by opened_date;`
	rows, err := db.Query(query, teamID)
	if err != nil {
		slog.Error("error getting a team's unassigned tickets", "error message", err)
		return []Ticket{}
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&ticket.Ticket_ID,
			&ticket.Ticket_Number,
			&ticket.Type,
			&ticket.Status,
			&ticket.Category,
			&ticket.Subcategory,
			&ticket.Title,
			&ticket.Description,
			&ticket.Customer_Contact,
			&ticket.Opened_Date,
			&ticket.Updated_at,
			&ticket.Closed_Date,
			&ticket.Assigned_Team,
			&ticket.Assigned_Analyst,
			&ticket.Opened_by,
			&ticket.Closed_by,
		)
		if err != nil {
			slog.Error("error scanning a team's unassigned tickets", "error message", err)
			return []Ticket{}
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

func GetTeamTickets(teamID string) []Ticket {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}

	query := `select * from ticket where assigned_team = $1 order by opened_date;`
	rows, err := db.Query(query, teamID)
	if err != nil {
		slog.Error("error getting team's tickets", "error message", err)
		return []Ticket{}
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&ticket.Ticket_ID,
			&ticket.Ticket_Number,
			&ticket.Type,
			&ticket.Status,
			&ticket.Category,
			&ticket.Subcategory,
			&ticket.Title,
			&ticket.Description,
			&ticket.Customer_Contact,
			&ticket.Opened_Date,
			&ticket.Updated_at,
			&ticket.Closed_Date,
			&ticket.Assigned_Team,
			&ticket.Assigned_Analyst,
			&ticket.Opened_by,
			&ticket.Closed_by,
		)
		if err != nil {
			slog.Error("error scanning team's tickets", "error message", err)
			return []Ticket{}
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

func DeleteTicket(ticketID string) error {
	var db *sql.DB = database.DB_Connection
	query := `delete from ticket where ticket_id = $1 returning assigned_analyst, opened_by, status;`
	var assigned_analyst uuid.UUID
	var ticket_status string
	var opened_by uuid.UUID

	err := db.QueryRow(query, ticketID).Scan(&assigned_analyst, &opened_by, &ticket_status)
	if err != nil {
		return err
	}

	if assigned_analyst == opened_by {
		if ticket_status == "Open" {
			query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1, number_of_opened_tickets = number_of_opened_tickets - 1 where analyst_id = $1;`
			_, err = db.Exec(query, assigned_analyst)
			if err != nil {
				return err
			}
		} else {
			query = `update analyst set number_of_closed_tickets = number_of_closed_tickets - 1, number_of_opened_tickets = number_of_opened_tickets - 1 where analyst_id = $1;`
			_, err = db.Exec(query, assigned_analyst)
			if err != nil {
				return err
			}
		}
	} else {
		if ticket_status == "Open" {
			query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1 where analyst_id = $1;`
			_, err = db.Exec(query, assigned_analyst)
			if err != nil {
				return err
			}
		} else {
			query = `update analyst set number_of_closed_tickets = number_of_closed_tickets - 1 where analyst_id = $1;`
			_, err = db.Exec(query, assigned_analyst)
			if err != nil {
				return err
			}
		}
		query = `update analyst set number_of_opened_tickets = number_of_opened_tickets - 1 where analyst_id = $1;`
		_, err = db.Exec(query, opened_by)
		if err != nil {
			return err
		}
	}

	return nil
}

func TicketExists(ticketID string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	query := `select ticket_id from ticket where ticket_id = $1;`
	var returnedTicketID uuid.UUID
	err := db.QueryRow(query, ticketID).Scan(&returnedTicketID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func CloseTicket(ticket_id string, analyst_id string) (Ticket, error) {
	var db *sql.DB = database.DB_Connection
	ticket := Ticket{}
	query := `update ticket set status = 'Closed', closed_date = current_timestamp, closed_by = $1, updated_at = current_timestamp where ticket_id = $2 returning *;`

	err := db.QueryRow(query, analyst_id, ticket_id).Scan(
		&ticket.Ticket_ID,
		&ticket.Ticket_Number,
		&ticket.Type,
		&ticket.Status,
		&ticket.Category,
		&ticket.Subcategory,
		&ticket.Title,
		&ticket.Description,
		&ticket.Customer_Contact,
		&ticket.Opened_Date,
		&ticket.Updated_at,
		&ticket.Closed_Date,
		&ticket.Assigned_Team,
		&ticket.Assigned_Analyst,
		&ticket.Opened_by,
		&ticket.Closed_by,
	)
	if err != nil {
		return Ticket{}, err
	}

	if ticket.Assigned_Analyst.UUID.String() == analyst_id {
		query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1, number_of_closed_tickets = number_of_closed_tickets + 1 where analyst_id = $1;`
		_, err = db.Exec(query, analyst_id)
		if err != nil {
			return Ticket{}, err
		}
	} else {
		query = `update analyst set number_of_open_tickets = number_of_open_tickets - 1 where analyst_id = $1;`
		_, err = db.Exec(query, ticket.Assigned_Analyst.UUID)
		if err != nil {
			return Ticket{}, err
		}

		query = `update analyst set number_of_closed_tickets = number_of_closed_tickets + 1 where analyst_id = $1;`
		_, err = db.Exec(query, analyst_id)
		if err != nil {
			return Ticket{}, err
		}
	}

	return ticket, nil
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
