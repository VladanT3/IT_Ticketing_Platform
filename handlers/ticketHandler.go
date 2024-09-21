package handlers

import (
	"net/http"
	"time"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
	"github.com/VladanT3/IT_Ticketing_Platform/views/tickets"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func ShowNewTicketForm(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, tickets.TicketForm(models.Ticket{}, LoggedInUser, LoggedInUserType, "create", "", "", models.Ticket{}))
}

func ShowTicket(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	ticket_id := chi.URLParam(r, "ticket_id")
	exists, err := models.TicketExists(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror checking if ticket exists: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if exists {
		ticket_to_show, err := models.GetTicket(ticket_id)
		if err != nil {
			err_msg := "Internal server error:\nerror getting ticket to show: " + err.Error()
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}
		return Render(w, r, tickets.TicketForm(ticket_to_show, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
	} else {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "No ticket found: ticket doesn't exist or it was removed!"))
	}
}

func TicketRedirection(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "ticket_type",
		Value:   r.FormValue("ticket_type"),
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
		Value:   r.FormValue("customer_contact"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "save_type",
		Value:   r.FormValue("save_type"),
		Expires: time.Time.Add(time.Now(), time.Second*10),
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "ticket_id",
		Value:   r.FormValue("ticket_id"),
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
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	type_cookie, err := r.Cookie("ticket_type")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	ticket_type := type_cookie.Value

	category_cookie, err := r.Cookie("category")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'category': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	category := category_cookie.Value

	subcategory_cookie, err := r.Cookie("subcategory")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'subcategory': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	subcategory := subcategory_cookie.Value

	title_cookie, err := r.Cookie("title")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'title': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	title := title_cookie.Value

	desc_cookie, err := r.Cookie("desc")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'desc': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	desc := desc_cookie.Value

	customer_contact_cookie, err := r.Cookie("customer_contact")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'customer_contact': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	customer_contact := customer_contact_cookie.Value

	save_type_cookie, err := r.Cookie("save_type")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'save_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
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
		err_msg := "Internal server error:\nerror creating ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user statistics after creating ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
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
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	type_cookie, err := r.Cookie("ticket_type")
	if err != nil {
		err_msg := "Internal server errpr:\nno cookie with name 'ticket_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	ticket_type := type_cookie.Value

	category_cookie, err := r.Cookie("category")
	if err != nil {
		err_msg := "Internal server error: \nno cookie with name 'category': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	category := category_cookie.Value

	subcategory_cookie, err := r.Cookie("subcategory")
	if err != nil {
		err_msg := "Internal server error: \nno cookie with name 'subcategory': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	subcategory := subcategory_cookie.Value

	title_cookie, err := r.Cookie("title")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'title': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	title := title_cookie.Value

	desc_cookie, err := r.Cookie("desc")
	if err != nil {
		err_msg := "Internal server error: \nno cookie with name 'desc': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	desc := desc_cookie.Value

	customer_contact_cookie, err := r.Cookie("customer_contact")
	if err != nil {
		err_msg := "Internal server error: \nno cookie with name 'customer_contact': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	customer_contact := customer_contact_cookie.Value

	save_type_cookie, err := r.Cookie("save_type")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'save_type': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	save_type := save_type_cookie.Value

	ticket_id_cookie, err := r.Cookie("ticket_id")
	if err != nil {
		err_msg := "Internal server error:\nno cookie with name 'ticket_id': " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}
	ticket_id := ticket_id_cookie.Value

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

		ticket, err := models.GetTicket(ticket_id)
		if err != nil {
			err_msg := "Internal server error:\nerror getting ticket for updating: " + err.Error()
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}
		return Render(w, r, tickets.TicketForm(ticket, LoggedInUser, LoggedInUserType, "update", "Please select a valid category!", "Please select a valid subcategory!", new_ticket))
	}

	new_ticket := models.Ticket{
		Ticket_ID: uuid.MustParse(ticket_id),
		Type:      ticket_type,
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

	new_ticket_id, err := models.UpdateTicket(new_ticket)
	if err != nil {
		err_msg := "Internal server error:\nerror updating ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
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

func DeleteTicket(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "manager" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of managerial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	err := models.DeleteTicket(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror deleting ticket: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user statistics after deleting ticket: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.DeletedTicket())
}

func ShowAllTicketSearch(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "All Ticket Search"))
}

func ShowTeamTickets(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "manager" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of managerial credentials!"))
	}

	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "Team Tickets"))
}

func ShowUnassignedTickets(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	return Render(w, r, tickets.TicketSearch(LoggedInUser, LoggedInUserType, "Unassigned Tickets"))
}

func FilterTickets(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	search := r.FormValue("search")
	customer := r.FormValue("customer")
	ticket_type := r.FormValue("type")
	status := r.FormValue("status")
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	search_type := r.FormValue("search_type")

	filetered_tickets, err := models.FilterTickets(search, customer, ticket_type, status, category, subcategory, search_type, LoggedInUser.Team_ID.UUID.String())
	if err != nil {
		err_msg := "Internal server error:\nerror filtering tickets: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.Tickets(filetered_tickets))
}

func CloseTicket(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	ticket_to_show, err := models.CloseTicket(ticket_id, LoggedInUser.Analyst_ID.String())
	if err != nil {
		err_msg := "Internal server error:\nerror closing ticket: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user statistics after closing ticket: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	w.Header().Add("HX-Redirect", "/ticket/"+ticket_id)
	return Render(w, r, tickets.TicketForm(ticket_to_show, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
}

func ShowTicketReopenForm(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	ticket, err := models.GetTicket(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting ticket for seeing reopens: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.ReopenForm(LoggedInUserType, ticket))
}

func ReopenTicket(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	reason := r.FormValue("reason")
	ticket_id := chi.URLParam(r, "ticket_id")

	err := models.ReopenTicket(ticket_id, reason, LoggedInUser.Analyst_ID.String())
	if err != nil {
		err_msg := "Internal server error:\nerror reopening ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user statistics after reopening ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	ticket, err := models.GetTicket(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting ticket after reopening: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.TicketForm(ticket, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
}

func ShowTicketReopenHistory(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	reopens, err := models.GetTicketReopens(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting all ticket reopens: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.ReopenHistory(LoggedInUserType, reopens, ticket_id))
}

func ShowTicketAssignmentForm(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	ticket, err := models.GetTicket(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting ticket for assignment: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.AssignmentForm(LoggedInUserType, ticket, false, "", ""))
}

func AssignTicket(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	ticket_id := chi.URLParam(r, "ticket_id")
	assign_to_analyst := r.FormValue("analyst")
	assign_to_team := r.FormValue("team")
	message := r.FormValue("message")

	if assign_to_team == "none" {
		ticket, err := models.GetTicket(ticket_id)
		if err != nil {
			err_msg := "Internal server error:\nerror getting ticket for assignment: " + err.Error()
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}
		return Render(w, r, tickets.AssignmentForm(LoggedInUserType, ticket, true, assign_to_team, assign_to_analyst))
	}

	err := models.AssignTicket(ticket_id, LoggedInUser.Analyst_ID.String(), assign_to_analyst, assign_to_team, message)
	if err != nil {
		err_msg := "Internal server error:\nerror assigning ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user statistics after assigning ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	ticket, err := models.GetTicket(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting ticket for assignment: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.TicketForm(ticket, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
}

func ShowTicketAssignmentHistory(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	ticket_assignments, err := models.GetAllTicketsAssignments(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting all ticket assignments: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, tickets.AssignmentHistory(LoggedInUserType, ticket_id, ticket_assignments))
}

func AssignTicketToMe(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	ticket_id := chi.URLParam(r, "ticket_id")

	err := models.AssignTicket(ticket_id, LoggedInUser.Analyst_ID.String(), LoggedInUser.Analyst_ID.String(), LoggedInUser.Team_ID.UUID.String(), "Assigned to self.")
	if err != nil {
		err_msg := "Internal server error:\nerror assigning ticket to self: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	LoggedInUser, err = models.UpdateLoggedInUser(LoggedInUser)
	if err != nil {
		err_msg := "Internal server error:\nerror updating user statistics after assigning ticket: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	ticket, err := models.GetTicket(ticket_id)
	if err != nil {
		err_msg := "Internal server error:\nerror getting ticket for assignment: " + err.Error()
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	w.Header().Add("HX-Redirect", "/ticket/"+ticket_id)

	return Render(w, r, tickets.TicketForm(ticket, LoggedInUser, LoggedInUserType, "update", "", "", models.Ticket{}))
}
