package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
)

func SelectSubcategories(w http.ResponseWriter, r *http.Request) error {
	category := r.FormValue("category")
	if category == "none" || category == "" {
		return Render(w, r, layouts.SelectSubcategories([]models.Subcategory{}))
	}
	subcategories := models.GetSubcategories(category)

	return Render(w, r, layouts.SelectSubcategories(subcategories))
}

func EditingSubcategories(w http.ResponseWriter, r *http.Request) error {
	category := r.FormValue("category")

	if category == "" {
		if LoggedInUserType == "admin" {
			http.Redirect(w, r, "/categories", http.StatusSeeOther)
			return nil
		} else {
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administrational credentials!"))
		}
	}
	subcategories := models.GetSubcategories(category)

	return Render(w, r, layouts.EditingSubcategories(subcategories))
}
