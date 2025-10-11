package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// Conditional Operatiors : > < >= <= == !=
	// Logical Operatiors : && || !
	iAge := 8
	if (iAge >= 1) && (iAge <= 17) {
		pl("Important Birthday")
	}
}
