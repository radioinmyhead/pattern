package main

import "fmt"

// fly: fly behavior
type fly func()

var (
	FlyWithWings     fly = func() { fmt.Println("fly with wings") }
	FlyNoWay         fly = func() { fmt.Println("i can't fly") }
	FlyRocketPowered fly = func() { fmt.Println("fly with rocket") }
)

// quack: quack behavior
type quack func()

var (
	Quack     quack = func() { fmt.Println("quack quack") }
	Squeak    quack = func() { fmt.Println("squeak squeak") }
	MuteQuack quack = func() { fmt.Println("mute quack") }
)

// Duck
type Duck struct {
	fly   fly
	quack quack
}

func PerformQuack(d *Duck)      { d.quack() }
func PerformFly(d *Duck)        { d.fly() }
func Swin(d *Duck)              { fmt.Println("all ducks float, even decoys!") }
func SetQuack(d *Duck, q quack) { d.quack = q }
func SetFly(d *Duck, f fly)     { d.fly = f }

// ModelDuck
type ModelDuck struct {
	Duck
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{
		Duck{
			fly:   FlyNoWay,
			quack: Quack,
		},
	}
}
func ModelDuckDisplay(md *ModelDuck) {
	fmt.Println("i am a model duck")
}

// main
func main() {
	md := NewModelDuck()
	ModelDuckDisplay(md)
	PerformQuack(&md.Duck)
	PerformFly(&md.Duck)
	SetFly(&md.Duck, FlyRocketPowered)
	PerformFly(&md.Duck)
}
