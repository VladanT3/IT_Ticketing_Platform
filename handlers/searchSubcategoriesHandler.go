package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/subcategories"
)

func SearchSubcategories(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("subcategory_search")
	category := r.FormValue("category")

	searchedSubcategories := models.SubcategorySearchByName(search, category)

	return Render(w, r, subcategories.SearchSubcategories(searchedSubcategories))
}
