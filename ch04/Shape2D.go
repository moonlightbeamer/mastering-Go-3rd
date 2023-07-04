package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.X+r.Y)
}

func main() {
	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", a.R, a.Perimeter())
    
	var j  Shape2D
	j = a
	// _, ok := interface{}(a).(Shape2D)
	_, ok := j.(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}
    
	b := rectangle{X: 3, Y: 5}
	fmt.Printf("X %.2f, Y %.2f -> Perimeter %.3f \n", b.X, b.Y, b.Perimeter())

	_, ok = interface{}(b).(Shape2D)
	if ok {
		fmt.Println("b is a Shape2D!")
	}

	i := 12
	_, ok = interface{}(i).(Shape2D)
	if ok {
		fmt.Println("i is a Shape2D!")
	}
	fmt.Println(ok)
}
