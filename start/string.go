package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "this is a string"
	fmt.Printf("%v \n", str)
	var isString = true

	isString = strings.Contains(str, "string")

	fmt.Println(isString)
}
