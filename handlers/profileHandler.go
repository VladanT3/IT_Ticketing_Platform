package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/views/profile"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, profile.Profile(LoggedInUser, LoggedInUserType))
}
