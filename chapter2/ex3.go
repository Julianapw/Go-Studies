package main

import "fmt"

type Book struct {
	Title   string
	Author  string
	Price   float64
	Premium bool
}

//Exercise 6

func NewBook(title string, author string, price float64) *Book {
	if price < 0 {
		price = 0
	}

	return &Book{
		Title:  title,
		Author: author,
		Price:  price,
	}
}

//Exercise 7

func isExpensive(book *Book) bool {
	if book.Price > 1000 {
		return true
	}
	return false
}

// Exercise 13
func UpdatePrice(book *Book, newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("Price cannot be negative")
	}
	book.Price = newPrice
	return nil
}

// Exercise 14
func TotalPrice(books []Book) float64 {
	var total float64
	for _, book := range books {
		total += book.Price
	}
	return total
}

// Exercise 15
func IncreaseAllPrices(books []Book, percent float64) {
	for i := range books {
		books[i].Price = books[i].Price * (1 + percent/100)
	}
}

// Exercise 16
func (book Book) GetPrice() float64 {
	return book.Price
}

// Exercise 17
func (book *Book) SetPrice(newPrice float64) *Book {
	if newPrice >= 0 {
		book.Price = newPrice
	}
	return book
}

// Exercise 18
func NewPremiumBook(title string, author string, price float64) *Book {
	return &Book{
		Title:   title,
		Author:  author,
		Price:   price,
		Premium: true,
	}
}

type Cart struct {
	Books []Book
}

func (c *Cart) AddBook(book Book) {
	c.Books = append(c.Books, book)
}

func (c *Cart) TotalPrice() float64 {
	return TotalPrice(c.Books)
}
