package main

import (
	"fmt"
	"os"
)

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum = sum + a
	}
	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input...)
}

func main() {
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)
	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum:", sum)
	fmt.Println("s:", s)
	fmt.Print("everything(s,s,s): ")
	everything(s, s, s)
	temp := []interface{}{}
	for _, v := range s {
		temp = append(temp, v)
	}
	fmt.Println("temp:", temp)
	fmt.Print("everything(temp): ")
	everything(temp)
	fmt.Print("everything(temp...): ")
	everything(temp...)
	fmt.Println("os.Args[1:]:", os.Args[1:])
	fmt.Print("everything(os.Args[1:]): ")
    everything(os.Args[1:])
	// Cannot directly pass []string as []interface{}
	// You have to convert it first!
	emp := []interface{}{}
	for _, v := range os.Args[1:] {
		emp = append(emp, v)
	}
	fmt.Println("emp:", emp)
	fmt.Print("everything(emp): ")
	everything(emp)
	fmt.Print("everything(emp...): ")
	everything(emp...)

	// There is a slightly different way to do the conversion
	arguments := os.Args[1:]
	empty := make([]interface{}, len(arguments))
	for i := range arguments {
		empty[i] = arguments[i]
	}
	fmt.Println("empty:", empty)
	fmt.Print("everything(empty...): ")
	everything(empty...)

	// This will work!
	str := []string{"One", "Two", "Three"}
	everything(str, str, str)
	tmp := []interface{}{}
	for _, v := range str {
		tmp = append(tmp, v)
	}
	fmt.Println("tmp:", tmp)
	fmt.Print("everything(tmp...): ")
	everything(tmp...)
    fmt.Println(1,2,3,4)
	fmt.Print("everything(1,2,3,4): ")
	everything(1,2,3,4)
	fmt.Print("everything([]interface{}{1,2,3,4,5}...): ")
	everything([]interface{}{1,2,3,4,5}...)
}
