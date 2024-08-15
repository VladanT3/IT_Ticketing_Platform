package models

import (
	"database/sql"
	"log"
	"strings"

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

func GetAllSubcategories() []Subcategory {
	var db *sql.DB = database.DB_Connection
	var subcategories []Subcategory

	query := `select subcategory_id, subcategory_name from subcategory;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("error getting subcategories: ", err)
	}
	defer rows.Close()

	subcategory := Subcategory{}
	for rows.Next() {
		err = rows.Scan(&subcategory.Subcategory_ID, &subcategory.Subcategory_Name)
		if err != nil {
			log.Fatal("error scanning subcategory: ", err)
		}

		subcategories = append(subcategories, subcategory)
	}

	return subcategories
}

func GetSubcategories(category_id string) []Subcategory {
	var db *sql.DB = database.DB_Connection
	var subcategories []Subcategory

	query := `select subcategory_id, subcategory_name from subcategory where category_id = $1;`
	rows, err := db.Query(query, category_id)
	if err != nil {
		log.Fatal("error getting categories subcategories: ", err)
	}
	defer rows.Close()

	subcategory := Subcategory{}
	for rows.Next() {
		err = rows.Scan(&subcategory.Subcategory_ID, &subcategory.Subcategory_Name)
		if err != nil {
			log.Fatal("error scanning categories subcategories: ", err)
		}

		subcategories = append(subcategories, subcategory)
	}

	return subcategories
}

func SubcategorySearchByName(search_term string) []Subcategory {
	var db *sql.DB = database.DB_Connection
	var subcategories []Subcategory
	subcategory := Subcategory{}

	search_term = strings.ToLower(search_term)
	search_term = "%" + search_term + "%"

	query := `select subcategory_id, subcategory_name from subcategory where lower(subcategory_name) like $1;`
	rows, err := db.Query(query, search_term)
	if err != nil {
		log.Fatal("Error getting subcategories by name: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&subcategory.Subcategory_ID, &subcategory.Subcategory_Name)
		if err != nil {
			log.Fatal("error scanning subcategories by name: ", err)
		}
		subcategories = append(subcategories, subcategory)
	}

	return subcategories

}
