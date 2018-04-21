package main

import "fmt"

// interface
type beverager interface {
	getDescription() string
	cost() float64
}

// base
type Beverage struct {
	description string
}

func (b *Beverage) getDescription() string {
	return b.description
}

// b1
type Espresso struct{ Beverage }

func NewEspresso() beverager {
	ret := &Espresso{}
	ret.description = "Espresso"
	return ret
}

func (b *Espresso) cost() float64 {
	return 1.99
}

// b2
type HouseBlend struct{ Beverage }

func NewHouseBlend() beverager {
	ret := &HouseBlend{}
	ret.description = "HouseBlend"
	return ret
}

func (b *HouseBlend) cost() float64 {
	return 0.89
}

// decorator
type Condiment interface {
	getDescription() string
}

type Mocha struct {
	b beverager
}

func NewMocha(b beverager) *Mocha {
	return &Mocha{b: b}
}

func (c *Mocha) getDescription() string {
	return c.b.getDescription() + ", Mocha"
}
func (c *Mocha) cost() float64 {
	return c.b.cost() + 0.2
}

// main
func main() {
	b := NewEspresso()
	fmt.Printf("b=%v,c=%v\n", b.getDescription(), b.cost())

	b = NewEspresso()
	b = NewMocha(b)
	fmt.Printf("b=%v,c=%v\n", b.getDescription(), b.cost())
}
