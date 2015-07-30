package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"log"
	"encoding/json"
)

type Userdata struct {
	Name	string `json:"first_name"`
	Lname	string `json:"last_name"`
	Age		string `json:"age"`
	Id		string `json:"user_id"`
}

type User struct {
	Result	bool   `json:"result"`
	Userdata Userdata `json:"userdata"`
}

func mockApi(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	user := User {
		true,
		Userdata {"Sergio", "Polo", "25", "1"},
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Write(response)	
}

func main() {
	m := martini.Classic()
	m.Use(func(r http.ResponseWriter) {
		r.Header().Set("Access-Control-Allow-Origin", "*")
	})
	m.Use(martini.Static("xrepo.manager/"))
	m.Post("/xRepo/index.php/AuthenticationService/Login", mockApi)
	m.Post("/xRepo/index.php/EventService/GetEventList", mockApi)
	m.Post("/xRepo/index.php/EventService/GetEventDetails", mockApi)	
	log.Fatal(http.ListenAndServe(":8080", m))
}
