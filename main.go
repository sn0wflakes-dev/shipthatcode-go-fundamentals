package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    // Print the greeting below.
		
		// exercise 01 - Hello, World!
		// fmt.Println("Hello, Go!")

		// exercise 02 - Variable and Types
		// var num1, num2 int
		// fmt.Scan(&num1, &num2)
		//
		// result := num1 + num2
		// fmt.Println(result)

		// exercise 03 -  Numbers and Math
		// var width, height int
		// fmt.Scan(&width, &height)
		//
		// fmt.Println(rectArea(width, height))

		// exercise 04 - Strings
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		line = strings.ToUpper(line)

		fmt.Println(line)

		_ = fmt.Sprint
}

func rectArea(w, h int) int {
	return w * h
}
