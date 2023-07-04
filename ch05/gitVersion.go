package main

import (
	"fmt"
	"os"
)

var VERSION string = "1"

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "version" {
			fmt.Println("Version:", VERSION)
		}
	}
}
