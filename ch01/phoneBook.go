package main

import (
	"fmt"
	"os"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *[]Entry {
	var data_result = []Entry{}
	for i, v := range data {
		if v.Surname == key {
			data_result = append(data_result, data[i])
		}
	}
	return &data_result
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := arguments[0]
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"Mike", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	// Differentiate between the commands
	switch arguments[1] {
	// The search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if len(*result) == 0 {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		for _, v := range *result {
			fmt.Println(v)
		}
	// The list command
	case "list":
		list()
	// Anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
