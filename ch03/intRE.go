package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	sa := arguments[1:]
	t := 0
	f := 0
	for _, s := range sa {
	  ret := matchInt(s)
	  fmt.Println(ret)
	  if ret == true {t +=1} else {f +=1}
	}
	fmt.Printf("true: %d, false: %d", t, f)
}
