package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

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

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}

func JSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	// Create random records
	var i int
	var t Data
	for i = 0; i < 2; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}
  /*
	fmt.Println("Last record:", t)
	json, err := json.MarshalIndent(t, "", "\t")
	if err == nil {
		fmt.Println(string(json))
	} else {
    return
  }
  */
	_ = PrettyPrint(t)

  buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)  // one func can't have json.MarshalIndent and json.NewEncoder at same time.
	encoder.SetIndent("", "\t")

	enc_err := encoder.Encode(DataRecords)
	if enc_err == nil {
		fmt.Print(buf)
	}
	
	val, _ := JSONstream(DataRecords)
	fmt.Println(val)
}
