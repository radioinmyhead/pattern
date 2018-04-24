package main

import "fmt"

type namer interface {
	Name() string
}
type Animal interface {
	namer
	Hi()
}

type animal struct{ namer }

func (a *animal) Hi() { fmt.Println("Hi,", a.Name()) }

type dog struct{}

func NewDog() Animal        { return &animal{&dog{}} }
func (a *dog) Name() string { return "doge" }

type cat struct{}

func NewCat() Animal        { return &animal{&cat{}} }
func (a *cat) Name() string { return "kitty" }

func main() {
	var d, c Animal = NewDog(), NewCat()
	d.Hi()
	c.Hi()
}
