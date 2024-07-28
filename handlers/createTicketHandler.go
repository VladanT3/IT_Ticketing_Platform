package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/VladanT3/IT_Ticketing_Platform/models"
	"github.com/VladanT3/IT_Ticketing_Platform/views/ticket"
	"github.com/google/uuid"
)

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) error {
	session, _ := Store.Get(r, "ticket")
	if ticketType, ok := session.Values["ticketType"].(string); !ok {
		log.Fatal("unable to convert ticket type")
	} else if category, ok := session.Values["category"].(string); !ok {
		log.Fatal("unable to convert category")
	} else if subcategory, ok := session.Values["subcategory"].(string); !ok {
		log.Fatal("unable to convert subcategory")
	} else if title, ok := session.Values["title"].(string); !ok {
		log.Fatal("unable to convert title")
	} else if desc, ok := session.Values["desc"].(string); !ok {
		log.Fatal("unable to convert description")
	} else if customerContact, ok := session.Values["customerContact"].(string); !ok {
		log.Fatal("unable to convert customer contact")
	} else if saveType, ok := session.Values["saveType"].(string); !ok {
		log.Fatal("unable to convert save type")
	} else {
		teamID := models.GetAnalystsTeam(LoggedInUser.Analyst_id.String()).Team_ID

		if category == "none" {
			newTicket := models.Ticket{
				Type: ticketType,
				Category: uuid.NullUUID{
					UUID:  uuid.Nil,
					Valid: false,
				},
				Subcategory: uuid.NullUUID{
					UUID:  uuid.Nil,
					Valid: false,
				},
				Title:            title,
				Description:      desc,
				Customer_Contact: customerContact,
			}

			return Render(w, r, ticket.Ticket(models.Ticket{}, LoggedInUserType, "update", "Please select a valid category!", "Please select a valid subcategory!", newTicket))
		}

		var db *sql.DB = database.DB_Connection
		query := `insert into ticket values(gen_random_uuid(), default, $1, 'Open', $2, $3, $4, $5, $6, default, default, null, $7, $8, $9, null) returning ticket_id;`
		newTicketID := ""
		err := db.QueryRow(query, ticketType, category, subcategory, title, desc, customerContact, teamID, LoggedInUser.Analyst_id, LoggedInUser.Analyst_id).Scan(&newTicketID)
		if err != nil {
			errMsg := "error inserting ticket: " + err.Error()
			http.Error(w, errMsg, http.StatusInternalServerError)
		}

		if saveType == "Save" {
			http.Redirect(w, r, "/ticket/"+newTicketID, http.StatusSeeOther)
			return nil
		} else if saveType == "Save and Exit" {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return nil
		}

		return nil
	}

	return nil
}
