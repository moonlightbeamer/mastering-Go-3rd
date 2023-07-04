package main

import (
	"fmt"
)

func main(){
	value, err := fmt.Printf("Go has strict rules for curly braces!")
	fmt.Println()
	fmt.Println("value is: ", value, "error is:", err)
}

