package main

import "fmt"

// flyBehavior: fly behavior
type flyBehavior interface {
	fly()
}

type FlyWithWings struct{}

func (f *FlyWithWings) fly() { fmt.Println("fly with wings") }

type FlyNoWay struct{}

func (f *FlyNoWay) fly() { fmt.Println("i can't fly") }

type FlyRocketPowered struct{}

func (f *FlyRocketPowered) fly() { fmt.Println("fly with rocket") }

// quackBehavior: quack behavior
type quackBehavior interface {
	quack()
}

type Quack struct{}

func (q *Quack) quack() { fmt.Println("quack quack") }

type Squeak struct{}

func (q *Squeak) quack() { fmt.Println("squeak squeak") }

type MuteQuack struct{}

func (q *MuteQuack) quack() { fmt.Println("mute quack") }

// Duck
type Ducker interface {
	Display()
	PerformQuack()
	PerformFly()
	Swin()
	SetQuack(quackBehavior)
	SetFly(flyBehavior)
}

type Duck struct {
	flyBehavior   flyBehavior
	quackBehavior quackBehavior
}

func (d *Duck) PerformQuack()             { d.quackBehavior.quack() }
func (d *Duck) PerformFly()               { d.flyBehavior.fly() }
func (d *Duck) Swin()                     { fmt.Println("all ducks float, even decoys!") }
func (d *Duck) SetQuack(qb quackBehavior) { d.quackBehavior = qb }
func (d *Duck) SetFly(fb flyBehavior)     { d.flyBehavior = fb }

// ModelDuck
type ModelDuck struct {
	Duck
}

func NewModelDuck() Ducker {
	return &ModelDuck{
		Duck{
			flyBehavior:   &FlyNoWay{},
			quackBehavior: &Quack{},
		},
	}
}
func (md *ModelDuck) Display() {
	fmt.Println("i am a model duck")
}

// main
func main() {
	md := NewModelDuck()
	md.Display()
	md.PerformQuack()
	md.PerformFly()
	md.SetFly(&FlyRocketPowered{})
	md.PerformFly()
}
