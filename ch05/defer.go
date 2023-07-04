package main

import (
	"fmt"
)

func d1() {
	fmt.Print("d1():")
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	fmt.Print("d2():")
	for i := 3; i > -1; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	fmt.Print("d3():")
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	fmt.Print("\nmain1 done\n")
	d2()
	fmt.Print("\nmain2 done\n")
	d3()
	fmt.Print("\nmain3 done\n")
}
