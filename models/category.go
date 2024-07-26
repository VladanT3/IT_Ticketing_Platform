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

func GetAllCategories() []Category {
	var db *sql.DB = database.DB_Connection
	var categories []Category

	query := `select * from category;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("error getting categories: ", err)
	}
	defer rows.Close()

	category := Category{}
	for rows.Next() {
		err = rows.Scan(&category.Category_ID, &category.Category_Name)
		if err != nil {
			log.Fatal("error scanning category: ", err)
		}

		categories = append(categories, category)
	}

	return categories
}
