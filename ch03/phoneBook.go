package main

import (
	"encoding/csv"
	"fmt"
  "math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// CSVFILE resides in the home directory of the current user
var CSVFILE = "./csv.data"
var INDFILE = "./index.data"

var data = []Entry{}
var index = map[string]int{}
var data_match = []int{}

func readCSVFile(filepath string, d *[]Entry) error {
	// If the filepath does not exist, create an new one and seed it with 100 rows of random data
	fileInfo, info_err := os.Stat(filepath)
	// If error is not nil, it means that the file does not exist
	if info_err != nil {
		fmt.Println("Didn't find pre-existing \"", filepath, "\" data file, creating a new one.")
		csv_file, create_err := os.Create(filepath)
		if create_err != nil {
			return create_err
		}
    csv_writer := csv.NewWriter(csv_file)
    csv_writer.Comma = ','
    l := 100
    for i := 0; i < l; i++ {
      name := getString(4)
      surname := getString(5)
      tel := strconv.Itoa(random(1999999999, 9999999999))
      time := strconv.FormatInt(time.Now().Unix(), 10)
      temp := []string{name, surname, tel, time}
      csv_writer.Write(temp)
    }
    csv_writer.Flush()
		csv_file.Close()
	} else if !(fileInfo.Mode().IsRegular()) {
		return fmt.Errorf("%s is not a regular file!", filepath)
	}
  
  // loading csv file into data structure after guarantee file availability
	data_file, open_err := os.Open(filepath)
	if open_err != nil {
		return open_err
	}
	defer data_file.Close()

	// CSV file read all at once
	lines, read_err := csv.NewReader(data_file).ReadAll()
	if read_err != nil {
		return read_err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		// Storing to global variable
    if matchNameSur(temp.Name) && matchNameSur(temp.Surname) && matchTel(temp.Tel) && matchTel(temp.LastAccess) {
		  *d = append(*d, temp)
    } else {
      return fmt.Errorf("name and tel fields in data file %s have incorrect format in line %v.", filepath, line)
    }
	}

	return nil
}

func saveCSVFile(filepath string, d *[]Entry) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range *d {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func readIndex(filepath string, d *[]Entry, ind *map[string]int) error {
  // If the filepath does not exist, create an new one and seed it with freshly created index data
	fileInfo, info_err := os.Stat(filepath)
	// If error is not nil, it means that the file does not exist
	if info_err != nil {
		fmt.Println("Didn't find pre-existing \"", filepath, "\" index file, creating a new one.")
		csv_file, create_err := os.Create(filepath)
		if create_err != nil {
			return create_err
		}
    csv_writer := csv.NewWriter(csv_file)
    csv_writer.Comma = ','
    for i, v := range *d {
      temp := []string{v.Tel, strconv.Itoa(i)}
      csv_writer.Write(temp)
    }
    csv_writer.Flush()
		csv_file.Close()
	} else if !(fileInfo.Mode().IsRegular()) {
		return fmt.Errorf("%s is not a regular file!", filepath)
	}

  // loading csv file into index after guarantee file availability
	index_file, open_err := os.Open(filepath)
	if open_err != nil {
		return open_err
	}
	defer index_file.Close()

	// CSV file read all at once
	lines, read_err := csv.NewReader(index_file).ReadAll()
	if read_err != nil {
		return read_err
	}
  
  for _, line := range lines {
		key := line[0]
    i, conv_err := strconv.Atoi(line[1])
    if conv_err != nil { 
      return fmt.Errorf("index in index file %s is not an integer in line %v", filepath, line)
    }
		(*ind)[key] = i
	}
	return nil
}

func saveIndexFile(filepath string, ind *map[string]int) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for k, v := range *ind {
		temp := []string{k, strconv.Itoa(v)}
		csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func initS(N, S, T string) *Entry {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

func insert(pS *Entry, d *[]Entry, ind *map[string]int) error {
	// If it already exists, do not add it
	_, ok := (*ind)[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	*d = append(*d, *pS)
	// Update the index
  (*ind)[(*pS).Tel] = len(*d)-1
	index_err := saveIndexFile(INDFILE, ind)
  if index_err != nil {
		return index_err
	}

	data_err := saveCSVFile(CSVFILE, d)
	if data_err != nil {
		return data_err
	}
	return nil
}

func deleteEntry(key string, d *[]Entry, ind *map[string]int) error {
	i, ok := (*ind)[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	*d = append((*d)[:i], (*d)[i+1:]...)
	// rebuild the index - key does not exist any more and index in data changed too.
  // remove the deleted key before rebuild map to avoid potentially duplicated index, or clean it entirely: (*ind) = map[string]int{} 
  delete(*ind, key)   // or (*ind) = map[string]int{} 
  for re_i, re_v := range *d {
    (*ind)[re_v.Tel] = re_i
  }

	data_err := saveCSVFile(CSVFILE, d)
	if data_err != nil {
		return data_err
	}

  index_err := saveIndexFile(INDFILE, ind)
	if index_err != nil {
		return index_err
	}

	return nil
}

func search(key string, d *[]Entry, ind *map[string]int, re *[]int) error {
	*re = []int{}
  for i, v := range *ind {
		if strings.Count(i, key) > 0 {
			*re = append(*re, v)
		}
	}
  if len(*re) == 0 {
    return fmt.Errorf("%s not found.", key)
  } else {
    for _, re_v := range *re {
      (*d)[re_v].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
    }
    saveCSVFile(CSVFILE, d)
    return nil
  }
}

func list(d *[]Entry) {
	for _, v := range *d {
		fmt.Println(v)
	}
}

func matchNameSur(s string) bool {
	return regexp.MustCompile("^[A-Z][a-z]*$").Match([]byte(s))
}

func matchTel(s string) bool {
	return regexp.MustCompile(`\d+$`).Match([]byte(s))
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(0, 26)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return strings.Title(strings.ToLower(temp))
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	data_err := readCSVFile(CSVFILE, &data)
	if data_err != nil {
		fmt.Println(data_err)
		return
	}

	index_err := readIndex(INDFILE, &data, &index)
	if index_err != nil {
		fmt.Println(index_err)
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := initS(arguments[2], arguments[3], t)
		// If it was nil, there was an error
		if temp != nil {
			insert_err := insert(temp, &data, &index)
			if insert_err != nil {
				fmt.Println(insert_err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		delete_err := deleteEntry(t, &data, &index)
		if delete_err != nil {
			fmt.Println(delete_err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		search_err := search(t, &data, &index, &data_match)
		if search_err != nil {
			fmt.Println(search_err)
			return
		}
		for _, re_v := range data_match {
      fmt.Println(data[re_v])
    }
	case "list":
		list(&data)
	default:
		fmt.Println("Not a valid option")
	}
}
