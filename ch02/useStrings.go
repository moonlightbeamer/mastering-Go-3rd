package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))

	f("%s\n", s.Title("tHis €wiLL be A title!"))

	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("extra test\n")
    f("12345678901234567890\n")
	f(s.TrimLeft("12227534445678901234567890\n", "12547"))
	f(s.TrimPrefix("12345678901234567890\n", "123"))
	f("end of extra test\n")

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis"))
    f("Compare: %v\n", s.Compare("bA", "Ba"))
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	for _, v := range t {
		fmt.Println(v)
	}
	f("%s\n", t)
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))
    for _, v := range t {
		fmt.Println(v)
	}
	f("%s\n", t)
	f("%s\n", s.Split("abcd e\tfg", ""))
	fmt.Printf("%s", s.Replace("abcd efg.", "", "_", -1))
	fmt.Println()
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123+++432+++", "++"))

	trimFunction := func(c rune) bool {
		return unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("a\t12a3 abc-ABC. \t .A", trimFunction))
}
