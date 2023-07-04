package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
  "strings"
	"time"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}
var MIN = 0
var MAX = 26
var data_result = []Entry{}
func search(key string, d *[]Entry) {
	for _, v := range *d {
		if strings.Count(v.Tel, key) > 0 {
			data_result = append(data_result, v)
		}
	}
	//return &data_result
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func populate(n int, d *[]Entry) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100000000, 9999999999))
		*d = append(*d, Entry{name, surname, n})
    (*d)[0] = Entry{"Hello", "World", "123456789"}
	}
  fmt.Printf("inside of populate func, D has %d entries.\n", len(*d))
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: search|list <arguments>")
		return
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	// How many records to insert
	n := 100
  fmt.Printf("Before populate func, Data has %d entries.\n", len(data))
  fmt.Println("It seems populate func need to pass on the pointers of slice to make it work")
	populate(n, &data)
	fmt.Printf("Data has %d entries.\n", len(data))

	// Differentiate between the commands
	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Tel number")
			return
		}
		search(arguments[2], &data)
		if len(data_result) == 0 {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
    fmt.Println("Although search func can take slice just fine since no change to it, only use, but to maintain consistency, passing pointers as well")
    fmt.Printf("Below are matches found in Data for phone number includes: %s\n", arguments[2])
		for _, v := range data_result {
      fmt.Println(v)
    }
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
