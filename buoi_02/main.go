package main

import (
	"fmt"
)

// ------------------- khai bao bien
func main() {
	var x int
	x = 1
	fmt.Println(x)

	var y int = 1
	fmt.Println(y)

	// inferred
	var z = 4.5
	fmt.Println(z)

	// var t int
	// t = 8
	t := 8
	fmt.Println(t)

	var r = float64(x)
	fmt.Println(r)

	c := complex(4, 3)
	fmt.Println(c)
	fmt.Println(real(c))
	fmt.Println(imag(c))

	d := 4 + 3i
	fmt.Println(d)
	fmt.Println(real(d))
	fmt.Println(imag(d))

	// rune = int32 -- ma unicode
	// byte = uint8 -- ma ASCII
	s := "ok boi"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for _, c := range s {
		fmt.Println(c)
	}
}

// ------------------- iota, point

const (
	x = 1 << (iota * 10)
	y
	z
	t
)

func main() {
	// fmt.Println("******************")

	//t := 3 // 11
	//fmt.Println(t << 1) // 110

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(t)

	t := 3
	addOne(&t)

	fmt.Println(t)
}

func addOne(x *int) {
	*x++
}
