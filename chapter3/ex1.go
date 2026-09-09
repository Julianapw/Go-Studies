package main

import "fmt"

type Printer interface {
	Print(message string) error
}

// Exercise 2
type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(message string) error {
	fmt.Println("Printer:", message)
	return nil
}

type FilePrinter struct{}

func (fp FilePrinter) Print(message string) error {
	fmt.Println("FilePrinter:", message)
	return nil
}

var _ Printer = (*ConsolePrinter)(nil)
var _ Printer = (*FilePrinter)(nil) //This line ensures, at compile time, that the type really implements the interface.
