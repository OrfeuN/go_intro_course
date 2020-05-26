package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func handleHome(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "<a href=\"/people/\">People</a>")
}

//comment
type AllPeople struct {
	People []Person `json:"results"`
}

//comment
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

//comment
type Person struct {
	Name         string `json:"name"`
	Homeworldurl string `json:"homeworld"`
	Homeworld    Planet
}

func handlePeople(writer http.ResponseWriter, req *http.Request) {
	var people AllPeople

	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	// fmt.Fprintln(writer, res)

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	err = json.Unmarshal(data, &people)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	for _, person := range people.People {
		var planet Planet = Planet{}

		err, planet = person.getHomeWorld()

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		person.Homeworld = planet
	}

	fmt.Fprintln(writer, people)
}

func (person Person) getHomeWorld() (error, Planet) {

	var planet Planet = Planet{}

	res, err := http.Get(person.Homeworldurl)

	if err != nil {
		return err, planet
	}

	var bytes []byte

	bytes, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return err, planet
	} else {
		err = json.Unmarshal(bytes, &planet)
	}

	return err, planet
}

func main() {
	http.HandleFunc("/people/", handlePeople)
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
