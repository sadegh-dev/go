package main

import (
	"fmt"
	"runtime"
	"time"
)

func compare(x, y int) string {

	var result string
	if x > y {
		result = "x greater than y"
	} else if x < y {
		result = "x smaller than y"
	} else {
		result = "x and y in same"
	}
	return result

}

func checkOS() string {

	return runtime.GOOS

}

func timeChecker() int {
	return time.Now().Hour()
}

func main() {
	var (
		name string
	)
	name = "charlie"
	fmt.Println(name)

	t := compare(50, 50)
	fmt.Println(t)

	fmt.Println(checkOS())
	fmt.Println(timeChecker())

}
