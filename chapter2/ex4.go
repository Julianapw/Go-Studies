package main

func (b *Book) Discount(percent float64) {
	b.Price = b.Price * (1 - percent/100)
}