package main

type Adress struct {
	City    string
	Country string
}

type Customer struct {
	Name string
	Adress
}
