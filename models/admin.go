package models

import (
	"database/sql"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Admin struct {
	Admin_ID uuid.UUID
}

func IsUserAdmin(analyst_id string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	admin := Admin{}

	query := `select * from administrator where administrator_id = $1;`
	err := db.QueryRow(query, analyst_id).Scan(&admin.Admin_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
