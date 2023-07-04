package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

type Shape interface {}
var shape Shape
func Teststruct(x interface{}) {
	// type switch
	switch T := x.(type) {
	case Secret:
		fmt.Printf("supported type: %T, value is: %v\n", T, T)
	case Entry:
		fmt.Printf("supported type: %T, value is: %v\n", T, T)
	default:
		fmt.Printf("Not supported type: %T, value is: %v: \n", T, T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
  case int:
    fmt.Printf("Int Data type: %T, value is: %v\n", T, T)
	default:
		fmt.Printf("Data type: %T, value is: %v\n", T, T)
	}
}

func main() {
  Learn(123)
  Learn(12.23)
	Learn('€')
	//Learn(A)
	Learn(200i+10)
	Learn(shape)

	A := Entry{100, "F2", Secret{"myPassword"}}
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("A string")
}
