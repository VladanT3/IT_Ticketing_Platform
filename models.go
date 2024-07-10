package main

type User struct {
	id       string
	name     string
	username string
	password string
}

var mockdb = map[string]User{
	"analyst": {
		id:       "1",
		name:     "analyst",
		username: "analyst1",
		password: "123",
	},
	"manager": {
		id:       "2",
		name:     "manager",
		username: "manager1",
		password: "456",
	},
	"admin": {
		id:       "3",
		name:     "admin",
		username: "admin1",
		password: "321",
	},
}
