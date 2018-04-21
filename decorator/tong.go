package main

type driver interface {
	name() string
	cost() int
	with()
	show()
}

func main() {
	b := NewTee()
	b.show()
	b.with(NewSugar())
	b.show()
}

type Tee struct {
	list []interface{}
}

func NewTee() driver {
	return &Tee{}
}
func (d *Tee) name() string { return "tee" }
func (d *Tee) cost() int    { return 100 }
func (d *Tee) with()        {}
func (d *Tee) show()        {}
