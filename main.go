package main

import "fmt"

func main() {
	fmt.Printf("Hello Gophers Again!")

	for i := 0; i < 100; i++ {
		println(i)
	}

	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			fmt.Printf("This is an even number \t %v\n", j)
		}
	}

	add := func(a, b int) int {
		return a + b
	}

	result := add(5, 3)
	fmt.Printf("Result: %d\n", result)

	// Anonymous function with parameters (example)
	multiply := func(x int, y int) int {
		return x * y
	}
	mulResult := multiply(4, 7)
	fmt.Printf("Multiplication Result: %d\n", mulResult)

	// Anonymous function with multiple parameters and immediate execution
	func(name string, age int, scores ...float64) {
		sum := 0.0
		for _, score := range scores {
			sum += score
		}
		average := sum / float64(len(scores))
		fmt.Printf("Student %s, age %d, has average score: %.2f\n", name, age, average)
	}("Alice", 20, 85.5, 92.0, 88.5)
}
