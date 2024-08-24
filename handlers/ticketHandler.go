package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func TicketRedirection(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:    "ticketType",
		Value:   r.FormValue("ticketType"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "category",
		Value:   r.FormValue("category"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "subcategory",
		Value:   r.FormValue("subcategory"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "title",
		Value:   r.FormValue("title"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "desc",
		Value:   r.FormValue("desc"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "customerContact",
		Value:   r.FormValue("customerContact"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "saveType",
		Value:   r.FormValue("saveType"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "ticketID",
		Value:   r.FormValue("ticketID"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})

	mode := r.FormValue("mode")
	if mode == "create" {
		http.Redirect(w, r, "/ticket/create", http.StatusSeeOther)
		return nil
	} else if mode == "update" {
		http.Redirect(w, r, "/ticket/update", http.StatusSeeOther)
		return nil
	}

	return nil
}

func CreateTicket(w http.ResponseWriter, r *http.Request) error {
	typeCookie, err := r.Cookie("ticketType")
	if err != nil {
		log.Fatal("No cookie with name 'ticketType': ", err)
	}
	ticketType := typeCookie.Value

	categoryCookie, err := r.Cookie("category")
	if err != nil {
		log.Fatal("No cookie with name 'category': ", err)
	}
	category := categoryCookie.Value

	subcategoryCookie, err := r.Cookie("subcategory")
	if err != nil {
		log.Fatal("No cookie with name 'subcategory': ", err)
	}
	subcategory := subcategoryCookie.Value

	titleCookie, err := r.Cookie("title")
	if err != nil {
		log.Fatal("No cookie with name 'title': ", err)
	}
	title := titleCookie.Value

	descCookie, err := r.Cookie("desc")
	if err != nil {
		log.Fatal("No cookie with name 'desc': ", err)
	}
	desc := descCookie.Value

	customerContactCookie, err := r.Cookie("customerContact")
	if err != nil {
		log.Fatal("No cookie with name 'customerContact': ", err)
	}
	customerContact := customerContactCookie.Value

	saveTypeCookie, err := r.Cookie("saveType")
	if err != nil {
		log.Fatal("No cookie with name 'saveType': ", err)
	}
	saveType := saveTypeCookie.Value

	teamID := models.GetAnalystsTeam(LoggedInUser.Analyst_ID.String()).Team_ID

	if category == "none" || category == "" {
		newTicket := models.Ticket{
			Type: ticketType,
			Category: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
			Subcategory: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
			Title:            title,
			Description:      desc,
			Customer_Contact: customerContact,
		}

		return Render(w, r, tickets.TicketForm(models.Ticket{}, LoggedInUser, LoggedInUserType, "create", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
	}

	var db *sql.DB = database.DB_Connection
	query := `insert into ticket values(gen_random_uuid(), default, $1, 'Open', $2, $3, $4, $5, $6, default, default, null, $7, $8, $9, null) returning ticket_id;`
	newTicketID := ""
	err = db.QueryRow(query, ticketType, category, subcategory, title, desc, customerContact, teamID, LoggedInUser.Analyst_ID, LoggedInUser.Analyst_ID).Scan(&newTicketID)
	if err != nil {
		errMsg := "error inserting ticket: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	query = `update analyst set number_of_open_tickets = number_of_open_tickets + 1, number_of_opened_tickets = number_of_opened_tickets + 1 where analyst_id = $1;`
	_, err = db.Exec(query, LoggedInUser.Analyst_ID)
	if err != nil {
		errMsg := "error updating ticket number trackers: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	LoggedInUser = models.UpdateLoggedInUser(LoggedInUser)

	if saveType == "Save" {
		http.Redirect(w, r, "/ticket/"+newTicketID, http.StatusSeeOther)
		return nil
	} else if saveType == "Save and Exit" {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	return nil
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) error {
	typeCookie, err := r.Cookie("ticketType")
	if err != nil {
		errMsg := "No cookie with name 'ticketType': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	ticketType := typeCookie.Value

	categoryCookie, err := r.Cookie("category")
	if err != nil {
		errMsg := "No cookie with name 'category': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	category := categoryCookie.Value

	subcategoryCookie, err := r.Cookie("subcategory")
	if err != nil {
		errMsg := "No cookie with name 'subcategory': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	subcategory := subcategoryCookie.Value

	titleCookie, err := r.Cookie("title")
	if err != nil {
		errMsg := "No cookie with name 'title': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	title := titleCookie.Value

	descCookie, err := r.Cookie("desc")
	if err != nil {
		errMsg := "No cookie with name 'desc': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	desc := descCookie.Value

	customerContactCookie, err := r.Cookie("customerContact")
	if err != nil {
		errMsg := "No cookie with name 'customerContact': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	customerContact := customerContactCookie.Value

	saveTypeCookie, err := r.Cookie("saveType")
	if err != nil {
		errMsg := "No cookie with name 'saveType': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	saveType := saveTypeCookie.Value

	ticketIDCookie, err := r.Cookie("ticketID")
	if err != nil {
		errMsg := "No cookie with name 'ticketID': " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	ticketID := ticketIDCookie.Value

	if category == "none" || category == "" {
		newTicket := models.Ticket{
			Type: ticketType,
			Category: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
			Subcategory: uuid.NullUUID{
				UUID:  uuid.Nil,
				Valid: false,
			},
			Title:            title,
			Description:      desc,
			Customer_Contact: customerContact,
		}

		return Render(w, r, tickets.TicketForm(models.GetTicket(ticketID), LoggedInUser, LoggedInUserType, "update", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
	}

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
	newTicketID := ""
	err = db.QueryRow(query, ticketType, category, subcategory, title, desc, customerContact, ticketID).Scan(&newTicketID)
	if err != nil {
		errMsg := "error updating ticket: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	if saveType == "Save" {
		http.Redirect(w, r, "/ticket/"+newTicketID, http.StatusSeeOther)
		return nil
	} else if saveType == "Save and Exit" {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	return nil
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) error {
	ticketID := chi.URLParam(r, "ticketID")

	models.DeleteTicket(ticketID)
	LoggedInUser = models.UpdateLoggedInUser(LoggedInUser)

	return Render(w, r, tickets.DeletedTicket())
}

func ShowNewTicketForm(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, tickets.TicketForm(models.Ticket{}, LoggedInUser, LoggedInUserType, "create", "", "", models.Ticket{}))
}

func ShowTicket(w http.ResponseWriter, r *http.Request) error {
	ticketID := chi.URLParam(r, "ticketID")

	if models.TicketExists(ticketID) {
		ticketToShow := models.GetTicket(ticketID)
		return Render(w, r, tickets.TicketForm(ticketToShow, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
	} else {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "No ticket found: ticket doesn't exist or it was removed!"))
	}
}

func ShowAllTicketSearch(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "All Ticket Search"))
}

func ShowTeamTickets(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "manager" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of managerial credentials!"))
	}

	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "Team Tickets"))
}

func ShowUnassignedTickets(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "Unassigned Tickets"))
}

func FilterTickets(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("search")
	customer := r.FormValue("customer")
	ticketType := r.FormValue("type")
	status := r.FormValue("status")
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	searchType := r.FormValue("searchType")

	fileteredTickets := models.FilterTickets(search, customer, ticketType, status, category, subcategory, searchType, LoggedInUser.Team_ID.UUID.String())

	return Render(w, r, tickets.Tickets(fileteredTickets))
}

func CloseTicket(w http.ResponseWriter, r *http.Request) error {
	ticket_id := chi.URLParam(r, "ticketID")

	ticket_to_show := models.CloseTicket(ticket_id, LoggedInUser.Analyst_ID.String())
	LoggedInUser = models.UpdateLoggedInUser(LoggedInUser)

	w.Header().Add("HX-Redirect", "/ticket/"+ticket_id)
	return Render(w, r, tickets.TicketForm(ticket_to_show, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
}

func ShowTicketReopenForm(w http.ResponseWriter, r *http.Request) error {
	ticket_id := chi.URLParam(r, "ticketID")

	ticket := models.GetTicket(ticket_id)

	return Render(w, r, tickets.ReopenForm(LoggedInUserType, ticket))
}

func ReopenTicket(w http.ResponseWriter, r *http.Request) error {
	reason := r.FormValue("reason")
	ticket_id := chi.URLParam(r, "ticketID")

	models.ReopenTicket(ticket_id, reason, LoggedInUser.Analyst_ID.String())
	LoggedInUser = models.UpdateLoggedInUser(LoggedInUser)

	ticket := models.GetTicket(ticket_id)
	return Render(w, r, tickets.TicketForm(ticket, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
}

func ShowTicketReopenHistory(w http.ResponseWriter, r *http.Request) error {
	ticket_id := chi.URLParam(r, "ticketID")

	reopens := models.GetTicketReopens(ticket_id)

	return Render(w, r, tickets.ReopenHistory(LoggedInUserType, reopens))
}
