// Declare package main - this file contains example functions
package main

// Import necessary packages
import (
	// Import fmt package for string formatting
	"fmt"

	// Import custom model package
	"github.com/salens/golang-project/model"
)

// ReturnPerson returns a Person struct with the given id, name, and age
func ReturnPerson(id int, name string, age int) model.Person {
	// Create and return a Person struct with the provided values
	return model.Person{
		// Set the ID field
		ID:   id,
		// Set the Name field
		Name: name,
		// Set the Age field
		Age:  age,
	}
}

// GetUserInfo takes a name and age, returns a greeting string and doubled age
func GetUserInfo(name string, age int) (string, int) {
	// Create a formatted greeting string using the provided name
	greeting := fmt.Sprintf("Hello, %s!", name)
	// Return the greeting and the age multiplied by 2
	return greeting, age * 2
}

// ReturnTwoInts returns two integers
func ReturnTwoInts(x int, y int) (int, int) {
	// Return both input parameters unchanged
	return x, y
}

// ProcessThreeStrings takes 3 string arguments and returns 3 strings
func ProcessThreeStrings(first string, second string, third string) (string, string, string) {
	// Process the strings and return modified versions
	return fmt.Sprintf("First: %s", first), fmt.Sprintf("Second: %s", second), fmt.Sprintf("Third: %s", third)
}