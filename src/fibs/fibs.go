// fibs.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fibs := []int{0, 1}
	if len(os.Args) > 1 {
		numberString := os.Args[1]
		number, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
		for len(fibs) < number {
			fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2])
		}
		fmt.Println(fibs[number-1])
	} else {
		fmt.Println("But what fib do you want my dear?")
	}
}
