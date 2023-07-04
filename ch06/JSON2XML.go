package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type XMLrec struct {
	Name    string `xml:"username"`
	Surname string `xml:"surname,omitempty"`
	Year    int    `xml:"creationyear,omitempty"`
}

type JSONrec struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need XML or JSON input")
		return
	}
	// This can be a JSON or an XML record
	input := []byte(arguments[1])
	fmt.Println(string(input))

	// Check if it is XML
	checkJSON := false
	tempXML := XMLrec{}
	tempJSON := JSONrec{}
	xml_err := xml.Unmarshal(input, &tempXML)
	if xml_err != nil {
		checkJSON = true
	} else {
		tempJSON = JSONrec{Name: tempXML.Name}
		if tempXML.Surname != "" {
			tempJSON.Surname = tempXML.Surname
		}
		if tempXML.Year != 0 {
			tempJSON.Year = tempXML.Year
		}
		s, j_err := json.Marshal(&tempJSON)
		if j_err != nil {
			fmt.Println(j_err)
			return
		}
		fmt.Println(string(s))
		return
	}

	if !checkJSON {
		return
	}

	// ELSE Check if it is JSON
	tempXML = XMLrec{}
	tempJSON = JSONrec{}
	json_err := json.Unmarshal(input, &tempJSON)
	if json_err != nil {
		fmt.Println("Not valid input")
		return
	} else {
		tempXML = XMLrec{Name: tempJSON.Name}
		if tempJSON.Surname != "" {
			tempXML.Surname = tempJSON.Surname
		}
		if tempJSON.Year != 0 {
			tempXML.Year = tempJSON.Year
		}
		s, x_err := xml.Marshal(&tempXML)
		if x_err != nil {
			fmt.Println(x_err)
			return
		}
		fmt.Println(string(s))
	}
}
