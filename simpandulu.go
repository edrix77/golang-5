package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	InputReader := bufio.NewReader(os.Stdin)
	input, _ := InputReader.ReadString('\n')

	var str string = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	fmt.Printf("%d", len(str))

	// if strings.HasPrefix(str, "i") && strings.ContainsAny(str, "a") && strings.HasSuffix(str, "n") {
	if "n" == str[len(str)-1:] {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}

}
