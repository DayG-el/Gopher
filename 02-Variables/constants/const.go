package main

import "fmt"

func main() {
	fmt.Println("starting new session for constans")
	const pi float64 = 3.14
	const (
		name   = "tunisia"
		city   = "Tunis"
		number = 10
	)
	fmt.Printf("pi is %f , name is %s , city is %s and number is %d\n", pi, name, city, number)
	const googleBaseUrl = "https://www.google.com"
	const googleMapUrl = "/maps"

	println(googleBaseUrl, googleMapUrl)
}
