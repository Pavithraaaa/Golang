package main

import "fmt"

func main() {
	var y string
	fmt.Println("Enter any word")
	fmt.Scanln(&y)
	for i := 0; i < len(y)/2; i++ {
		if y[i] != y[len(y)-i-1] {
			fmt.Println("The entered string is not a palindrome")
			break

		} else {
			fmt.Println("Is a palindrome")
			break
		}
	}
}
