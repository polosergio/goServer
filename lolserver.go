package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type RitoResponse struct {
	Response []byte 
}

var devKey string

func getSummonerNames(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	params := w.URL.Query().Get("names")
	url := "https://na.api.pvp.net/api/lol/na/v1.4/summoner/by-name/" + params + devKey
	response, err := http.Get(url)
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	r.Write(body)
}

func getLatestMatches(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	w.ParseForm()
	Ids := w.Form["Ids[]"]
	Response := make([]string, 0)
	for _, summonerId := range Ids {
		log.Println(summonerId)
		url := "https://na.api.pvp.net/api/lol/na/v2.2/matchhistory/" + summonerId + devKey
		response, err := http.Get(url)
		if err != nil {
			http.Error(r, err.Error(), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		Response = append(Response, string(body))
	}
	encoded, err := json.Marshal(Response)
	if err != nil {
		log.Println(err)
	}
	r.Write(encoded)
}

func main() {
	devKey = "?api_key=349258f8-e96e-45f9-a89a-249dee4c607b"
	m := martini.Classic()
	m.Use(func(r http.ResponseWriter) {
		r.Header().Set("Access-Control-Allow-Origin", "*")
	})
	m.Use(martini.Static("whosthefeeder/"))
	m.Get("/summonerNames", getSummonerNames)
	m.Post("/latestMatches", getLatestMatches)
	log.Fatal(http.ListenAndServe(":8080", m))
}

