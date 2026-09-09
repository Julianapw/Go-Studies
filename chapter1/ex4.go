package main

import "fmt"

func Ex4() {

	age := 23
	name := "Juliana"
	height := 1.58
	student := true

	fmt.Printf("My name is %s, I am %d years old, my height is %.2f meters and it is %t that I am a student.\n", name, age, height, student)
}

//When can you not use :=?
//Outside of a function, when the variable is already declared in the same scope, when assigning to an existing variable (not declaring), when declaring struct fields, package-level constants, or imports
