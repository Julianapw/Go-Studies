package main

import "fmt"

func Ex5() {
	var a int
	var b string
	var c bool

	fmt.Printf("The value of a is %d, the value of b is %s and the value of c is %t.\n", a, b, c)
}

// What are zero values?
// Zero values are the default values assigned to variables in Go when they are declared without an explicit initialization. Each type has its own zero value: for numeric types (int, float, etc.) it is 0, for strings it is an empty string "", for booleans it is false, and for pointers, interfaces, slices, channels, maps, and function types it is nil. Zero values allow developers to work with uninitialized variables without encountering undefined behavior, as they will have a predictable default state.

//Why is this useful?
//Zero values are useful because they provide a safe and predictable default state for variables that have not been explicitly initialized. This helps prevent bugs and errors that can arise from using uninitialized variables, as they will have a known value instead of containing garbage data. Additionally, zero values can simplify code by allowing developers to rely on default values without needing to explicitly set them, which can lead to cleaner and more concise code.
