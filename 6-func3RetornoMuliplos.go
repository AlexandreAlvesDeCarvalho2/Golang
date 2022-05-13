package main

import "fmt"

func multiplosReturns(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := multiplosReturns("hello", "world")
	fmt.Println(a, b)
}
