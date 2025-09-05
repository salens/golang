// Declare package main - this is the entry point of the program
package main

// Import standard and third-party packages
import (
	// Import fmt package for formatted I/O
	"fmt"
	// Import runtime package for runtime information
	"runtime"
	// Import time package for time-related functions
	"time"

	// Import custom model package
	"github.com/salens/golang-project/model"
)

// init function runs before main
func init() {
	// Print the number of CPUs available on the system
	fmt.Println("Number of CPUs:", runtime.NumCPU())
}

// Main function - program entry point
func main() {
	// Record the start time to measure execution time
	start := time.Now()
	// Instantiate two Person objects
	// Create first person with ID 1, name Alice, age 30
	person1 := model.Person{ID: 1, Name: "Alice", Age: 30}
	// Create second person with ID 2, name Bob, age 25
	person2 := model.Person{ID: 2, Name: "Bob", Age: 25}
	// Print person1 information in a readable format
	fmt.Printf("Person 1: ID=%d, Name=%s, Age=%d years old\n", person1.ID, person1.Name, person1.Age)
	// Print person2 information in a readable format
	fmt.Printf("Person 2: ID=%d, Name=%s, Age=%d years old\n", person2.ID, person2.Name, person2.Age)

	// Call ReturnPerson from functionExamples.go
	// Create a person using the ReturnPerson function with ID 3, name Charlie, age 40
	p := ReturnPerson(3, "Charlie", 40)
	// Print the returned person's details
	fmt.Printf("ReturnPerson: ID=%d, Name=%s, Age=%d\n", p.ID, p.Name, p.Age)
	// Print a greeting message
	fmt.Printf("Hello Gophers Again!")

	// Loop from 0 to 99 and print each number
	for i := 0; i < 100; i++ {
		// Print the current value of i
		println(i)
	}

	// Loop from 0 to 99 to check for even numbers
	for j := 0; j < 100; j++ {
		// Check if the number is even using modulo operator
		if j%2 == 0 {
			// Print even numbers with tab formatting
			fmt.Printf("This is an even number \t %v\n", j)
		}
	}

	// Define an anonymous function that adds two integers
	add := func(a, b int) int {
		// Return the sum of a and b
		return a + b
	}

	// Call the add function with arguments 5 and 3
	result := add(5, 3)
	// Print the result of the addition
	fmt.Printf("Result: %d\n", result)

	// Anonymous function with parameters (example)
	// Define an anonymous function that multiplies two integers
	multiply := func(x int, y int) int {
		// Return the product of x and y
		return x * y
	}
	// Call the multiply function with arguments 4 and 7
	mulResult := multiply(4, 7)
	// Print the multiplication result
	fmt.Printf("Multiplication Result: %d\n", mulResult)

	// Anonymous function with multiple parameters and immediate execution and GO for concurrent execution
	// Launch a goroutine with an anonymous function that calculates average scores
	go func(name string, age int, scores ...float64) {
		// Initialize sum to 0.0 for calculating total score
		sum := 0.0
		// Iterate through all scores using range
		for _, score := range scores {
			// Add each score to the sum
			sum += score
		}
		// Calculate average by dividing sum by number of scores
		average := sum / float64(len(scores))
		// Print student information and average score
		fmt.Printf("Student %s, age %d, has average score: %.2f\n", name, age, average)
	}("Alice", 20, 85.5, 92.0, 88.5)

	// Print OS and architecture
	// Print the operating system
	fmt.Println("runtime.GOOS\t", runtime.GOOS)
	// Print the system architecture
	fmt.Println("runtime.GOARCH\t", runtime.GOARCH)
	// Print the number of CPUs
	fmt.Println("runtime.NumCPU\t", runtime.NumCPU())
	// Print the number of active goroutines
	fmt.Println("runtime.NumGoroutine\t", runtime.NumGoroutine())
	// Print the Go version
	fmt.Println("runtime.Version\t", runtime.Version())
	// Commented out code - call GetUserInfo function with name "John" and age 25
	greeting, doubledAge := GetUserInfo("John", 25)
	// Commented out code - print greeting and doubled age
	fmt.Printf("%s You are %d years old (doubled)\n", greeting, doubledAge)

	// Call ReturnTwoInts function with arguments 10 and 20
	first, second := ReturnTwoInts(10, 20)
	// Print the two integers returned by the function
	fmt.Printf("ReturnTwoInts returned: %d, %d\n", first, second)

	// Record the end time
	stop := time.Now()
	// Calculate elapsed time in seconds
	elapsed := stop.Sub(start).Seconds()
	// Print the total execution time
	fmt.Printf("Elapsed time: %.4f seconds\n", elapsed)
}
