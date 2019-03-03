package main

import "fmt"

func main() {

	var ab [4]string

	ab[0] = "first"
	ab[1] = "second"
	ab[2] = "third"
	ab[3] = "fourth"

	fmt.Println(ab[0], ab[1], ab[2])
	fmt.Println(ab)
}
