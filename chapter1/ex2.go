package main

//What does go.mod do?
//go.mod is a file used in Go projects to manage dependencies and specify module information. It defines the module's name, its dependencies, and their versions. This allows for better organization and version control of the project's dependencies, making it easier to build and maintain the project.

//Why is module naming important?
//Module naming is important because it provides a unique identifier for the module, which helps avoid conflicts with other modules. A well-chosen module name can also convey the purpose and functionality of the module, making it easier for developers to understand its role in the project. Additionally, using a consistent naming convention can improve code readability and maintainability.

import "fmt"

func Ex2() {
	fmt.Println("Welcome to Go Programming")
}

//What is the difference between go run and go build?
//The difference between go run and go build is that go run compiles and executes the Go program in one step, while go build compiles the Go program into an executable binary file without running it. go run is typically used for quick testing and development, while go build is used for creating a standalone executable that can be distributed and run independently of the Go source code.
