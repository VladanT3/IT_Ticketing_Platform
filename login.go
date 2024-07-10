package main

import "net/http"

func Login(user User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, mockUser := range mockdb {
			if user.username == mockUser.username && user.password == mockUser.password {
				w.Write([]byte(user.name + " logged in"))
				return
			}
		}
		w.Write([]byte("no user found"))
	}
}
