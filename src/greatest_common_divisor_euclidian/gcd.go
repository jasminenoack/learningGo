package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func euclidean(a, b int) int {
	if b == 0 {
		return a
	}
	remainder := a % b
	return euclidean(b, remainder)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	string1, _ := reader.ReadString('\n')
	string2, _ := reader.ReadString('\n')
	num1, _ := strconv.Atoi(strings.Trim(string1, " \n"))
	num2, _ := strconv.Atoi(strings.Trim(string2, " \n"))
	if num2 > num1 {
		c := num1
		num1 = num2
		num2 = c
	}
	fmt.Println(euclidean(num1, num2))
}
