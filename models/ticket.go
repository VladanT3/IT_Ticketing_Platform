package models

import (
	"database/sql"
	"log"
	"strconv"
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
		log.Fatal("error getting tickets: ", err)
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
			log.Fatal("error scaning ticket: ", err)
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}

func GetTicket(ticketID string) Ticket {
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
			return Ticket{}
		}
		log.Fatal("error getting ticket: ", err)
	}

	return ticket
}

func FilterTickets(search string, customer string, ticketType string, status string, category string, subcategory string, searchType string, teamID string) []Ticket {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}
	var queryArgs []any

	search = "%" + search + "%"
	customer = "%" + customer + "%"

	queryArgs = append(queryArgs, search)
	queryArgs = append(queryArgs, customer)
	var queryArgIter int = 2

	query := `
		select *
		from ticket
		where (ticket_number::varchar like $1 or
		title like $1 or
		description like $1) and
		customer_contact like $2
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
		log.Fatal("error getting filtered tickets: ", err)
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
			log.Fatal("error scanning filtered tickets: ", err)
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

func GetAllTickets() []Ticket {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}

	query := `select * from ticket order by opened_date;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("error getting all tickets: ", err)
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
			log.Fatal("error scanning all tickets: ", err)
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
		log.Fatal("error getting unassigned tickets: ", err)
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
			log.Fatal("error scanning unassigned tickets: ", err)
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
		log.Fatal("error getting team tickets: ", err)
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
			log.Fatal("error scanning team tickets: ", err)
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

func DeleteTicket(ticketID string) {
	var db *sql.DB = database.DB_Connection
	query := `delete from ticket where ticket_id = $1;`

	_, err := db.Exec(query, ticketID)
	if err != nil {
		log.Fatal("error deleting ticket: ", err)
	}
}

func TicketExists(ticketID string) bool {
	var db *sql.DB = database.DB_Connection
	query := `select ticket_id from ticket where ticket_id = $1;`
	var returnedTicketID uuid.UUID
	err := db.QueryRow(query, ticketID).Scan(&returnedTicketID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal("error checking if ticket exists: ", err)
	}

	return true
}
