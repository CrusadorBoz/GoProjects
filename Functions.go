// how to create functions in go
package main

import (
	"fmt"
)

func addtwopoints(a int, b int) int {
	return a + b
}

func addseries(a, b, c int) int {
	return a + b + c
}

func main() {
	var tot int = 0
	// or you can tot := addtwopoints(1,3)
	tot = addtwopoints(14, 33)
	// Example of :=  declares var when it is not already
	sumOfTwo := addtwopoints(1, 4)
	fmt.Println("14 + 33 = ", tot)
	fmt.Println("1 + 4 = ", sumOfTwo)
	sumSeries := addseries(4, 6, 8)
	fmt.Println("Sum of series ", sumSeries)
}
