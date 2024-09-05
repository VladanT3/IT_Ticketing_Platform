package handlers

import (
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/analyst"
	"github.com/VladanT3/IT_Ticketing_Platform/views/team"
)

func Profile(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, analyst.Profile(LoggedInUser, LoggedInUserType))
}

func GetTeamsAnalysts(w http.ResponseWriter, r *http.Request) error {
	team := r.FormValue("team")

	if team == "none" {
		return Render(w, r, analyst.SelectAnalyst(models.GetAllAnalysts()))
	} else {
		analysts := models.GetTeamsAnalysts(team)
		return Render(w, r, analyst.SelectAnalyst(analysts))
	}
}

func GetAnalystsTeam(w http.ResponseWriter, r *http.Request) error {
	analyst_id := r.FormValue("analyst")

	analyst := models.GetAnalyst(analyst_id)

	if analyst_id == "none" {
		return Render(w, r, team.SelectTeam(analyst))
	} else {
		return Render(w, r, team.SelectTeam(analyst))
	}
}
