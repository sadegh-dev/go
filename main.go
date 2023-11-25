package main

import (
	"fmt"
)

var (
	co int
)

func main() {
	co = 20
	switch co {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default is done.")

	}

}
