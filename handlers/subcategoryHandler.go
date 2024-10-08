package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/VladanT3/IT_Ticketing_Platform/views/login"
	"github.com/VladanT3/IT_Ticketing_Platform/views/subcategories"
	"github.com/go-chi/chi/v5"
)

func SelectSubcategories(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	category := r.FormValue("category")
	if category == "none" || category == "" {
		return Render(w, r, subcategories.SelectSubcategories([]models.Subcategory{}))
	}
	subcategoryOutput := models.GetSubcategories(category)

	return Render(w, r, subcategories.SelectSubcategories(subcategoryOutput))
}

func SearchSubcategories(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	search := r.FormValue("subcategory_search")
	category := r.FormValue("category")

	searchedSubcategories, err := models.SubcategorySearchByName(search, category)
	if err != nil {
		err_msg := "Internal server error:\nerror searching subcategories: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	return Render(w, r, subcategories.SearchSubcategories(searchedSubcategories))
}

func ShowModifiableSubcategories(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	category_id := r.FormValue("category")

	if category_id == "" {
		if LoggedInUserType == "admin" {
			w.Header().Add("HX-Redirect", "/categories")
			return Render(w, r, categories.Categories(LoggedInUserType))
		} else {
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
		}
	}
	subcategoryOutput := models.GetSubcategories(category_id)

	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func ShowSubcategoryPopup(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

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
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	subcategory_name := r.FormValue("subcategory_name")
	category_id := r.FormValue("category_id")

	subcategory_name_exists, err := models.DoesSubcategoryNameExist(subcategory_name, category_id)
	if err != nil {
		err_msg := "Internal server error:\nerror checking whether subcategory name exists: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if subcategory_name_exists {
		subcategoryOutput := models.GetSubcategories(category_id)
		return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, true))
	}
	err = models.CreateSubcategory(subcategory_name, category_id)
	if err != nil {
		err_msg := "Internal server error:\nerror creating subcategory: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	subcategoryOutput := models.GetSubcategories(category_id)
	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func UpdateSubcategory(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	subcategory_id := chi.URLParam(r, "subcategory_id")
	subcategory_name := r.FormValue("subcategory_name")
	category_id := r.FormValue("category_id")

	subcategory_name_exists, err := models.DoesSubcategoryNameExist(subcategory_name, category_id)
	if err != nil {
		err_msg := "Internal server error:\nerror checking whether subcategory name exists: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	if subcategory_name_exists {
		old_name, err := models.IsOldSubcategoryName(subcategory_id, category_id, subcategory_name)
		if err != nil {
			err_msg := "Internal server error:\nerror checking whether subcategory name is the old name: " + err.Error()
			w.Header().Add("HX-Redirect", "/error")
			return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
		}

		if !old_name {
			subcategoryOutput := models.GetSubcategories(category_id)
			return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, true))
		} else {
			subcategoryOutput := models.GetSubcategories(category_id)
			return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
		}
	}

	err = models.UpdateSubcategory(subcategory_id, subcategory_name, category_id)
	if err != nil {
		err_msg := "Internal server error:\nerror updating subcategory: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	subcategoryOutput := models.GetSubcategories(category_id)
	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func DeleteSubcategory(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	subcategory_id := chi.URLParam(r, "subcategory_id")
	category_id := r.FormValue("category_id")

	err := models.DeleteSubcategory(subcategory_id, category_id)
	if err != nil {
		err_msg := "Internal server error:\nerror deleting subcategory: " + err.Error()
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	}

	subcategoryOutput := models.GetSubcategories(category_id)
	return Render(w, r, subcategories.ModifiableSubcategories(subcategoryOutput, category_id, false))
}

func ShowSubcategoryAlreadyExistsError(w http.ResponseWriter, r *http.Request) error {
	if LoggedInUserType != "admin" {
		w.Header().Add("HX-Redirect", "/error")
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, "Access Denied: Lack of administratorial credentials!"))
	}

	if LoggedInUserType == "" {
		w.Header().Add("HX-Redirect", "/")
		return Render(w, r, login.Login(false, false, "", ""))
	}

	return Render(w, r, subcategories.SubcategoryExistsError())
}
