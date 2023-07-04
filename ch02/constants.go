package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		zero int = 3 << iota
		_
		two
		three
		four
		_
		six
	)

	fmt.Println("2^0:", zero)
	fmt.Println("2^2:", two)
	fmt.Println("2^2:", three)
	fmt.Println("2^4:", four)
	fmt.Println("2^6:", six)
}
