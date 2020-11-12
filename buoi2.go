package main

import (
	"fmt"
)

const (
	x = 1 << (iota * 10)
	y
	z
	t
)

func main() {
	// rune = int32 -- ma unicode
	// byte = uint8 -- ma ASCII
	// s := "ok boi"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }

	// fmt.Println("******************")

	// for _, c := range s {
	// 	fmt.Println(c)
	// }

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
