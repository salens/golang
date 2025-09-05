// Package model contains data structures used throughout the application
package model

// Person represents a person with basic information
type Person struct {
	// ID is the unique identifier for the person
	ID   int
	// Name is the person's full name
	Name string
	// Age is the person's age in years
	Age  int
}