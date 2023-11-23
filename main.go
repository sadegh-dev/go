package main

import (
	"fmt"
)

var (
	name string
	//age      int
	//is_admin bool
)

func main() {
	name = "charlie"

	fmt.Println(name)
	//fmt.Println(*name)
	fmt.Println(&name)

}
