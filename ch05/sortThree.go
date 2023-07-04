package main

import (
	"fmt"
  "sort"
  "os"
  _ "strconv"
)

type SortThree []string
var parameters SortThree

func (s SortThree) Len() int {
  return len(s)
}

func (s SortThree) Less (i, j int) bool {
  ii:= s[i]
  jj:= s[j]
  return ii < jj
}

func (s SortThree) Swap (i, j int) {
  s[i], s[j] = s[j], s[i]
}

func main() {
  if len(os.Args) != 4 {
    fmt.Println("Usaege: sortThree param1 param2 param3")
    return
  }
  for i:=1; i<4; i++ {
    parameters = append(parameters, os.Args[i])
  }
  fmt.Println(parameters)
  sort.Sort(parameters)
  fmt.Println(parameters)

  
}