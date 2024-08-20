package models

import (
	"database/sql"
	"log"
	"strings"

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

	query := `select * from category order by lower(category_name);`
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

func CategorySearchByName(search_term string) []Category {
	var db *sql.DB = database.DB_Connection
	var categories []Category
	category := Category{}

	search_term = strings.ToLower(search_term)
	search_term = "%" + search_term + "%"

	query := `select * from category where lower(category_name) like $1 order by lower(category_name);`
	rows, err := db.Query(query, search_term)
	if err != nil {
		log.Fatal("Error getting categories by name: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&category.Category_ID, &category.Category_Name)
		if err != nil {
			log.Fatal("error scanning categories by name: ", err)
		}
		categories = append(categories, category)
	}

	return categories
}

func CreateCategory(name string) {
	var db *sql.DB = database.DB_Connection
	query := `insert into category values(gen_random_uuid(), $1);`

	_, err := db.Exec(query, name)
	if err != nil {
		log.Fatal("error inserting category: ", err)
	}
}

func UpdateCategory(id string, name string) {
	var db *sql.DB = database.DB_Connection
	query := `update category set category_name = $1 where category_id = $2;`

	_, err := db.Exec(query, name, id)
	if err != nil {
		log.Fatal("error updating category: ", err)
	}
}

func DeleteCategory(id string) {
	var db *sql.DB = database.DB_Connection
	query := `delete from category where category_id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("error deleting category: ", err)
	}
}
