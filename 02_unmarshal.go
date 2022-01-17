package main

import (
	"encoding/json"
	"log"
	"time"
)

/*
	Unmarshalling (decoding) is conversion of raw JSON (strings or bytes) into Go types:
	Structured data to structs, arrays/slices
	Unstructured data to maps and empty interfaces
*/

var jsonData = `{
	"firstName": "Jane",
	"lastName": "Doe",
	"marks": 91.7,
	"dob": "2012-04-23T18:25:43.511Z",
	"courses": ["mathematics", "physics", "computer science"],
	"others": {
		"ssn": "BA91653567289Z90",
		"maritalStatus": "unmarried"
	}
}`

/*
	For unmarshalling to a struct, you must know the structure of JSON.
	The key for a field can be modified by giving JSON tags for every field.
*/

type student struct {
	FName   string    `json:"firstName"`
	LName   string    `json:"lastName"`
	Marks   float64   `json:"marks"`
	DOB     time.Time `json:"dob"`
	Courses []string  `json:"courses"`
	Others  others    `json:"others"`
}

type others struct {
	SSN           string `json:"ssn"`
	MaritalStatus string `json:"maritalStatus"`
}

func unmarshalToStruct() {
	var s student

	err := json.Unmarshal([]byte(jsonData), &s)
	if err != nil {
		log.Printf("error in unmarshalling JSON to struct %v", err)
	}

	log.Printf("Student: %v", s)
}

func unmarshalToAnonymousStruct() {
	s := struct {
		FName   string    `json:"firstName"`
		LName   string    `json:"lastName"`
		Marks   float64   `json:"marks"`
		DOB     time.Time `json:"dob"`
		Courses []string  `json:"courses"`
		Others  struct {
			SSN           string `json:"ssn"`
			MaritalStatus string `json:"maritalStatus"`
		} `json:"others"`
	}{}

	err := json.Unmarshal([]byte(jsonData), &s)
	if err != nil {
		log.Printf("error in unmarshalling JSON to anonymous struct %v", err)
	}

	log.Printf("Student: %v", s)
}

/*
	If you don’t know the structure of your JSON properties beforehand,
	you cannot use structs to unmarshal your data.
*/

func unmarshalToMap() {
	var s map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &s)
	if err != nil {
		log.Printf("error in unmarshalling JSON to map %v", err)
	}

	log.Printf("Student: %v", s)
}

func unmarshalToInterface() {
	var s interface{}

	err := json.Unmarshal([]byte(jsonData), &s)
	if err != nil {
		log.Printf("error in unmarshalling JSON to map %v", err)
	}

	log.Printf("Student: %v", s)
}

func main() {
	unmarshalToStruct()
	unmarshalToAnonymousStruct()

	unmarshalToMap()
	unmarshalToInterface()
}
