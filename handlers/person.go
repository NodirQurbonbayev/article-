package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"httpserver/models"
)

func HandlePerson(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createPerson(w, r)
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPut:
		updatePerson(w, r)
	case http.MethodDelete:
		deletePerson(w, r)

	}
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// max ni aniqlab +1 id ga beradi
	newPerson.ID = len(models.People) + 1
	models.People = append(models.People, newPerson)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	newPerson.ID = len(models.People) + 1
	models.People = append(models.People, newPerson)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var newperson models.Person
	idString := r.URL.Query()["id"][0]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	for i, person := range models.People {
		if person.ID == id {
			models.People[i] = newperson
		}
	}
	var i int
	if i == len(models.People) {
		http.Error(w, "Person not found", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.People)

}

func deletePerson(w http.ResponseWriter, r *http.Request) {

}
