package models

import (
	"database/sql"
	"log/slog"
	"strings"

	"github.com/VladanT3/IT_Ticketing_Platform/internal/database"
	"github.com/google/uuid"
)

type Analyst struct {
	Analyst_ID               uuid.UUID
	First_Name               string
	Last_Name                string
	Email                    string
	Password                 string
	Phone_Number             string
	Team_ID                  uuid.NullUUID
	Number_of_Open_Tickets   int
	Number_of_Opened_Tickets int
	Number_of_Closed_Tickets int
}

func CheckEmail(email string) (Analyst, bool, error) {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}
	query := `select * from analyst where email = $1;`

	err := db.QueryRow(query, email).Scan(
		&analyst.Analyst_ID,
		&analyst.First_Name,
		&analyst.Last_Name,
		&analyst.Email,
		&analyst.Password,
		&analyst.Phone_Number,
		&analyst.Team_ID,
		&analyst.Number_of_Open_Tickets,
		&analyst.Number_of_Opened_Tickets,
		&analyst.Number_of_Closed_Tickets,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Analyst{}, false, nil
		}
		return Analyst{}, false, err
	}

	return analyst, true, nil
}

func CheckPassword(password string, email string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	var correctPassword bool
	query := `select (password = crypt($1, password)) as password from analyst where email = $2;`
	err := db.QueryRow(query, password, email).Scan(&correctPassword)
	if err != nil {
		return false, err
	}
	if !correctPassword {
		return false, nil
	} else {
		return true, nil
	}
}

func GetAnalyst(analyst_id string) Analyst {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}

	query := `select * from analyst where analyst_id = $1;`
	err := db.QueryRow(query, analyst_id).Scan(
		&analyst.Analyst_ID,
		&analyst.First_Name,
		&analyst.Last_Name,
		&analyst.Email,
		&analyst.Password,
		&analyst.Phone_Number,
		&analyst.Team_ID,
		&analyst.Number_of_Open_Tickets,
		&analyst.Number_of_Opened_Tickets,
		&analyst.Number_of_Closed_Tickets,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Analyst{}
		}
		slog.Error("error getting analyst", "error message", err)
		return Analyst{}
	}

	return analyst
}

func UpdateLoggedInUser(loggedInUser Analyst) (Analyst, error) {
	var db *sql.DB = database.DB_Connection
	analyst := Analyst{}
	query := `select * from analyst where analyst_id = $1;`

	err := db.QueryRow(query, loggedInUser.Analyst_ID).Scan(
		&analyst.Analyst_ID,
		&analyst.First_Name,
		&analyst.Last_Name,
		&analyst.Email,
		&analyst.Password,
		&analyst.Phone_Number,
		&analyst.Team_ID,
		&analyst.Number_of_Open_Tickets,
		&analyst.Number_of_Opened_Tickets,
		&analyst.Number_of_Closed_Tickets,
	)
	if err != nil {
		return Analyst{}, err
	}

	return analyst, nil
}

func GetAllAnalysts() []Analyst {
	var db *sql.DB = database.DB_Connection
	analysts := []Analyst{}

	query := `select * from analyst order by first_name, last_name;`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("error getting all analysts", "error message", err)
		return []Analyst{}
	}
	defer rows.Close()

	analyst := Analyst{}
	for rows.Next() {
		err = rows.Scan(
			&analyst.Analyst_ID,
			&analyst.First_Name,
			&analyst.Last_Name,
			&analyst.Email,
			&analyst.Password,
			&analyst.Phone_Number,
			&analyst.Team_ID,
			&analyst.Number_of_Open_Tickets,
			&analyst.Number_of_Opened_Tickets,
			&analyst.Number_of_Closed_Tickets,
		)
		if err != nil {
			slog.Error("error scanning all analysts", "error message", err)
			return []Analyst{}
		}

		analysts = append(analysts, analyst)
	}

	return analysts
}

func GetTeamsAnalysts(team_id string) []Analyst {
	var db *sql.DB = database.DB_Connection
	analysts := []Analyst{}

	query := `select * from analyst where team_id = $1 order by first_name, last_name;`
	rows, err := db.Query(query, team_id)
	if err != nil {
		slog.Error("error getting a teams analysts", "error message", err)
		return []Analyst{}
	}
	defer rows.Close()

	analyst := Analyst{}
	for rows.Next() {
		err = rows.Scan(
			&analyst.Analyst_ID,
			&analyst.First_Name,
			&analyst.Last_Name,
			&analyst.Email,
			&analyst.Password,
			&analyst.Phone_Number,
			&analyst.Team_ID,
			&analyst.Number_of_Open_Tickets,
			&analyst.Number_of_Opened_Tickets,
			&analyst.Number_of_Closed_Tickets,
		)
		if err != nil {
			slog.Error("error scanning a teams analysts", "error message", err)
			return []Analyst{}
		}

		analysts = append(analysts, analyst)
	}

	return analysts
}

func FilterUsers(search_term string, view_type string, team_id string, user_type string) []Analyst {
	var db *sql.DB = database.DB_Connection
	analysts := []Analyst{}
	analyst := Analyst{}
	search_term = strings.ToLower(search_term)
	search_term = "%" + search_term + "%"
	var queryArgs []any
	var query string

	if user_type == "managers" {
		query = `
			select an.*
			from analyst an join manager m on an.analyst_id = m.manager_id
			where (lower(an.first_name) like $1 or
			lower(an.last_name) like $1)
		`
	} else if user_type == "administrators" {
		query = `
			select an.*
			from analyst an join administrator ad on an.analyst_id = ad.administrator_id
			where (lower(an.first_name) like $1 or
			lower(an.last_name) like $1)
		`
	} else {
		query = `
			select an.* 
			from analyst an
			where (lower(an.first_name) like $1 or
			lower(an.last_name) like $1)
		`
	}

	queryArgs = append(queryArgs, search_term)

	if view_type == "Team View" {
		query += `
			and an.team_id = $2
		`
		queryArgs = append(queryArgs, team_id)
	}

	query += `
		order by an.first_name, an.last_name;
	`
	rows, err := db.Query(query, queryArgs...)
	if err != nil {
		slog.Error("error filtering analysts", "error message", err)
		return []Analyst{}
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&analyst.Analyst_ID,
			&analyst.First_Name,
			&analyst.Last_Name,
			&analyst.Email,
			&analyst.Password,
			&analyst.Phone_Number,
			&analyst.Team_ID,
			&analyst.Number_of_Open_Tickets,
			&analyst.Number_of_Opened_Tickets,
			&analyst.Number_of_Closed_Tickets,
		)
		if err != nil {
			slog.Error("error scanning filtered analysts", "error message", err)
			return []Analyst{}
		}
		analysts = append(analysts, analyst)
	}

	return analysts
}

func UpdateAnalyst(new_analyst Analyst, user_type string) error {
	var db *sql.DB = database.DB_Connection
	query := `
		update analyst set
		first_name = $1,
		last_name = $2,
		email = $3,
		phone_number = $4,
		team_id = $5
		where analyst_id = $6;
	`
	_, err := db.Exec(query, new_analyst.First_Name, new_analyst.Last_Name, new_analyst.Email, new_analyst.Phone_Number, new_analyst.Team_ID.UUID, new_analyst.Analyst_ID)
	if err != nil {
		return err
	}

	manager_check, err := IsUserManager(new_analyst.Analyst_ID.String())
	if err != nil {
		return err
	}
	admin_check, err := IsUserAdmin(new_analyst.Analyst_ID.String())
	if err != nil {
		return err
	}

	if user_type == "manager" {
		if !manager_check {
			query = `insert into manager values($1);`
			_, err := db.Exec(query, new_analyst.Analyst_ID)
			if err != nil {
				return err
			}
		}
		if admin_check {
			query = `delete from administrator where administrator_id = $1;`
			_, err := db.Exec(query, new_analyst.Analyst_ID)
			if err != nil {
				return err
			}
		}
	} else if user_type == "admin" {
		if !admin_check {
			query = `insert into administrator values($1);`
			_, err := db.Exec(query, new_analyst.Analyst_ID)
			if err != nil {
				return err
			}
		}
		if manager_check {
			query = `delete from manager where manager_id = $1;`
			_, err := db.Exec(query, new_analyst.Analyst_ID)
			if err != nil {
				return err
			}
		}
	} else {
		if admin_check {
			query = `delete from administrator where administrator_id = $1;`
			_, err := db.Exec(query, new_analyst.Analyst_ID)
			if err != nil {
				return err
			}
		} else if manager_check {
			query = `delete from manager where manager_id = $1;`
			_, err := db.Exec(query, new_analyst.Analyst_ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DeleteAnalyst(analyst_id string) error {
	var db *sql.DB = database.DB_Connection
	query := `delete from analyst where analyst_id = $1;`

	_, err := db.Exec(query, analyst_id)
	if err != nil {
		return err
	}

	manager_check, err := IsUserManager(analyst_id)
	if err != nil {
		return err
	}
	admin_check, err := IsUserAdmin(analyst_id)
	if err != nil {
		return err
	}

	if admin_check {
		query = `delete from administrator where administrator_id = $1;`
		_, err := db.Exec(query, analyst_id)
		if err != nil {
			return err
		}
	} else if manager_check {
		query = `delete from manager where manager_id = $1;`
		_, err := db.Exec(query, analyst_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateAnalyst(new_analyst Analyst, user_type string) error {
	var db *sql.DB = database.DB_Connection
	query := `insert into analyst values(gen_random_uuid(), $1, $2, $3, $4, $5, $6, default, default, default) returning analyst_id;`
	var new_analyst_id uuid.UUID

	err := db.QueryRow(query, new_analyst.First_Name, new_analyst.Last_Name, new_analyst.Email, new_analyst.Password, new_analyst.Phone_Number, new_analyst.Team_ID.UUID).Scan(&new_analyst_id)
	if err != nil {
		return err
	}

	if user_type == "manager" {
		query = `insert into manager values($1);`
		_, err := db.Exec(query, new_analyst_id)
		if err != nil {
			return err
		}
	} else if user_type == "admin" {
		query = `insert into administrator values($1);`
		_, err := db.Exec(query, new_analyst_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func IsEmailSame(analyst_id string, email string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	var count int

	query := `select count(*) from analyst where analyst_id = $1 and email = $2;`
	err := db.QueryRow(query, analyst_id, email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func UserExists(analyst_id string) (bool, error) {
	var db *sql.DB = database.DB_Connection
	var count int

	query := `select count(*) from analyst where analyst_id = $1;`
	err := db.QueryRow(query, analyst_id).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
