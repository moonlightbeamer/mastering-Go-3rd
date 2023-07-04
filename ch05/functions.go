package main

import (
	"fmt"
  "os"
  "strconv"
)

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

// Sorting from smaller to bigger value
func sortTwo(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func main() {
  if len(os.Args) == 1 {
    fmt.Println("usage: functions number")
    return
  }
	n, _ := strconv.Atoi(os.Args[1])
	d, s := doubleSquare(n)
	fmt.Println("Double of", n, "is", d)
	fmt.Println("Square of", n, "is", s)

	// An anonymous function
	anF:= func(param int) int {
		return param * param
	}
	fmt.Println("the anF of", n, "is:", anF(n))
	fmt.Println(sortTwo(1, -3))
	fmt.Println(sortTwo(-1, 0))
	fmt.Println(sortTwo(doubleSquare(n)))
}
