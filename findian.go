package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string

	fmt.Printf("Enter a string:")
	fmt.Scan(&a)

	var str string = strings.ToLower(a)

	if strings.HasPrefix(str, "i") && strings.ContainsAny(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}

}
