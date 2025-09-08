// Declare package main - this is the entry point of the program
package main

// Import standard and third-party packages
import (
	// Import fmt package for formatted I/O
	"fmt"
	"sync"

	// Import runtime package for runtime information
	"runtime"
	// Import time package for time-related functions
	"time"

	// Import custom model package
	"github.com/salens/golang-project/model"
)

var wg sync.WaitGroup

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
		fmt.Printf("Student %s, age %d, has average score: %.2f\n\n", name, age, average)
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

	// Call ProcessThreeStrings function with three string arguments
	str1, str2, str3 := ProcessThreeStrings("hello", "world", "example")
	// Print the three strings returned by the function
	fmt.Printf("ProcessThreeStrings returned: %s, %s, %s\n", str1, str2, str3)

	wg.Add(2)
	go func() {
		fmt.Print("Hello from goroutine ome\n")
		wg.Done()
	}()

	go func() {
		fmt.Print("Hello from goroutine two\n")
		wg.Done()
	}()

	// Create two new channels
	// Create integer channel for sending numbers between goroutines
	ch1 := make(chan int)
	// Create string channel for sending text messages between goroutines
	ch2 := make(chan string)

	// Anonymous function that sends integer to ch1
	go func() {
		// Send the number 100 to ch1 channel
		ch1 <- 100
	}()

	// Anonymous function that sends string to ch2
	go func() {
		// Send a message to ch2 channel
		ch2 <- "Message from new channel"
	}()

	// Receive from new channels
	// Receive integer value from ch1 channel
	number := <-ch1
	// Receive string message from ch2 channel
	message := <-ch2
	// Print the received number
	fmt.Printf("Received number: %d\n", number)
	// Print the received message
	fmt.Printf("Received message: %s\n", message)

	// Create a buffered channel with capacity of 3
	// Buffered channels allow goroutines to send values without blocking until buffer is full
	ch3 := make(chan string, 3)

	// Anonymous function that sends multiple messages to buffered channel
	go func() {
		// Send first message to buffered channel - won't block
		ch3 <- "First buffered message"
		// Send second message to buffered channel - won't block
		ch3 <- "Second buffered message"
		// Send third message to buffered channel - won't block
		ch3 <- "Third buffered message"
		// Close the channel to signal no more values will be sent
		close(ch3)
	}()

	// Receive all messages from buffered channel using range
	// Range will automatically stop when channel is closed
	for msg := range ch3 {
		// Print each message received from the buffered channel
		fmt.Printf("Buffered channel message: %s\n", msg)
	}

	// Create send-only channel ch4 example
	ch4 := make(chan int, 2)

	// Anonymous function with send-only channel parameter
	go func(sendOnly chan<- int) {
		// Send values through the send-only channel
		sendOnly <- 42
		sendOnly <- 100
	}(ch4)

	// Receive values from ch4 in main
	val1 := <-ch4
	val2 := <-ch4
	fmt.Printf("Received from send-only channel ch4: %d, %d\n", val1, val2)

	// Make send-only channel example
	//sendCh := make(chan<- string)

	// Make receive-only channel example
	//receiveCh := make(<-chan int)

	// Demonstrate send-only channel
	chSend := make(chan string)
	sendCh := chan<- string(chSend) // convert to send-only
	go func(sendOnly chan<- string) {
		sendOnly <- "Hello from send-only channel!"
		// close is not strictly needed for send-only, but good practice
		close(chSend)
	}(sendCh)

	// Receive from the original channel (as receive-only)
	for msg := range chSend {
		fmt.Printf("Received from send-only channel: %s\n", msg)
	}

	// Demonstrate receive-only channel
	chRecv := make(chan int)
	go func() {
		chRecv <- 123
		close(chRecv)
	}()
	// Pass the channel as receive-only to the goroutine
	go func(receiveOnly <-chan int) {
		for v := range receiveOnly {
			fmt.Printf("Received from receive-only channel: %d\n", v)
		}
	}(chRecv)

	wg.Wait()
	fmt.Println("All goroutines complete.")

	// Demonstrate API examples
	DemoAPI()

	// Start API server in goroutine for testing
	go StartSimpleAPI()

	// Test the API endpoints
	TestAPIEndpoints()

	// Give server time to start
	time.Sleep(1 * time.Second)

	// Record the end time
	stop := time.Now()
	// Calculate elapsed time in seconds
	elapsed := stop.Sub(start).Seconds()
	// Print the total execution time
	fmt.Printf("Elapsed time: %.4f seconds\n", elapsed)
}
