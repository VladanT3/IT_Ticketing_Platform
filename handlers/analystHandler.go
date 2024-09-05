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

func ShowUserView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, analyst.UserView(LoggedInUserType, LoggedInUser, "User View"))
}

func ShowTeamView(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, analyst.UserView(LoggedInUserType, LoggedInUser, "Team View"))
}

func FilterUsers(w http.ResponseWriter, r *http.Request) error {
	search_term := r.FormValue("search")
	view_type := r.FormValue("view_type")

	analysts := models.FilterUsers(search_term, view_type, LoggedInUser.Team_ID.UUID.String())

	return Render(w, r, analyst.Users(analysts))
}
