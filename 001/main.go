package main 

import "fmt"


func main() {

	sum := 0
	for i := 0; i < 1000 ; i++ {
		if i % 3 == 0  || i % 5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of all mutiple of 3 or 5 below 1000 is: ", sum)
}