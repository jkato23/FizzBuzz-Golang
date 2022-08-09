package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numRange int
	fmt.Println("Hello")
	fmt.Print("Please enter a non-negative number range to check: ")
	for {
		_, err := fmt.Scanln(&numRange)
		if numRange < 0 || err != nil {
			fmt.Print("You entered in something that was not a non-negative number, please try again: ")
			if err != nil {
				var discard string
				fmt.Scanln(&discard)
			}
		} else {
			break
		}
	}
	for i := 1; i <= numRange; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("FizzBuzz" + " ")
		} else if i % 3 == 0 {
			fmt.Printf("Fizz" + " ")
		} else if i % 5 == 0 {
			fmt.Printf("Buzz" + " ")
		} else {
			concatenated := strconv.Itoa(i) + " "
			fmt.Printf(concatenated)
		}
	}
	

}
