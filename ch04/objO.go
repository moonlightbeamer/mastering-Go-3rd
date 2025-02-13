package main

import (
	"fmt"
)

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

// Structure c has two fields
type c struct {
	A a
	B b
}

type d struct {
	a
	b
}

// Satisfying IntA
func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

// Satisfying IntB
func (varC c) bar() {
	fmt.Println("Bar Processing", varC)
}

// Structure compose gets the fields of structure a
type compose struct {
	field1 int
	a
}

// Different structures can have methods with the same name
func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var iC c = c{A: a{120, 12}, B: b{"-12", -12}}
	iD := d{a{111,222}, b{"13", 13}}

	iC.A.A()
	iC.B.A()
    fmt.Println(iD.YY, iD.AA) //iD.XX is ambigious, can't be referenced
	// The following two statements does not work
	// ./objO.go:71:33: mixture of field:value and value initializers
	// iComp := compose{field1: 123, a{456, 789}}
	// ./objO.go:75:32: cannot use promoted field a.XX in struct literal of type compose
	//./objO.go:75:41: cannot use promoted field a.YY in struct literal of type compose
	// iComp := compose{field1: 123, XX: 456, YY: 789}

	iComp := compose{123, a{456, 789}}
	// iComp := compose{field1: 123, a: a{456, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)

	iC.bar()
	iC.foo()
	processA(iC)
}
