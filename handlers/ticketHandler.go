package handlers

import (
	"database/sql"
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
		Name:    "ticket_type",
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
		Name:    "customer_contact",
		Value:   r.FormValue("customerContact"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "save_type",
		Value:   r.FormValue("saveType"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "ticket_id",
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
	type_cookie, err := r.Cookie("ticket_type")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	ticket_type := type_cookie.Value

	category_cookie, err := r.Cookie("category")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'category': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	category := category_cookie.Value

	subcategory_cookie, err := r.Cookie("subcategory")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'subcategory': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	subcategory := subcategory_cookie.Value

	title_cookie, err := r.Cookie("title")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'title': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	title := title_cookie.Value

	desc_cookie, err := r.Cookie("desc")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'desc': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	desc := desc_cookie.Value

	customer_contact_cookie, err := r.Cookie("customer_contact")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'customer_contact': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	customer_contact := customer_contact_cookie.Value

	save_type_cookie, err := r.Cookie("save_type")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'save_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	save_type := save_type_cookie.Value

	team_id := models.GetAnalystsTeam(LoggedInUser.Analyst_ID.String()).Team_ID

	if category == "none" || category == "" {
		new_ticket := models.Ticket{
			Type: ticket_type,
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
			Customer_Contact: customer_contact,
		}

		return Render(w, r, tickets.TicketForm(models.Ticket{}, LoggedInUser, LoggedInUserType, "create", "Please select a valid category!", "Please select a valid subcategory!", new_ticket))
	}

	new_ticket := models.Ticket{
		Type: ticket_type,
		Category: uuid.NullUUID{
			UUID:  uuid.MustParse(category),
			Valid: true,
		},
		Subcategory: uuid.NullUUID{
			UUID:  uuid.MustParse(subcategory),
			Valid: true,
		},
		Title:            title,
		Description:      desc,
		Customer_Contact: customer_contact,
	}

	new_ticket_id, err := models.CreateTicket(new_ticket, team_id, LoggedInUser.Analyst_ID)
	if err != nil {
		errMsg := "Internal server error:\nerror creating ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		errMsg := "Internal server error:\nerror updating user statistics after creating ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage("", errMsg))
	}

	if save_type == "Save" {
		http.Redirect(w, r, "/ticket/"+new_ticket_id, http.StatusSeeOther)
		return nil
	} else if save_type == "Save and Exit" {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return nil
	}

	return nil
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) error {
	type_cookie, err := r.Cookie("ticket_type")
	if err != nil {
		errMsg := "Internal server errpr:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	ticket_type := type_cookie.Value

	category_cookie, err := r.Cookie("category")
	if err != nil {
		errMsg := "Internal server error: \nno cookie with name 'category': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	category := category_cookie.Value

	subcategory_cookie, err := r.Cookie("subcategory")
	if err != nil {
		errMsg := "Internal server error: \nno cookie with name 'subcategory': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	subcategory := subcategory_cookie.Value

	title_cookie, err := r.Cookie("title")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'title': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	title := title_cookie.Value

	desc_cookie, err := r.Cookie("desc")
	if err != nil {
		errMsg := "Internal server error: \nno cookie with name 'desc': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	desc := desc_cookie.Value

	customer_contact_cookie, err := r.Cookie("customer_contact")
	if err != nil {
		errMsg := "Internal server error: \nno cookie with name 'customer_contact': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	customer_contact := customer_contact_cookie.Value

	save_type_cookie, err := r.Cookie("save_type")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'save_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	save_type := save_type_cookie.Value

	ticket_id_cookie, err := r.Cookie("ticket_id")
	if err != nil {
		errMsg := "Internal server error:\nno cookie with name 'ticket_id': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, errMsg))
	}
	ticket_id := ticket_id_cookie.Value

	if category == "none" || category == "" {
		newTicket := models.Ticket{
			Type: ticket_type,
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
			Customer_Contact: customer_contact,
		}

		return Render(w, r, tickets.TicketForm(models.GetTicket(ticket_id), LoggedInUser, LoggedInUserType, "update", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
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
	err = db.QueryRow(query, ticket_type, category, subcategory, title, desc, customer_contact, ticket_id).Scan(&newTicketID)
	if err != nil {
		errMsg := "error updating ticket: " + err.Error()
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	if save_type == "Save" {
		http.Redirect(w, r, "/ticket/"+newTicketID, http.StatusSeeOther)
		return nil
	} else if save_type == "Save and Exit" {
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

	return Render(w, r, tickets.ReopenHistory(LoggedInUserType, reopens, ticket_id))
}
