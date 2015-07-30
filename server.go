package main

import (
	"encoding/json"
	//"fmt"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"strings"
	//	"strconv"
	"io/ioutil"
)

type Userdata struct {
	Name     string `json:"first_name"`
	Lname    string `json:"last_name"`
	Age      string `json:"age"`
	Id       string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type params struct {
	Count  int
	Offset int
	Query  string
}

type count int

type Event struct {
	Name          string `json:"name"`
	EventId       string `json:"event_id"`
	Description   string `json:"description"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end_datetime"`
	Image         string `json:"image_url"`
}

type Response struct {
	Result   bool     `json:"result"`
	Userdata Userdata `json:"userdata"`
	Count    int      `json:"count"`
	Events   []Event  `json:"event_list"`
}

type Venue struct {
	Id   string `json:"venue_id"`
	Name string `json:"name"`
}

type Venues struct {
	Venues []Venue `json:"venues"`
}

func mockApi(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	TotalEvents := []Event{
		Event{
			"Mask Halloween All Inclusive Party 1",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"", //"http://107.161.31.65/xrepo.service/assets/kuumba123.jpg",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"http://107.161.31.65/xrepo.service/assets/kuumba123.jpg",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"http://107.161.31.65/xrepo.service/assets/kuumba123.jpg",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"http://107.161.31.65/xrepo.service/assets/kuumba123.jpg",
		},
		Event{
			"Mask Halloween All Inclusive Party 5",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"http://107.161.31.65/xrepo.service/assets/kuumba123.jpg",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 10",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 15",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 20",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 25",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 30",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 32",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 35",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 40",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 45",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 15:00:00",
			"",
		},
		Event{
			"Mask Halloween All Inclusive Party 50",
			"59",
			"MASK the 14th Annual Halloween Party Miami # 1 Annual Theme Ritual Miami # 1 Open Bar Party - Costume Fete Miami # 1 Bottomless Mug Party",
			"2014-11-01 15:00:00",
			"2014-11-01 18:00:00",
			"",
		},
	}
	var parameters params
	Events := TotalEvents
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &parameters)
	if parameters.Query != "" {
		SearchedEvents := []Event{}
		for _, event := range Events {
			if strings.Contains(event.Name, parameters.Query) {
				SearchedEvents = append(SearchedEvents, event)
			}
		}
		Events = SearchedEvents
		TotalEvents = SearchedEvents
	}
	if parameters.Count != 0 {
		var index int
		if len(Events) > (parameters.Offset + parameters.Count) {
			index = parameters.Offset + parameters.Count
		} else {
			index = len(Events)
		}
		if index != 0 {
			Events = Events[parameters.Offset:index]
		}
	}
	user := Response{
		true,
		Userdata{"Sergio", "Polo", "25", "1", "polosergio@gmail.com", "polosergio"},
		len(TotalEvents),
		Events,
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Write(response)
}

func venueApi(r http.ResponseWriter, w *http.Request) {
	venues := Venues{
		[]Venue{
			Venue{"1", "testLocation"},
			Venue{"2", "testLocationTwo"},
		},
	}
	response, err := json.Marshal(venues)
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
		r.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		r.Header().Set("Access-Control-Allow-Headers", "GET,POST,PUT,DELETE")
	})
	m.Use(martini.Static("apis", martini.StaticOptions{Prefix: "/api/VenueService/GetVenueList"}))
	m.Use(martini.Static("xrepo.manager/", martini.StaticOptions{Prefix: "xrepo.manager", Fallback: "index.html"}))
	m.Use(martini.Static("xrepo.manager/", martini.StaticOptions{Fallback: "index.html"}))
	m.Post("/api/AuthenticationService/Login", mockApi)
	m.Post("/api/AuthenticationService/CreateAccount", mockApi)
	m.Post("/api/AuthenticationService/ResetPassword", mockApi)
	m.Post("/api/EventService/GetEventList", mockApi)
	m.Post("/api/EventService/GetEventDetails", mockApi)
	m.Post("/api/EventService/AddEvent", mockApi)
	m.Post("/api/EventService/DeleteEvent", mockApi)
	m.Post("/api/EventService/UpdateEvent", mockApi)
	log.Fatal(http.ListenAndServe(":8080", m))
}
