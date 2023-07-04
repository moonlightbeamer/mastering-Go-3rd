package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type emptyFile struct {
	Ended bool
	Read  int
}

// Implement error interface
func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%t) but read (%d) bytes", e.Ended, e.Read)
}

// Check values
func isFileEmpty(e error) bool {
	// Type assertion
	v, ok := e.(emptyFile)
	if ok {
		if v.Read == 0 && v.Ended == true {
			return true
		}
	}
	return false
}

func readFile(file string) error {
	//var err error
	fd, open_err := os.Open(file)
	if open_err != nil {
		return open_err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	n := 0
	for {
		line, read_err := reader.ReadString('\n')
		n += len(line)
		if read_err == io.EOF {
			// End of File: nothing more to read
			//if n == 0 {
				return emptyFile{Ended: true, Read: n}
			//}
			//break
		} else if read_err != nil {
			return read_err
		}
	}
	//return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: errorInt <file1> [<file2> ...]")
		return
	}

	for _, file := range flag.Args() {
		err := readFile(file)
		if isFileEmpty(err) {
			fmt.Println(file, "is an empty file, ", err)
		} else {
			v, ok := err.(emptyFile)
			if ok {
                fmt.Println(file, "is a normal file, ", v)
			} else if err != nil {
				fmt.Println(file, "has other error: ", err)
			}
		}	
	}
}
