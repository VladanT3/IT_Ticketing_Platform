package models

import (
	"database/sql"
	"log"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Subcategory struct {
	Subcategory_ID   uuid.UUID
	Category_ID      uuid.UUID
	Subcategory_Name string
}

func GetSubcategoryName(subcategory_id uuid.UUID) string {
	var db *sql.DB = database.DB_Connection
	var subcategory_name string

	query := `select subcategory_name from subcategory where subcategory_id = $1;`
	err := db.QueryRow(query, subcategory_id).Scan(&subcategory_name)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		log.Fatal("error getting subcategory name: ", err)
	}

	return subcategory_name
}
