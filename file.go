package main

import "fmt"

var count int = 5
var text string
var number int
var float float32

var a int = 5
var b int = 6

var addition int = a + b
var subtraction int = a - b
var multiplication int = a * b
var division int = a / b

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(count)
	fmt.Println("Enter Your name: ")
	fmt.Scanln(&text)
	fmt.Printf("Hello, %s\n", text )
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)
	fmt.Println(number)
	fmt.Println("Enter a float: ")
	fmt.Scanln(&float)
	fmt.Println(float)
	fmt.Println("addition: ")
	fmt.Println(addition)
	fmt.Println("subtraction: ")
	fmt.Println(subtraction)
	fmt.Println("multiplication: ")
	fmt.Println(multiplication)
	fmt.Println("division: ")
	fmt.Println(division)
	fmt.Println("Reminder:")
	fmt.Println(b % a)
}
