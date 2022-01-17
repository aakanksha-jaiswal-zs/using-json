package main

import (
	"encoding/json"
	"log"
	"time"
)

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

func marshalToStruct() ([]byte, error) {
	const layout = "2006-01-02"
	date := "1999-12-31"

	dob, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}

	s := student{
		FName:   "Jane",
		LName:   "Doe",
		Marks:   91.5,
		DOB:     dob,
		Courses: []string{"mathematics", "physics", "computer science"},
		Others: others{
			SSN:           "BA91653567289Z90",
			MaritalStatus: "unmarried",
		},
	}

	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func marshalToMap() ([]byte, error) {
	student := map[string]interface{}{
		"firstName": "Jane",
		"lastName":  "Doe",
		"marks":     91.7,
		"dob":       "2012-04-23T18:25:43.511Z",
		"courses":   []string{"mathematics", "physics", "computer science"},
		"others": others{
			SSN:           "BA91653567289Z90",
			MaritalStatus: "unmarried",
		},
	}

	body, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	body, err := marshalToStruct()
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))

	body, err = marshalToMap()
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}
