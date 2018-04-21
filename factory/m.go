package main

import "fmt"

type namer interface {
	name() string
}
type animal interface {
	namer
	hi()
}

type Animal struct{ namer }

func (b *Animal) hi() { fmt.Println("hi,", b.name()) }

type dog struct{}

func (a *dog) name() string { return "doge" }
func NewDog() animal        { return &Animal{&dog{}} }

type cat struct{}

func (a *cat) name() string { return "kitty" }
func NewCat() animal        { return &Animal{&cat{}} }

func main() {
	var d, c animal = NewDog(), NewCat()
	d.hi()
	c.hi()
}
