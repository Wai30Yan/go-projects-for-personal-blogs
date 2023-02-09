package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var People []Person

var PORT = ":8080"

func main() {
	People = []Person{
		{1, "Wai Yan", 27},
		{2, "John", 48},
		{3, "Jane", 35},
	}

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:         PORT,
        Handler:      mux,
        IdleTimeout:  10 * time.Second,
        ReadTimeout:  time.Second,
        WriteTimeout: time.Second,
	}

	mux.Handle("/get", http.HandlerFunc(getPeople))
	mux.Handle("/person/", http.HandlerFunc(getPerson))
	mux.Handle("/add", http.HandlerFunc(addPerson))
	mux.Handle("/update/", http.HandlerFunc(updatePerson))
	mux.Handle("/delete/", http.HandlerFunc(deletePerson))

	fmt.Println("Servering running at", PORT)
	err := s.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPeople")
	json.NewEncoder(w).Encode(People)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPerson")
	vars := strings.Split(r.URL.Path, "/")
	param := vars[len(vars)-1]
	id, _ := strconv.Atoi(param)
	for _, person := range People {
		if person.ID == id {
			json.NewEncoder(w).Encode(person)
		}
	}
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addPerson")
	person := Person{
		ID: len(People) + 1,
	}

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	People = append(People, person)
	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updatePerson")
	vars := strings.Split(r.URL.Path, "/")
	param := vars[len(vars)-1]
	id, _ := strconv.Atoi(param)

	updatedPerson := Person{
		ID: id,
	}
	json.NewDecoder(r.Body).Decode(&updatedPerson)
	for i, person := range People {
		if person.ID == id {
			People[i] = updatedPerson
			json.NewEncoder(w).Encode(People[i])
			return
		}
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deletePerson")
	vars := strings.Split(r.URL.Path, "/")
	param := vars[len(vars)-1]
	id, _ := strconv.Atoi(param)

	for key, person := range People {
		if person.ID == id {
			People = append(People[:key], People[key+1:]...)

		}
	}
}
