package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		// reader := bufio.NewReader(os.Stdin)
		// line, _ := reader.ReadString('\n')
		// line = strings.TrimRight(line, "\r\n")
		// line = strings.ToUpper(line)
		//
		// fmt.Println(line)
		//
		// _ = fmt.Sprint

		// exercise 05 - Formatted Print
		// reader := bufio.NewReader(os.Stdin)
		// nameInput, _ := reader.ReadString('\n')
		// ageInput, _ := reader.ReadString('\n')
		// nameInput = strings.TrimRight(nameInput, "\r\n")
		// ageInput = strings.TrimRight(ageInput, "\r\n")
		// age, _ := strconv.Atoi(ageInput)
		//
		// fmt.Printf("Hi, %s! You are %d years old.\n", nameInput, age)
		//
		// _ = age
		// _ = fmt.Print

		// exercise 06 - if, Else Switch 
		// var n int
		// fmt.Scan(&n)
		//
		// if n % 3 == 0 && n % 5 == 0 {
		// 	fmt.Println("FizzBuzz")
		// } else if n % 3 == 0 {
		// 	fmt.Println("Fizz")
		// } else if n % 5 == 0 {
		// 	fmt.Println("Buzz")
		// } else {
		// 	fmt.Println(n)
		// }

		// exercise 07 - Loops with for
		// var n int
		// fmt.Scan(&n)
		//
		// sum := 0
		// for i := 1; i <= n; i++ {
		// 	sum += i
		// }
		//
		// fmt.Println(sum)

		// exercise 08 - Function
		// var input int
		// fmt.Scan(&input)
		// fmt.Println(square(input))

		// exercise 09 - Error handling
		// reader := bufio.NewReader(os.Stdin)
		// stringInput, _ := reader.ReadString('\n')
		// stringInput = strings.TrimRight(stringInput, "\r\n")
		//
		// toInteger, err := strconv.Atoi(stringInput)
		//
		// if err != nil {
		// 	fmt.Println("bad")
		// 	return
		// }
		//
		// fmt.Printf("ok %d", toInteger)

		// practice slice at exercise 10
		// println("Mutate Arr/Src from slice")
		// base1 := []int{10, 10, 10, 10}
		// mutateBase := base1[:]
		// fmt.Println(mutateBase)
		// mutateBase[0] = 20
		// fmt.Println(mutateBase)
		//
		// println("Copy Arr to Slice")
		// base2 := []int{10, 10, 10, 30}
		// copyView := make([]int, 4, 8)
		// copy(copyView, base2[:])
		// fmt.Println(base2)
		// copyView[3] = 40
		// fmt.Println(copyView)

		// exercise 10 - slice
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		inputParts := strings.Fields(strings.TrimSpace(input))

		nums := make([]int, 0, len(inputParts))
		for _, v := range inputParts {
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}

		var maxValue = nums[0]
		for _, v := range nums {
			if  maxValue <= v {
				maxValue = v
			}
		}

		fmt.Println(maxValue)

}

// func rectArea(w, h int) int {
// 	return w * h
// }

// func square(number int) int {
// 	return number * number
// }
