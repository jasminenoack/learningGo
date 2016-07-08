package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Echo1 test
func Echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args[1:]); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

// Echo2 test
func Echo2(args []string) string {
	var s, sep string
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// Echo3 print all
func Echo3(args []string) string {
	var s, sep string
	for _, arg := range args[:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// Echo4 modify to print the index and value of each argument
func Echo4(args []string) string {
	var s, sep string
	var indexSep = ":"
	for index, arg := range args[1:] {
		s += sep + strconv.Itoa(index) + indexSep + arg
		sep = " "
	}
	return s
}

// Echo5 more efficent version
func Echo5(args []string) string {
	return strings.Join(args[1:], " ")
}

// Echo6 just print array
func Echo6(args []string) []string {
	return args[1:]
}

func main() {
	arguments := os.Args
	fmt.Println("Echo1")
	result := Echo1(arguments)
	fmt.Println(result)
	fmt.Println("Echo2")
	result = Echo2(arguments)
	fmt.Println(result)
	fmt.Println("Echo3")
	result = Echo3(arguments)
	fmt.Println(result)
	fmt.Println("Echo4")
	result = Echo4(arguments)
	fmt.Println(result)
	fmt.Println("Echo5")
	result = Echo5(arguments)
	fmt.Println(result)
	fmt.Println("Echo6")
	result2 := Echo6(arguments)
	fmt.Println(result2)
}
