package main

type Printable interface {
	Print() string
}

func (b Book) Print() string {
	return b.Title
}
