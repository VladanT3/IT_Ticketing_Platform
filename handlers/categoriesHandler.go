package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/views/categories"
)

func CategoriesHandler(w http.ResponseWriter, r *http.Request) error {
	//TODO: check if user is admin
	return Render(w, r, categories.Categories(LoggedInUserType))
}
