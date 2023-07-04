package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	var min, max float64
  var first int
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
      first = i
		  //continue
		} else {
      min = n
			max = n
      first = i
			break
		}
	}
  if first == len(arguments)-1 {
    fmt.Println("no valid argument found.")
    return
  } else {
    fmt.Println("First numerical argument:", first)
    for i := first + 1; i < len(arguments); i++ {
      n, err := strconv.ParseFloat(arguments[i], 64)
      if err != nil {
        continue
      } else if n < min {
          min = n
          continue
      } else if n > max {
          max = n
          continue
      }
      fmt.Printf("Min: %.0f, Max: %.0f", min, max)
      fmt.Println()
    }
    fmt.Println("Min:", min)
    fmt.Println("Max:", max)
  }
}
