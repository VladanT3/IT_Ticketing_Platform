package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/subcategories"
)

func SelectSubcategories(w http.ResponseWriter, r *http.Request) error {
	category := r.FormValue("category")
	if category == "none" || category == "" {
		return Render(w, r, subcategories.SelectSubcategories([]models.Subcategory{}))
	}
	subcategoryOutput := models.GetSubcategories(category)

	return Render(w, r, subcategories.SelectSubcategories(subcategoryOutput))
}

func SearchSubcategories(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("subcategory_search")
	category := r.FormValue("category")

	searchedSubcategories := models.SubcategorySearchByName(search, category)

	return Render(w, r, subcategories.SearchSubcategories(searchedSubcategories))
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
	subcategoryOutput := models.GetSubcategories(category)

	return Render(w, r, subcategories.EditingSubcategories(subcategoryOutput, category))
}
