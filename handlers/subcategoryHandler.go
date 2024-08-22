package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/subcategories"
	"github.com/go-chi/chi/v5"
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

func ShowModifiableSubcategories(w http.ResponseWriter, r *http.Request) error {
	category_id := r.FormValue("category")

	if category_id == "" {
		if LoggedInUserType == "admin" {
			w.Header().Add("HX-Redirect", "/categories")
			return Render(w, r, categories.Categories(LoggedInUserType))
		} else {
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administrational credentials!"))
		}
	}
	subcategoryOutput := models.GetSubcategories(category_id)

	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func ShowSubcategoryPopup(w http.ResponseWriter, r *http.Request) error {
	operation := r.FormValue("subcategory_operation")
	category_id := r.FormValue("category_id")
	var subcategory_id string
	var subcategory_name string

	if operation != "create" {
		subcategory_id = r.FormValue("subcategory_id")
		subcategory_name = r.FormValue("name")
	}

	return Render(w, r, subcategories.SubcategoryPopup(operation, category_id, subcategory_id, subcategory_name))
}

func CreateSubcategory(w http.ResponseWriter, r *http.Request) error {
	subcategory_name := r.FormValue("subcategory_name")
	category_id := r.FormValue("category_id")

	if models.DoesSubcategoryNameExist(subcategory_name, category_id) {
		subcategoryOutput := models.GetSubcategories(category_id)
		return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, true))
	}
	models.CreateSubcategory(subcategory_name, category_id)

	subcategoryOutput := models.GetSubcategories(category_id)
	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func UpdateSubcategory(w http.ResponseWriter, r *http.Request) error {
	subcategory_id := chi.URLParam(r, "subcategoryID")
	subcategory_name := r.FormValue("subcategory_name")
	category_id := r.FormValue("category_id")

	if models.IsSubcategoryNameNew(subcategory_id, category_id, subcategory_name) {
		if models.DoesSubcategoryNameExist(subcategory_name, category_id) {
			subcategoryOutput := models.GetSubcategories(category_id)
			return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, true))
		} else {
			models.UpdateSubcategory(subcategory_id, subcategory_name, category_id)

			subcategoryOutput := models.GetSubcategories(category_id)
			return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
		}
	} else {
		subcategoryOutput := models.GetSubcategories(category_id)
		return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
	}
}

func DeleteSubcategory(w http.ResponseWriter, r *http.Request) error {
	subcategory_id := chi.URLParam(r, "subcategoryID")
	category_id := r.FormValue("category_id")

	models.DeleteSubcategory(subcategory_id, category_id)

	subcategoryOutput := models.GetSubcategories(category_id)
	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func ShowSubcategoryAlreadyExistsError(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, subcategories.SubcategoryExistsError())
}
