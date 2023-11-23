package main

import "testing"

func TestMain(t *testing.T) {
	// Create an empty map
	studentGrades := make(map[string]int)

	// Add grades to the map
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	// Check individual grades
	if studentGrades["Alice"] != 90 {
		t.Error("Expected Alice's grade to be 90, but got", studentGrades["Alice"])
	}
	if studentGrades["Bob"] != 85 {
		t.Error("Expected Bob's grade to be 85, but got", studentGrades["Bob"])
	}
	if studentGrades["Charlie"] != 95 {
		t.Error("Expected Charlie's grade to be 95, but got", studentGrades["Charlie"])
	}
}
