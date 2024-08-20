package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
)

func ShowCategoriesPage(w http.ResponseWriter, r *http.Request) error {
	//TODO: check if user is admin
	return Render(w, r, categories.Categories(LoggedInUserType))
}

func ShowCategoryPopup(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, categories.CategoryPopup())
}

func CreateCategory(w http.ResponseWriter, r *http.Request) error {
	return nil
}
