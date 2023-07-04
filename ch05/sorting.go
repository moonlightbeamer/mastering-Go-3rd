package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grades{{"J.", "Lewis", 10}, {"M.", "Tsoukalos", 7},
		{"D.", "Tsoukalos", 8}, {"J.", "Lewis", 9}}
    
	for i :=0; i < 2; i++ {
		isSorted := sort.SliceIsSorted(data, func(i, j int) bool { return data[i].Grade < data[j].Grade })

		if isSorted {
			fmt.Println("It is sorted! As is:")
		} else {
			fmt.Println("It is NOT sorted! Sorting...")
			sort.Slice(data, func(i, j int) bool { return data[i].Grade < data[j].Grade })
		}

		fmt.Println("By Grade:", data)
	}	
}
