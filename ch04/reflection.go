package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}

	r := reflect.ValueOf(A)
	fmt.Println("String value:", r) // r.String())

	iType := r.Type()
	fmt.Printf("i Type: %s of kind %s\n", iType, iType.Kind())
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value %v\n", r.Field(i)) // r.Field(i).Interface())

		// Check whether there are other structures embedded in Record
		k := r.Field(i).Type().Kind()
		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Printf("type: %s\tkind: %s\n", r.Field(i).Type(), k)
		}

		// Same as before but using the internal value
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
