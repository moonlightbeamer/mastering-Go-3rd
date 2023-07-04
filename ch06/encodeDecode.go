package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:""`
	Surname string `json:"-,"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{Name: "Mike", Surname: "Tsoukalos", Year: 2021}

	// Regular Structure
	// Encoding JSON data -> Convert Go Structure to JSON record with fields
	t, err := json.Marshal(&useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Type: %[1]T \t Value: %[1]s\n", t)
	}

	// Decoding JSON data given as a string
	str := `{"Name": "M.", "-": "Ts", "creat":2020}`
	// Convert string into a byte slice
	jsonRecord := []byte(str)
	// Create a structure variable to store the result
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %[1]T with value %[1]v\n", temp)
	}
}
