package main

import (
	"encoding/json"
	"log"
)

// "omitempty" tag allows us to omit empty field values.

func marshalWithOmitempty() {
	s := struct {
		FName   string   `json:"firstName,omitempty"`
		LName   string   `json:"lastName,omitempty"`
		Age     int      `json:"age,omitempty"`
		Marks   float64  `json:"marks,omitempty"`
		Courses []string `json:"courses,omitempty"`
		Others  struct {
			SSN           string `json:"ssn,omitempty"`
			MaritalStatus string `json:"maritalStatus,omitempty"`
		} `json:"others,omitempty"`
	}{
		FName:   "",  // "" string is omitted
		LName:   "",  // "" string is omitted
		Age:     0,   // 0 integer is omitted
		Marks:   0,   // 0 float is omitted
		Courses: nil, // nil slices and [] (empty) slices are omitted
		Others: struct {
			SSN           string `json:"ssn,omitempty"`           // "" string is omitted
			MaritalStatus string `json:"maritalStatus,omitempty"` // "" string is omitted
		}{}, // empty structs are not omitted
	}

	body, err := json.Marshal(&s)
	if err != nil {
		log.Printf("error in unmarshalling JSON to struct %v", err)
	}

	log.Printf("Student: %v", string(body))
}

func marshalWithPointer() {
	s := struct {
		FName   string   `json:"firstName,omitempty"`
		LName   string   `json:"lastName,omitempty"`
		Age     int      `json:"age,omitempty"`
		Marks   float64  `json:"marks,omitempty"`
		Courses []string `json:"courses,omitempty"`
		Others  *struct {
			SSN           string `json:"ssn,omitempty"`
			MaritalStatus string `json:"maritalStatus,omitempty"`
		} `json:"others,omitempty"`
	}{
		FName:   "",  // "" string is omitted
		LName:   "",  // "" string is omitted
		Age:     0,   // 0 integer is omitted
		Marks:   0,   // 0 float is omitted
		Courses: nil, // nil slices and [] (empty) slices are omitted
		Others:  nil, // nil pointers are omitted
	}

	body, err := json.Marshal(&s)
	if err != nil {
		log.Printf("error in unmarshalling JSON to struct %v", err)
	}

	log.Printf("Student: %v", string(body))
}

func main() {
	marshalWithOmitempty()
	marshalWithPointer()
}
