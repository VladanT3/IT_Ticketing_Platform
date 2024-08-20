package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
	"github.com/go-chi/chi/v5"
)

func ShowCategoriesPage(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, categories.Categories(LoggedInUserType))
}

func SearchCategories(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("category_search")

	searchedCategories := models.CategorySearchByName(search)

	return Render(w, r, categories.ShowCategories(searchedCategories))
}

func ShowCategoryPopup(w http.ResponseWriter, r *http.Request) error {
	operation := r.FormValue("category_operation")
	var category_id string
	var category_name string

	if operation != "create" {
		category_id = r.FormValue("id")
		category_name = r.FormValue("name")
	}

	return Render(w, r, categories.CategoryPopup(operation, category_id, category_name))
}

func CreateCategory(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("category_name")

	models.CreateCategory(name)

	return Render(w, r, categories.ShowCategories(models.GetAllCategories()))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) error {
	category_id := chi.URLParam(r, "categoryID")
	category_name := r.FormValue("category_name")

	models.UpdateCategory(category_id, category_name)

	return Render(w, r, categories.ShowCategories(models.GetAllCategories()))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) error {
	category_id := chi.URLParam(r, "categoryID")

	models.DeleteCategory(category_id)

	return Render(w, r, categories.ShowCategories(models.GetAllCategories()))
}
