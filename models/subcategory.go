package models

import (
	"database/sql"
	"log/slog"
	"strings"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Subcategory struct {
	Subcategory_ID   uuid.UUID
	Category_ID      uuid.UUID
	Subcategory_Name string
}

func GetSubcategories(category_id string) []Subcategory {
	var db *sql.DB = database.DB_Connection
	var subcategories []Subcategory

	query := `select * from subcategory where category_id = $1 order by lower(subcategory_name);`
	rows, err := db.Query(query, category_id)
	if err != nil {
		slog.Error("error getting a categories subcategories", "error message", err)
		return []Subcategory{}
	}
	defer rows.Close()

	subcategory := Subcategory{}
	for rows.Next() {
		err = rows.Scan(&subcategory.Subcategory_ID, &subcategory.Category_ID, &subcategory.Subcategory_Name)
		if err != nil {
			slog.Error("error scanning a categories subcategories", "error message", err)
			return []Subcategory{}
		}

		subcategories = append(subcategories, subcategory)
	}

	return subcategories
}

func SubcategorySearchByName(search_term string, categoryID string) ([]Subcategory, error) {
	var db *sql.DB = database.DB_Connection
	var subcategories []Subcategory
	subcategory := Subcategory{}

	search_term = strings.ToLower(search_term)
	search_term = "%" + search_term + "%"

	query := `select * from subcategory where category_id = $1 and lower(subcategory_name) like $2;`
	rows, err := db.Query(query, categoryID, search_term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&subcategory.Subcategory_ID, &subcategory.Category_ID, &subcategory.Subcategory_Name)
		if err != nil {
			return nil, err
		}
		subcategories = append(subcategories, subcategory)
	}

	return subcategories, nil

}

func CreateSubcategory(name string, category_id string) error {
	var db *sql.DB = database.DB_Connection
	query := `insert into subcategory values(gen_random_uuid(), $1, $2);`

	_, err := db.Exec(query, category_id, name)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSubcategory(subcategory_id string, name string, category_id string) error {
	var db *sql.DB = database.DB_Connection
	query := `update subcategory set subcategory_name = $1 where subcategory_id = $2 and category_id = $3;`

	_, err := db.Exec(query, name, subcategory_id, category_id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubcategory(subcategory_id string, category_id string) error {
	var db *sql.DB = database.DB_Connection
	query := `delete from subcategory where subcategory_id = $1 and category_id = $2;`

	_, err := db.Exec(query, subcategory_id, category_id)
	if err != nil {
		return err
	}

	return nil
}

func DoesSubcategoryNameExist(name string, category_id string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	var count int
	name = strings.ToLower(name)
	query := `select count(*) from subcategory where category_id = $1 and lower(subcategory_name) = $2;`

	err := db.QueryRow(query, category_id, name).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

func GetSubcategoryIDByName(name string) (uuid.UUID, error) {
	var db *sql.DB = database.DB_Connection
	query := `select subcategory_id from subcategory where subcategory_name = $1;`
	var subcategory_id uuid.UUID

	err := db.QueryRow(query, name).Scan(&subcategory_id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return subcategory_id, nil
}
