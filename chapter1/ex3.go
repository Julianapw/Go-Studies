package main

import "fmt"

func Ex3() {

	var age int = 23
	var name string = "Juliana"
	var height float64 = 1.58
	var student bool = true

	fmt.Printf("My name is %s, I am %d years old, my height is %.2f meters and it is %t that I am a student.\n", name, age, height, student)
}

//What is static typing?
//Static typing is a programming language feature where the type of a variable is determined at compile time and cannot be changed at runtime. This means that once a variable is declared with a specific type, it can only hold values of that type. Static typing helps catch type-related errors early in the development process, as the compiler can check for type mismatches and enforce type safety. It also allows for better performance, as the compiler can optimize code based on known types.

//How is it different from Python?
//The main difference between static typing in Go and dynamic typing in Python is that in Go, variables must be declared with a specific type, and the type cannot change throughout the program. In contrast, Python allows variables to hold values of any type, and the type can change at runtime. This means that in Python, you can assign a string to a variable and later assign an integer to the same variable without any issues. In Go, you would need to declare separate variables for each type or use interfaces to achieve similar flexibility. Static typing in Go can lead to better performance and early error detection, while dynamic typing in Python offers more flexibility and ease of use.
