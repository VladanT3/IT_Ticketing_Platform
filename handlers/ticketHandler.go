package handlers

import (
	"net/http"
)

func TicketHandler(w http.ResponseWriter, r *http.Request) error {
	ticketType := r.FormValue("type")
	category := r.FormValue("category")
	subcategory := r.FormValue("subcategory")
	title := r.FormValue("title")
	desc := r.FormValue("desc")
	customerContact := r.FormValue("customerContact")
	saveType := r.FormValue("save")

	session, _ := Store.Get(r, "ticket")
	session.Values["ticketType"] = ticketType
	session.Values["category"] = category
	session.Values["subcategory"] = subcategory
	session.Values["title"] = title
	session.Values["desc"] = desc
	session.Values["customerContact"] = customerContact
	session.Values["saveType"] = saveType
	session.Save(r, w)

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
