package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// For loop used as do-while loop
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	i = 100
	fmt.Println(i)

	j := 0
	for {
		if j == 10 {
			break
		}
		fmt.Print(j*j, " ")
		j++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}

	bSlice := []int{-1, 2, 1, -1, 2, -2}
	for _, v := range bSlice {
		fmt.Println("value: ", v)
	}

	cSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, _ := range cSlice {
		fmt.Println("key: ", i)
	}

	dSlice := []int{-1, 2, 1, -1, 2, -2}
	for i:= range dSlice {
		fmt.Println("key: ", i)
	}
}
