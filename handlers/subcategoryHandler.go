package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
)

func GetSubcategories(w http.ResponseWriter, r *http.Request) error {
	category := r.FormValue("category")
	if category == "none" || category == "" {
		return Render(w, r, layouts.Subcategories([]models.Subcategory{}))
	}
	subcategories := models.GetSubcategories(category)

	return Render(w, r, layouts.Subcategories(subcategories))
}
