package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var peopleMap = make(map[int]Person)
var nextID = 1

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	person.ID = nextID
	peopleMap[nextID] = person
	nextID++
	writeJSON(w, person)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person
	for _, person := range peopleMap {
		people = append(people, person)
	}
	writeJSON(w, people)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	
	if _, exists := peopleMap[person.ID]; !exists {
		http.Error(w, "Person not found", 404)
		return
	}
	
	peopleMap[person.ID] = person
	writeJSON(w, person)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	var req struct{ ID int `json:"id"` }
	json.NewDecoder(r.Body).Decode(&req)
	
	if _, exists := peopleMap[req.ID]; !exists {
		http.Error(w, "Person not found", 404)
		return
	}
	
	delete(peopleMap, req.ID)
	w.WriteHeader(204)
}

func setupAPI() {
	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		switch r.Method {
		case "POST":
			createPerson(w, r)
		case "GET":
			getPeople(w, r)
		case "PUT":
			updatePerson(w, r)
		case "DELETE":
			deletePerson(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../index.html")
	})
}

func main() {
	setupAPI()
	fmt.Println("API server starting on :8080")
	fmt.Println("POST   /people - Create person")
	fmt.Println("GET    /people - Get all people") 
	fmt.Println("PUT    /people - Update person")
	fmt.Println("DELETE /people - Delete person")
	fmt.Println("Web interface: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}