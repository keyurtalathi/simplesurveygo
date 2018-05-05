package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

type MovieDetails struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Cast     string `json:"cast"`
	Genre    string `json:"genre"`
	Notes    string `json:"notes"`
}

func main() {
	movies := make(chan MovieDetails)
	var m []MovieDetails

	apiUrl := "https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil) // URL-encoded payload

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	log.Println(resp.StatusCode)

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(m))

	go insert_in_mongo(movies)
	go insert_in_mongo(movies)
	go insert_in_mongo(movies)
	go insert_in_mongo(movies)
	for _, item := range m {
		movies <- item
	}
}
func insert_in_mongo(movies chan MovieDetails) {
	session, er := mgo.Dial("127.0.0.1")
	if er != nil {
		panic(er)
	}
	defer session.Close()
	for {
		select {
		case movie := <-movies:

			collection := session.DB("movies").C("movie_details")
			err := collection.Insert(&movie)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
