package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
)

func ShowCategoriesPage(w http.ResponseWriter, r *http.Request) error {
	//TODO: check if user is admin
	return Render(w, r, categories.Categories(LoggedInUserType))
}

func SearchCategories(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("category_search")

	searchedCategories := models.CategorySearchByName(search)

	return Render(w, r, categories.SearchCategories(searchedCategories))
}

func ShowCategoryPopup(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, categories.CategoryPopup())
}

func CreateCategory(w http.ResponseWriter, r *http.Request) error {
	return nil
}
