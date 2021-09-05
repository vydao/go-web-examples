package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Age       int    `json:"age"`
}

func encode(w http.ResponseWriter, r *http.Request) {
	u := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       31,
	}

	json.NewEncoder(w).Encode(u)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var u User

	json.NewDecoder(r.Body).Decode(&u)

	fmt.Fprintf(w, "%s %s is %d years old!", u.Firstname, u.Lastname, u.Age)
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	http.ListenAndServe(":8080", nil)
}
