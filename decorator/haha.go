package main

import "fmt"

// interface
type beverager interface {
	getDescription() string
	cost() float64
}

// b1
type Espresso struct{}

func (b *Espresso) getDescription() string {
	return "Espresso"
}

func (b *Espresso) cost() float64 {
	return 1.99
}

// b2
type HouseBlend struct{}

func (b *HouseBlend) getDescription() string {
	return "HouseBlend"
}
func (b *HouseBlend) cost() float64 {
	return 0.89
}

// decorator
type decorator interface {
	beverager
	Set(beverager) beverager
}
type Mocha struct {
	b beverager
}

func NewMocha(b beverager) beverager {
	return &Mocha{b: b}
}
func (c *Mocha) Set(b beverager) beverager {
	c.b = b
	return c
}

func (c *Mocha) getDescription() string {
	return c.b.getDescription() + ", Mocha"
}
func (c *Mocha) cost() float64 {
	return c.b.cost() + 0.2
}

// main
func main() {
	var b beverager = &Espresso{}
	fmt.Printf("b=%v,c=%v\n", b.getDescription(), b.cost())

	b = &Espresso{}
	b = NewMocha(b)
	fmt.Printf("b=%v,c=%v\n", b.getDescription(), b.cost())

	var ret beverager = &Espresso{}
	list := []decorator{&Mocha{}, &Mocha{}}
	for _, d := range list {
		ret = d.Set(ret)
		// ret = ret.set(d)
	}
	fmt.Printf("b=%v,c=%v\n", ret.getDescription(), ret.cost())
}
