package main

import (
	"fmt"
)

func main() {
	c := complex(3, 4)
	fmt.Println(c)
	fmt.Println(real(c))
	fmt.Println(imag(c))

	hello()
}
