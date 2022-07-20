package main

import "fmt"
	

func main() {

	a := 1
	b := 1
	c := 2
	sum := 0
	n := 4_000_000
	for i := 0; c < n ; i++ {
		if c%2 == 0 {
			sum += c
		}
		a = b
		b = c
		c = a + b
	}
	fmt.Println(sum)
}
