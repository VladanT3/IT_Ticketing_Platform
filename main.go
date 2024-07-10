package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	testUser := User{
		id:       "1",
		name:     "analyst",
		username: "analyst1",
		password: "123",
	}
	router.HandleFunc("GET /", Login(testUser))

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	fmt.Println("Server started on: http://localhost:8000")
	log.Fatal(server.ListenAndServe())
}
