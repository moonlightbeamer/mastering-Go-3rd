package main

import (
  "fmt"
  "reflect"
)

type Location struct {
  Street string
  City   string
  State  string
  zip    int
}

type Person struct {
  Name    string
  Age     int
  Address Location
  Tel     int
}

func main() {
  p := Person{"Alice", 30, Location{"123 Main", "Loveland", "CO", 80301}, 8005551234}

  // Get the reflect.Type, reflect.Kind and reflect.Value of the variable
  t := reflect.TypeOf(p)   // reflect.ValueOf(p).Type() and reflect.TypeOf(p) are the same
  k := t.Kind()            // v.Kind(), v.Type().Kind() and t.Kind() are the same
  v := reflect.ValueOf(p)

  // Print the type and kind of the variable
  fmt.Println("Root Type:\t", t)
  fmt.Println("Root Value:\t", v)
  fmt.Println("Root Kind:\t", k)

  // Iterate over the fields of the struct
  for i := 0; i < t.NumField(); i++ {    // t.NumField() and v.NumField() are the same
    fieldType := t.Field(i)
    fieldKind := t.Field(i).Type.Kind()  // t.Field(i).Type.Kind() and v.Field(i).Kind() are the same
    fieldValue := v.Field(i)

    // Print the field sequence, name, type, kind and value
    fmt.Printf("Root Field %d: %s\t Type:(%s) of Kind:[%s] = %v\n", i, fieldType.Name, fieldType.Type, fieldKind, fieldValue)
    if fieldKind.String() == "struct" {     // fieldKind == reflect.struct and fieldKind.String() == "struct" are the same
      for j := 0; j < fieldType.Type.NumField(); j++ {  // fieldType.Type.NumField() and fieldValue.NumField() are the same
        nestfieldType := fieldType.Type.Field(j)  
        nestfieldKind := fieldType.Type.Field(j).Type.Kind()  // fieldValue.Field(j).Kind() and fieldType.Type.Field(j).Type.Kind() are the same
        nestfieldValue := fieldValue.Field(j)

        // Print the nest filed sequence, name, type, kind and value
        fmt.Printf("\tNest Field %d: %s\t Type:(%s) of Kind:[%s] = %v\n", j, nestfieldType.Name, nestfieldType.Type, nestfieldKind, nestfieldValue)
      }  
    }
  }
}
