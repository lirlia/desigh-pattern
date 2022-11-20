package main

import "fmt"

func main() {
	p := NewProduct(WithName("iPhone"), WithColor("Black"), WithId(1))
	fmt.Println(p)

	p = NewProduct(WithName("Android"), WithId(2), WithColor("White"))
	fmt.Println(p)
}

type Option func(*product)

type product struct {
	name  string
	id    int
	color string
}

func NewProduct(opts ...Option) *product {

	p := &product{}
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithName(name string) Option {
	return func(p *product) {
		p.name = name
	}
}

func WithColor(color string) Option {
	return func(p *product) {
		p.color = color
	}
}

func WithId(id int) Option {
	return func(p *product) {
		p.id = id
	}
}
