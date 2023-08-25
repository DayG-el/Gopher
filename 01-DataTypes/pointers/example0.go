// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//Pointers in Go are variables that store the memory address of another variable.

func main() {
	a, b := 20, 40
	//Declaration and Initialization
	var ap *int = &a
	var bp *int = &b

	fmt.Println("Address of a:", &a)
	fmt.Println("Address of ap  pointer:", ap)

	fmt.Println("Address of b:", &b)
	fmt.Println("Address of bp pointer:", bp)

	a2 := a
	a2 = a2 + 2
	fmt.Println("value of a after increase a2", a)
	fmt.Println("value of a2 after increase a2", a2)

	fmt.Println("Address of a2:", &a2)

	var ap2 *int = &a

	*ap2 = *ap2 + 2
	fmt.Println("value of ap2", *ap2)
	fmt.Println("value of a", a)
	num := 20
	println("Address of num:", &num)
	changeNumberByValue(num)
	println("value of num:", num)

	changeNumberByRef(&num)
	println("value of num:", num)
}
func changeNumberByRef(num *int) {
	println("changeNumberByRef Address of num:", num)
	*num = *num + 2
}

func changeNumberByValue(num int) {
	println("changeNumberByValue Address of num:", &num)
	num = num + 2
}
