package main

import "fmt"

func GetUserInfo(name string, age int) (string, int) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, age * 2
}