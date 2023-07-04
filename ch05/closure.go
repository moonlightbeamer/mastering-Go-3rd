package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
    // sum's initial value gets modified on every calling of the lambda func
		sum += x
		return sum
	}
}

func main() {
  // each of pos and neg are getting its own initialized sum from below assignment
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Printf("%[1]s %s %s %[1]s %s\n", "a", "b", "c", "d")
	fmt.Printf("%s %[3]s %s %[2]s %s %[4]s\n", "a", "b", "c", "d")
	fmt.Printf("%s %%m 15%%z \n", "a")
	fmt.Printf("%*s\n", 4, "hi")
	fmt.Printf("%.*f", 5, 5.321)
}
