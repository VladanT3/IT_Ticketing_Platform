package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/go-chi/chi/v5"
)

func ShowCategoriesPage(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, categories.Categories(LoggedInUserType))
}

func SearchCategories(w http.ResponseWriter, r *http.Request) error {
	search := r.FormValue("category_search")

	searchedCategories, err := models.CategorySearchByName(search)
	if err != nil {
		err_msg := "Internal server error:\nerror searching for categories: " + err.Error()
		w.Header().Add("ErrorMessage", err_msg)
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, categories.ShowCategories(searchedCategories, false))
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
	category_name := r.FormValue("category_name")

	category_name_exists, err := models.DoesCategoryNameExist(category_name)
	if err != nil {
		err_msg := "Internal server error:\nerror checking if category name exists: " + err.Error()
		w.Header().Add("ErrorMessage", err_msg)
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if category_name_exists {
		return Render(w, r, categories.ShowCategories(models.GetAllCategories(), true))
	}

	err = models.CreateCategory(category_name)
	if err != nil {
		err_msg := "Internal server error:\nerror creating category: " + err.Error()
		w.Header().Add("ErrorMessage", err_msg)
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, categories.ShowCategories(models.GetAllCategories(), false))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) error {
	category_id := chi.URLParam(r, "category_id")
	category_name := r.FormValue("category_name")

	category_name_exists, err := models.DoesCategoryNameExist(category_name)
	if err != nil {
		err_msg := "Internal server error:\nerror checking if category name exists: " + err.Error()
		w.Header().Add("ErrorMessage", err_msg)
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if category_name_exists {
		return Render(w, r, categories.ShowCategories(models.GetAllCategories(), true))
	} else {
		err = models.UpdateCategory(category_id, category_name)
		if err != nil {
			err_msg := "Internal server error:\nerror updating category: " + err.Error()
			w.Header().Add("ErrorMessage", err_msg)
			w.Header().Add("HX-Redirect", "/error")
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}

		return Render(w, r, categories.ShowCategories(models.GetAllCategories(), false))
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) error {
	category_id := chi.URLParam(r, "category_id")

	err := models.DeleteCategory(category_id)
	if err != nil {
		err_msg := "Internal server error:\nerror deleting category: " + err.Error()
		w.Header().Add("ErrorMessage", err_msg)
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, categories.ShowCategories(models.GetAllCategories(), false))
}

func ShowCategoryAlreadyExistsError(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, categories.CategoryExistsError())
}
