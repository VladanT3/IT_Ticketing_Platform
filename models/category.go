package models

import (
	"database/sql"
	"log"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Category struct {
	Category_ID   uuid.UUID
	Category_Name string
}

func GetCategoryName(category_id uuid.UUID) string {
	var db *sql.DB = database.DB_Connection
	var category_name string

	query := `select category_name from category where category_id = $1;`
	err := db.QueryRow(query, category_id).Scan(&category_name)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		log.Fatal("error getting category name: ", err)
	}

	return category_name
}
