package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

// Format 1
type F1 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// Format 2
type F2 struct {
	Name       string
	Surname    string
	Areacode   string
	Tel        string
	LastAccess string
}

type Book1 []F1
type Book2 []F2

// CSVFILE resides in the home directory of the current user
var CSVFILE = ""
var d1 = Book1{}
var d2 = Book2{}

func readCSVFile(filepath string) error {
	fileInfo, info_err := os.Stat(filepath)
	if info_err != nil {
		return fmt.Errorf("%s not found!", filepath)
	} else if !(fileInfo.Mode().IsRegular()) {
		return fmt.Errorf("%s is not a regular file!", filepath)
	}

	csv_file, open_err := os.Open(filepath)
	if open_err != nil {
		return open_err
	}
	defer csv_file.Close()

	lines, read_err := csv.NewReader(csv_file).ReadAll()
	if read_err != nil {
		return read_err
	}

	var firstLine bool = true
	var csv_format string = "0"
	for _, line := range lines {
		if firstLine {
			if len(line) == 4 {
				csv_format = "1"
			} else if len(line) == 5 {
				csv_format = "2"
			} else {
				csv_format = "0"
				return errors.New("Unsupported File Format!")
			}
			firstLine = false
		}

		switch csv_format {
    case "1":
      temp := F1{
        Name:       line[0],
        Surname:    line[1],
        Tel:        line[2],
        LastAccess: line[3],
      }
      d1 = append(d1, temp)
    case "2":
      temp := F2{
        Name:       line[0],
        Surname:    line[1],
        Areacode:   line[2],
        Tel:        line[3],
        LastAccess: line[4],
      }
      d2 = append(d2, temp)
    case "0":
      return errors.New("unsupported format")
    }
  }
  return nil
}
// Implement sort.Interface for Book1
func (a Book1) Len() int {
	return len(a)
}

// First based on surname. If they have the same
// surname take into account the name.
func (a Book1) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		if a[i].Name == a[j].Name {
      return a[i].Tel < a[j].Tel
    }
    return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a Book1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Implement sort.Interface for Book2
func (a Book2) Len() int {
	return len(a)
}

// First based on areacode. If they have the same
// areacode take into account the surname.
func (a Book2) Less(i, j int) bool {
	if a[i].Areacode == a[j].Areacode {
    if a[i].Surname == a[j].Surname {
      if a[i].Name == a[j].Name {
        return a[i].Tel < a[j].Tel
      }
      return a[i].Name < a[j].Name
    }
    return a[i].Surname < a[j].Surname
	}
	return a[i].Areacode < a[j].Areacode
}

func (a Book2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func sortData(data interface{}) {
	// type switch
	switch T := data.(type) {
	case Book1:
		sort.Sort(T)
		for _, v := range T {
      fmt.Println(v)
    }
	case Book2:
		sort.Sort(T)
		for _, v := range T {
      fmt.Println(v)
    }
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func main() {
	if len(os.Args) != 1 {
		CSVFILE = os.Args[1]
	} else {
		fmt.Println("No data file!")
		return
	}

	err := readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println("main:", err)
		return
	}

	if len(d1) != 0 {
		sortData(d1)
	} else {
		sortData(d2)
	}
}
