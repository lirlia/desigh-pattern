package main

import "fmt"

func main() {
	p := Factory(1)
	fmt.Println(p.GetPrice())
}

type ProductIF interface {
	GetPrice() int
}

type Product struct {
	Price int
}

func (p *Product) GetPrice() int {
	return p.Price
}

func NewProduct(price int) ProductIF {
	return &Product{Price: price}
}

func Factory(n int) ProductIF {
	return NewProduct(100)
}
