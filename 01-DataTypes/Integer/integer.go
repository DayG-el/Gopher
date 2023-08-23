package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

var (
	num   int   = 100000000000000000 // 0 ->100000000000000000
	num8  int8  = 100                // 0 ->100
	num16 int16 = 10000              // 0 ->10000
	num32 int32 = 10000000           // 0 ->10000000
	num64 int64 = 100000000000000000 // 0 ->100000000000000000
)

func main() {
	fmt.Println("1 byte is 8 bite ")
	fmt.Printf("num %d bytes \n", unsafe.Sizeof(num))
	fmt.Printf("num8 %d bytes \n", unsafe.Sizeof(num8))
	fmt.Printf("num16 %d bytes \n", unsafe.Sizeof(num16))
	fmt.Printf("num32 %d bytes \n", unsafe.Sizeof(num32))
	fmt.Printf("num64 %d bytes \n", unsafe.Sizeof(num64))
	var a = bits.UintSize
	fmt.Println(a)

	var fnum float32 = 10.2 // 3.4 * 10^ -38
	var fnum8 float64       // 1.7 * 10 ^ -308

	fmt.Printf("fnum %d bytes \n", unsafe.Sizeof(fnum))
	fmt.Printf("fnum8 %d bytes \n", unsafe.Sizeof(fnum8))
}
