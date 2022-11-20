package main

import (
	"fmt"
	"sync"
)

type Cart struct {
	item *sync.Map
}

func main() {
	c := NewCart()
	fmt.Println(c.GetItem("item1"))
	c.AddItem("item1")
	fmt.Println(c.GetItem("item1"))
}

func NewCart() *Cart {
	return &Cart{
		item: &sync.Map{},
	}
}
func (c *Cart) AddItem(item string) {
	c.item.Store(item, 1)
}

func (c *Cart) GetItem(item string) int {
	if v, ok := c.item.Load(item); ok {
		return v.(int)
	}
	return 0
}

func (c *Cart) DeleteItem(item string) {
	c.item.Delete(item)
}
