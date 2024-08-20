package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
)

func SearchCategories(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("category_search")

	searchedCategories := models.CategorySearchByName(search)

	return Render(w, r, categories.SearchCategories(searchedCategories))
}
