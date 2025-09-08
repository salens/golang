package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/salens/golang-project/model"
)

var peopleMap = make(map[int]model.Person)
var nextID = 1

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	json.NewDecoder(r.Body).Decode(&person)
	person.ID = nextID
	peopleMap[nextID] = person
	nextID++
	writeJSON(w, person)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var people []model.Person
	for _, person := range peopleMap {
		people = append(people, person)
	}
	writeJSON(w, people)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
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
			http.Error(w, "Method not allowed", 405)
		}
	})
}

func StartSimpleAPI() {
	setupAPI()
	fmt.Println("API server starting on :8080")
	fmt.Println("POST   /people - Create person")
	fmt.Println("GET    /people - Get all people")
	fmt.Println("PUT    /people - Update person")
	fmt.Println("DELETE /people - Delete person")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DemoAPI() {
	fmt.Println("=== Simple API Demo ===")
	fmt.Println("Run StartSimpleAPI() in a goroutine to test")
	fmt.Println("Examples:")
	fmt.Println("  Create: curl -X POST -d '{\"name\":\"John\",\"age\":30}' http://localhost:8080/people")
	fmt.Println("  Read:   curl http://localhost:8080/people")
	fmt.Println("  Update: curl -X PUT -d '{\"id\":1,\"name\":\"John Updated\",\"age\":31}' http://localhost:8080/people")
	fmt.Println("  Delete: curl -X DELETE -d '{\"id\":1}' http://localhost:8080/people")
	fmt.Println("========================\n")
}

func TestAPIEndpoints() {
	fmt.Println("=== Testing API Endpoints ===")
	
	baseURL := "http://localhost:8080/people"
	client := &http.Client{Timeout: 5 * time.Second}
	
	// Test 1: Create person
	fmt.Println("1. Creating person...")
	createData := `{"name":"Alice","age":25}`
	resp, err := client.Post(baseURL, "application/json", strings.NewReader(createData))
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("   Status: %s\n", resp.Status)
	}
	
	// Test 2: Create another person
	fmt.Println("2. Creating another person...")
	createData2 := `{"name":"Bob","age":30}`
	resp, err = client.Post(baseURL, "application/json", strings.NewReader(createData2))
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("   Status: %s\n", resp.Status)
	}
	
	// Test 3: Get all people
	fmt.Println("3. Getting all people...")
	resp, err = client.Get(baseURL)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("   Status: %s\n", resp.Status)
		fmt.Printf("   Response: %s\n", string(body))
	}
	
	// Test 4: Update person
	fmt.Println("4. Updating person...")
	updateData := `{"id":1,"name":"Alice Updated","age":26}`
	req, _ := http.NewRequest("PUT", baseURL, strings.NewReader(updateData))
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("   Status: %s\n", resp.Status)
	}
	
	// Test 5: Delete person
	fmt.Println("5. Deleting person...")
	deleteData := `{"id":1}`
	req, _ = http.NewRequest("DELETE", baseURL, strings.NewReader(deleteData))
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("   Status: %s\n", resp.Status)
	}
	
	// Test 6: Get all people after deletion
	fmt.Println("6. Getting all people after deletion...")
	resp, err = client.Get(baseURL)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("   Status: %s\n", resp.Status)
		fmt.Printf("   Response: %s\n", string(body))
	}
	
	fmt.Println("=== API Testing Complete ===\n")
}