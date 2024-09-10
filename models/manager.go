package models

import (
	"database/sql"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Manager struct {
	Manager_ID uuid.UUID
}

func IsUserManager(analyst_id string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	manager := Manager{}

	query := `select * from manager where manager_id = $1;`
	err := db.QueryRow(query, analyst_id).Scan(&manager.Manager_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func IsUserManagerNoError(analyst_id string) bool {
	var db *sql.DB = database.DB_Connection
	manager := Manager{}

	query := `select * from manager where manager_id = $1;`
	err := db.QueryRow(query, analyst_id).Scan(&manager.Manager_ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return true
}
