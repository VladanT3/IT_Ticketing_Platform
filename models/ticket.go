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
	var query string
	var offset int
	search = "%" + search + "%"
	customer = "%" + customer + "%"

	if searchType == "All Ticket Search" {
		query = `
			select *
			from ticket
			where (ticket_number::varchar like $1 or
			title like $1 or
			description like $1) and
			customer_contact like $2
		`
		offset = 0
	} else if searchType == "Team Tickets" {
		query = `
			select *
			from ticket
			where (ticket_number::varchar like $1 or
			title like $1 or
			description like $1) and
			customer_contact like $2 and
			assigned_team = $3
		`
		offset = 1
	} else if searchType == "Unassigned Tickets" {
		query = `
			select *
			from ticket
			where (ticket_number::varchar like $1 or
			title like $1 or
			description like $1) and
			customer_contact like $2 and
			assigned_team = $3 and
			assigned_analyst is null
		`
		offset = 1
	}

	var rows *sql.Rows
	var err error

	if category == "none" {
		if ticketType == "Both" {
			if status == "Both" {
				query += `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer)
				} else {
					rows, err = db.Query(query, search, customer, teamID)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and status = $` + strconv.Itoa(3+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, status)
				} else {
					rows, err = db.Query(query, search, customer, teamID, status)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			}
		} else {
			if status == "Both" {
				query += `
					and type = $` + strconv.Itoa(3+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, ticketType)
				} else {
					rows, err = db.Query(query, search, customer, teamID, ticketType)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and type = $` + strconv.Itoa(3+offset) + ` and
					status = $` + strconv.Itoa(4+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, ticketType, status)
				} else {
					rows, err = db.Query(query, search, customer, teamID, ticketType, status)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			}
		}
	} else {
		if ticketType == "Both" {
			if status == "Both" {
				query += `
					and category_id = $` + strconv.Itoa(3+offset) + ` and
					subcategory_id = $` + strconv.Itoa(4+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, category, subcategory)
				} else {
					rows, err = db.Query(query, search, customer, teamID, category, subcategory)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and status = $` + strconv.Itoa(3+offset) + ` and
					category_id = $` + strconv.Itoa(4+offset) + ` and
					subcategory_id = $` + strconv.Itoa(5+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, status, category, subcategory)
				} else {
					rows, err = db.Query(query, search, customer, teamID, status, category, subcategory)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			}
		} else {
			if status == "Both" {
				query += `
					and type = $` + strconv.Itoa(3+offset) + ` and
					category_id = $` + strconv.Itoa(4+offset) + ` and
					subcategory_id = $` + strconv.Itoa(5+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, ticketType, category, subcategory)
				} else {
					rows, err = db.Query(query, search, customer, teamID, ticketType, category, subcategory)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and type = $` + strconv.Itoa(3+offset) + ` and
					status = $` + strconv.Itoa(4+offset) + ` and
					category_id = $` + strconv.Itoa(5+offset) + ` and
					subcategory_id = $` + strconv.Itoa(6+offset) + `
					order by opened_date;
				`

				if offset == 0 {
					rows, err = db.Query(query, search, customer, ticketType, status, category, subcategory)
				} else {
					rows, err = db.Query(query, search, customer, teamID, ticketType, status, category, subcategory)
				}
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			}
		}
	}

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
