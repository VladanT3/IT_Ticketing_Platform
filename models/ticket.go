package models

import (
	"database/sql"
	"log"
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

func FilterTickets(search string, customer string, ticketType string, status string, category string, subcategory string) []Ticket {
	var db *sql.DB = database.DB_Connection
	tickets := []Ticket{}
	ticket := Ticket{}
	var query string = `
		select *
		from ticket
		where (ticket_number::varchar like '%` + search + `%' or
		title like '%` + search + `%' or
		description like '%` + search + `%') and
		customer_contact like '%` + customer + `%'
	`
	var rows *sql.Rows
	var err error

	if category == "none" {
		if ticketType == "Both" {
			if status == "Both" {
				query += `
					order by opened_date;
				`
				rows, err = db.Query(query)
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and status = $1
					order by opened_date;
				`
				rows, err = db.Query(query, status)
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			}
		} else {
			if status == "Both" {
				query += `
					and type = $1
					order by opened_date;
				`
				rows, err = db.Query(query, ticketType)
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and type = $1 and
					status = $2
					order by opened_date;
				`
				rows, err = db.Query(query, ticketType, status)
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
					and category_id = $1 and
					subcategory_id = $2
					order by opened_date;
				`
				rows, err = db.Query(query, category, subcategory)
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and status = $1 and
					category_id = $2 and
					subcategory_id = $3
					order by opened_date;
				`
				rows, err = db.Query(query, status, category, subcategory)
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			}
		} else {
			if status == "Both" {
				query += `
					and type = $1 and
					category_id = $2 and
					subcategory_id = $3
					order by opened_date;
				`
				rows, err = db.Query(query, ticketType, category, subcategory)
				if err != nil {
					log.Fatal("error getting filtered tickets: ", err)
				}
				defer rows.Close()
			} else {
				query += `
					and type = $1 and
					status = $2 and
					category_id = $3 and
					subcategory_id = $4
					order by opened_date;
				`
				rows, err = db.Query(query, ticketType, status, category, subcategory)
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
