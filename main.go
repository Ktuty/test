package main

import "fmt"

func main() {
	var v1, v2 int

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&v1)
	if err != nil {
		fmt.Printf("Error reading user input: %s\n", err)
		panic(err)
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&v2)
	if err != nil {
		fmt.Printf("Error reading user input: %s\n", err)
		panic(err)
	}

	fmt.Printf("First number: %d\nSecond number: %d\n\nSum = %v", v1, v2, v1+v2)

}
