package main

import (
	"fmt"
)

func main() {
	// Create an empty map
	studentGrades := make(map[string]int)

	// Add grades to the map
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	// Access and print individual grades
	fmt.Println("Alice's grade:", studentGrades["Alice"])
	fmt.Println("Bob's grade:", studentGrades["Bob"])
	fmt.Println("Charlie's grade:", studentGrades["Charlie"])

	// Iterate over the map
	for name, grade := range studentGrades {
		fmt.Println(name, "scored", grade)
	}
}