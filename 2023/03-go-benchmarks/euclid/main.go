package main

import "fmt"

// Euclid implements a version of Euclids algorithm do find the greatest common divisor of two numbers.
// Pseudo code from Wikipedia:
// function gcd(a, b)
//
//	while b ≠ 0
//	    t := b
//	    b := a mod b
//	    a := t
//	return a
func Euclid(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func main() {
	a := 252
	b := 105
	fmt.Printf("GCD of %d and %d is %d.\n", a, b, Euclid(a, b))
}
