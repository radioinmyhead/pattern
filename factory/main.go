package main

import "fmt"

// pizza
type Pizza interface {
	prepare()
	bake()
	cut()
	box()
	Name() string
}

type basePizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (b *basePizza) prepare() {
	fmt.Println("Preparing", b.name)
	fmt.Println("Tossing dough...", b.dough)
	fmt.Println("Adding sauce...", b.sauce)
	fmt.Println("Adding toppings:")
	for _, t := range b.toppings {
		fmt.Println("   " + t)
	}
}
func (b *basePizza) bake()        { fmt.Println("Bake for 25 minutes at 350") }
func (b *basePizza) box()         { fmt.Println("Place pizza in official PizzaStore box") }
func (b *basePizza) cut()         { fmt.Println("Cutting the pizza into diagonal slices") }
func (b *basePizza) Name() string { return b.name }

// pizza 1
type NYStyleCheesePizza struct{ basePizza }

func NewNYStyleCheesePizza() Pizza {
	return &NYStyleCheesePizza{basePizza{
		name:     "NY style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"grated reggiano cheese"},
	}}
}

// pizza 2
type ChicagoStyleCheesePizza struct{ basePizza }

func NewChicagoStyleCheesePizza() Pizza {
	return &ChicagoStyleCheesePizza{basePizza{
		name:     "Chicago style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}}
}

func (p *ChicagoStyleCheesePizza) cut() {
	fmt.Println("Cutting the pizza into square slices")
}

// pizza store
type PizzaCreater interface {
	PizzaCreat(string) Pizza
}
type PizzaStore interface {
	PizzaCreater
	PizzaOrder(string) Pizza
}

type pizzaStore struct{ PizzaCreater }

func (b *pizzaStore) PizzaOrder(item string) Pizza {
	var pizza = b.PizzaCreat(item)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return pizza
}

// store 1
type NYPizzaStore struct{}

func NewNYPizzaStore() PizzaStore {
	return &pizzaStore{&NYPizzaStore{}}
}
func (s *NYPizzaStore) PizzaCreat(item string) Pizza {
	switch item {
	case "cheese":
		return NewNYStyleCheesePizza()
	default:
		panic("why")
	}
}

// store 2
type ChicagoPizzaStore struct{}

func NewChicagoPizzaStore() PizzaStore {
	return &pizzaStore{&ChicagoPizzaStore{}}
}
func (s *ChicagoPizzaStore) PizzaCreat(item string) Pizza {
	switch item {
	case "cheese":
		return NewChicagoStyleCheesePizza()
	default:
		panic("why")
	}
}

// custom
func main() {
	var (
		nyStore      PizzaStore = NewNYPizzaStore()
		chicagoStore PizzaStore = NewChicagoPizzaStore()
		pizza        Pizza
	)

	pizza = nyStore.PizzaOrder("cheese")
	fmt.Println("Ethan ordered a", pizza.Name())

	pizza = chicagoStore.PizzaOrder("cheese")
	fmt.Println("Joel ordered a", pizza.Name())
}
