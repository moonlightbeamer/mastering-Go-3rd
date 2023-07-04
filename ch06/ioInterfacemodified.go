package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1   S1
	text []byte
}

// Using pointer to S1 for changes to be persistent when the method exits
func (s *S1) Read(p *[]byte) (n int, err error) {
	fmt.Print("Give me your name: ")
	fmt.Scanln(p)
	fmt.Println(*p)
	s.F2 = string(*p)
	fmt.Println(s.F2)
	fmt.Println(len(*p))
	return len(*p), nil
}

func (s *S1) Write(p []byte) (n int, err error) {
	if s.F1 < 0 {
		return -1, nil
	}

	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s ", p)
	}
	fmt.Println()
	return s.F1, nil
}

func (s S2) eof() bool {
	if len(s.text) == 0 {fmt.Println("eof() running")}
	return len(s.text) == 0
}

func (s *S2) readByte() byte {
	// this function assumes that eof() check was done before
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return
	}
  fmt.Println("s:", s)
  fmt.Println("p:", p)
	l := len(p)
	if l > 0 {
		for n < l {
			p[n] = s.readByte()
      fmt.Println("l:", l, "n:", n, "p[n]:", p[n], "p[n] char:", string(p[n]))
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
    fmt.Println("p:", p)
	}
	return
}

func main() {
	s1var := S1{4, "Hello"}
	fmt.Println(s1var)

	buf := make([]byte, 2)
	_, err := s1var.Read(&buf)
	fmt.Println("len of buf:", len(buf), "capacity of buf:", cap(buf), "value of buf:", buf, "string of buf:", string(buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read:", s1var.F2)
	fmt.Println(s1var)
	_, _ = s1var.Write([]byte("Hello There!"))
	s1var.Write([]byte(s1var.F2))

	s2var := S2{F1: s1var, text: []byte("Hello world!!")}
	// Read s2var.text
	r := bufio.NewReaderSize(&s2var, 4)
  fmt.Println("r:", *r)
	for {
    fmt.Println("begin of for loop:", "len of buf:", len(buf), "capacity of buf:", cap(buf), "value of buf:", buf, "string of buf:", string(buf))
		fmt.Println("before read:", s2var)
    n, err := r.Read(buf)
    fmt.Println("after read:", s2var)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", n, "*", err, "*", string(buf[:n]), "of:", string(buf))
	}
  fmt.Println("end of for loop:", "len of buf:", len(buf), "capacity of buf:", cap(buf), "value of buf:", buf, "string of buf:", string(buf))
}
