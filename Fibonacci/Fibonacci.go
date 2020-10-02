/*---------------------------------------------------------------------
Shirish Amatya'
10/01/2020
This Program iterates the Fibonacci Sequence for the first 25 numbers
using a recursive function
---------------------------------------------------------------------*/
package main

import (
	"fmt"
	"strconv"
)

func FibonacciRec(num int) int {
	if num <= 1 {
		return num
	}
	return FibonacciRec(num-1) + FibonacciRec(num-2)
}

func main() {
	for i := 0; i <= 25; i++ {
		fmt.Print(strconv.Itoa(FibonacciRec(i)) + " ")
	}
	fmt.Println("")
}
