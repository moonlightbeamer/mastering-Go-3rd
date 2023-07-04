package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	buffer := "Data to write\n"

	f1, err := os.Create("./f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, buffer)

	f2, err := os.Create("./f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(buffer)
	fmt.Printf("wrote %d bytes to f2\n", n)

	f3, err := os.Create("./f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(buffer)
	fmt.Printf("wrote %d bytes to f3\n", n)
	w.Flush()

	f := "./f4.txt"
	f4, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	for i := 0; i < 5; i++ {
		n, err = io.WriteString(f4, buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("wrote %d bytes to f4\n", n)
	}

	// Append to a file
	f4, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	// Write() needs a byte slice
	n, err = f4.WriteString("Put some more data at the end.\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("appended %d bytes to f4\n", n)
}
