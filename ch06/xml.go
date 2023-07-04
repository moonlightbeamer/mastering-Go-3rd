package main

import (
	"encoding/xml"
	"fmt"
)

type Employee struct {
	XMLName   xml.Name `xml:"employee"`
	ID        int      `xml:"id,attr"`
	ID2       string   `xml:"id2,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
	          Address
	Omit      string   `xml:“-,”` 		  
	Chd	      string   `xml:“,chardata”`
	Cd        string   `xml:“,cdata”`
	Inn       string   `xml:“,innerxml”`
	Comment   string   `xml:",comment"`
}

type Address struct {
	City, Country string
}

func main() {
	r := Employee{FirstName: "Mihalis", LastName: "Tsoukalos", ID: 7, Inn: "Verbatim", Height: 7, Chd: "random string Chd", Cd: "another string"}
	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	output, err := xml.MarshalIndent(&r, "  ", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	final := output
	output = []byte(xml.Header + string(output))
	
	fmt.Printf("%s\n*********\n", output)
	fmt.Printf("%s\n******===******\n", final)
	fmt.Println(xml.Header)
}
