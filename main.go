package main

import "fmt"

var name string

//var age int = 22

//var array = [...]string{"Yash", "Krish", "Ravi", "Jaydeep", "Test"}

func main() {
	// Example 1: Print simple message

	fmt.Println("Enter your name ")
	fmt.Scan(&name)

	fmt.Println("Your name is = ", name)
	// Example 2: Print formatted string using variables

	// fmt.Printf("My name is %s and my age is %d", name, age)

	// Example 3: Print star pattern using nested loops

	// for i := 0; i <= 5; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < len(array); i++ {
	// 	fmt.Println("Name =", array[i])
	// }
}
