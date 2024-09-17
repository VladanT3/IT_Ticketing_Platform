package models

import (
	"database/sql"
	"log/slog"
	"strings"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Category struct {
	Category_ID   uuid.UUID
	Category_Name string
}

func GetAllCategories() []Category {
	var db *sql.DB = database.DB_Connection
	var categories []Category

	query := `select * from category order by lower(category_name);`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("error getting all categories", "error message", err)
		return []Category{}
	}
	defer rows.Close()

	category := Category{}
	for rows.Next() {
		err = rows.Scan(&category.Category_ID, &category.Category_Name)
		if err != nil {
			slog.Error("error scanning all categories", "error message", err)
			return []Category{}
		}

		categories = append(categories, category)
	}

	return categories
}

func CategorySearchByName(search_term string) ([]Category, error) {
	var db *sql.DB = database.DB_Connection
	var categories []Category
	category := Category{}

	search_term = strings.ToLower(search_term)
	search_term = "%" + search_term + "%"

	query := `select * from category where lower(category_name) like $1 order by lower(category_name);`
	rows, err := db.Query(query, search_term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&category.Category_ID, &category.Category_Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func CreateCategory(name string) error {
	var db *sql.DB = database.DB_Connection
	query := `insert into category values(gen_random_uuid(), $1);`

	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCategory(id string, name string) error {
	var db *sql.DB = database.DB_Connection
	query := `update category set category_name = $1 where category_id = $2;`

	_, err := db.Exec(query, name, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	var db *sql.DB = database.DB_Connection
	query := `delete from category where category_id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func DoesCategoryNameExist(name string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	var count int
	name = strings.ToLower(name)
	query := `select count(*) from category where lower(category_name) = $1;`

	err := db.QueryRow(query, name).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

func GetCategoryIDByName(name string) (uuid.UUID, error) {
	var db *sql.DB = database.DB_Connection
	query := `select category_id from category where category_name = $1;`
	var category_id uuid.UUID

	err := db.QueryRow(query, name).Scan(&category_id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return category_id, nil
}
