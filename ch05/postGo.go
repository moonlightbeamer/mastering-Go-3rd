package main

import (
	"fmt"
	"math/rand"
	"time"
    "strings"
	"github.com/moonlightbeamer/post05"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return strings.Title(strings.ToLower(temp))
}

func main() {
	//post05.Hostname = "localhost"
	//post05.Port = 5432
	//post05.Username = "mtsouk"
	//post05.Password = "pass"
	//post05.Database = "go"

	data, err := post05.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}
    fmt.Println("LIST DONE")
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(8)

	t := post05.Userdata{
		Username:    random_username,
		Name:        getString(4),
		Surname:     getString(5),
		Description: "This is me!",
	}

	id := post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}
    username, err := post05.SearchUser(id)
	if err != nil {
		fmt.Println(err)
	}
	
	err = post05.DeleteUser(username)
	if err != nil {
		fmt.Println(err)
	}

	// Trying to delete it again!
	err = post05.DeleteUser(username)
	if err != nil {
		fmt.Println(err)
	}

	id = post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = post05.Userdata{
		Username:    random_username,
		Name:        getString(4),
		Surname:     getString(5),
		Description: "This might not be me!"}

	err = post05.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}


}
