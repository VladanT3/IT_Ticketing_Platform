package models

import (
	"database/sql"
	"log"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Manager struct {
	Manager_ID uuid.UUID
}

func IsUserManager(analystID string) bool {
	var db *sql.DB = database.DB_Connection
	manager := Manager{}

	query := `select * from manager where manager_id = $1;`
	err := db.QueryRow(query, analystID).Scan(&manager.Manager_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal("error checking if user is manager: ", err)
	}

	return true
}
