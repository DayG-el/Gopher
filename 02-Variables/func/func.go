package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "this is golang tutorial go go"
	fmt.Println(strings.Contains(myString, "go"))
	fmt.Println(strings.ContainsAny(myString, "go"))
	fmt.Println(strings.Count(myString, "go"))
	fmt.Println(strings.Index(myString, "go"))
	fmt.Println(strings.LastIndex(myString, "go"))
	fmt.Println(strings.Split(myString, " "))
	fmt.Println(strings.Join(strings.Split(myString, " "), "-"))
	fmt.Println(strings.Trim(myString, ";"))
	fmt.Println(strings.TrimLeft(myString, ";"))
	fmt.Println(strings.TrimRight(myString, ";"))
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))

}
