package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %v, error: %s", f, err)
			break
		}
		fmt.Println(line)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
