package main

import "fmt"

// pizza ingredient
const (
	FreshClams  = "Fresh Clams"
	FrozenClams = "Frozen Clams"

	MarinaraSauce   = "Marinara Sauce"
	PlumTomatoSauce = "Plum Tomato Sauce"

	ThinCrushDough  = "Thin Crush Dough"
	ThickCrustDough = "Thick Crust Dough"

	ReggianoCheese   = "Reggiano Cheese"
	MozzarellaCheese = "Mozzarella Cheese"
	GoatCheese       = "Goat Cheese"

	SlicedPepperoni = "Sliced Pepperoni"

	Garlic    = "Garlic"
	Onion     = "Onion"
	RedPepper = "Red Pepper"
	Mushroom  = "Mushroom"
)

// pizza ingredient factory
type PizzaIngredientFactory interface {
	createClam() string
	createSauce() string
	createDouth() string
	createCheese() string
	createPepperoni() string
	createVeggies() []string
}

type NYPizzaIngredientFactory struct{}

func (f *NYPizzaIngredientFactory) createClam() string      { return FreshClams }
func (f *NYPizzaIngredientFactory) createSauce() string     { return MarinaraSauce }
func (f *NYPizzaIngredientFactory) createDouth() string     { return ThinCrushDough }
func (f *NYPizzaIngredientFactory) createCheese() string    { return ReggianoCheese }
func (f *NYPizzaIngredientFactory) createPepperoni() string { return SlicedPepperoni }
func (f *NYPizzaIngredientFactory) createVeggies() []string {
	return []string{Garlic, Onion, RedPepper, Mushroom}
}

type ChicagoPizzaIngredientFactory struct{}

func (f *ChicagoPizzaIngredientFactory) createClam() string      { return FrozenClams }
func (f *ChicagoPizzaIngredientFactory) createSauce() string     { return PlumTomatoSauce }
func (f *ChicagoPizzaIngredientFactory) createDouth() string     { return ThickCrustDough }
func (f *ChicagoPizzaIngredientFactory) createCheese() string    { return MozzarellaCheese }
func (f *ChicagoPizzaIngredientFactory) createPepperoni() string { return SlicedPepperoni }
func (f *ChicagoPizzaIngredientFactory) createVeggies() []string {
	return []string{Garlic, Onion, RedPepper, Mushroom}
}

// pizza
type Pizza interface {
	prepare()
	bake()
	cut()
	box()
	call(string)
	Name() string
}

type basePizza struct {
	name              string
	dough             string
	sauce             string
	veggies           []string
	cheese            string
	pepperoni         string
	clam              string
	ingredientFactory PizzaIngredientFactory
}

func (b *basePizza) bake()         { fmt.Println("Bake for 25 minutes at 350") }
func (b *basePizza) cut()          { fmt.Println("Cutting the pizza into diagonal slices") }
func (b *basePizza) box()          { fmt.Println("Place pizza in official PizzaStore box") }
func (b *basePizza) call(n string) { b.name = n }
func (b *basePizza) Name() string  { return b.name }
func (b *basePizza) String() string {
	return "cheese=" + b.cheese + ",dough=" + b.dough + ",sauce=" + b.sauce
}

// cheese pizza
type CheesePizza struct{ basePizza }

func NewCheesePizza(i PizzaIngredientFactory) Pizza {
	return &CheesePizza{basePizza{ingredientFactory: i}}
}
func (p *CheesePizza) prepare() {
	fmt.Println("prepare", p.name)
	p.dough = p.ingredientFactory.createDouth()
	p.sauce = p.ingredientFactory.createSauce()
	p.cheese = p.ingredientFactory.createCheese()
}

// clam pizza
type ClamPizza struct{ basePizza }

func NewClamPizza(i PizzaIngredientFactory) Pizza {
	return &ClamPizza{basePizza{ingredientFactory: i}}
}
func (p *ClamPizza) prepare() {
	fmt.Println("prepare", p.name)
	p.dough = p.ingredientFactory.createDouth()
	p.sauce = p.ingredientFactory.createSauce()
	p.cheese = p.ingredientFactory.createCheese()
	p.clam = p.ingredientFactory.createClam()
}

// pizza store
type PizzaCreater interface {
	CreatePizza(int) Pizza
}
type PizzaStore interface {
	PizzaCreater
	OrderPizza(int) Pizza
}

type pizzaStore struct{ PizzaCreater }

func (b *pizzaStore) OrderPizza(item int) Pizza {
	var pizza = b.CreatePizza(item)
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
func (s *NYPizzaStore) CreatePizza(item int) (pizza Pizza) {
	i := &NYPizzaIngredientFactory{}
	switch item {
	case PTCheese:
		pizza = NewCheesePizza(i)
		pizza.call("New York Style cheese Pizza")
	case PTClam:
		pizza = NewClamPizza(i)
		pizza.call("New York Style clam Pizza")
	default:
		fmt.Println("not support")
		panic("why")
	}
	return
}

// store 2
type ChicagoPizzaStore struct{}

func NewChicagoPizzaStore() PizzaStore {
	return &pizzaStore{&ChicagoPizzaStore{}}
}
func (s *ChicagoPizzaStore) CreatePizza(item int) (pizza Pizza) {
	i := &ChicagoPizzaIngredientFactory{}
	switch item {
	case PTCheese:
		pizza = NewCheesePizza(i)
		pizza.call("Chicago Style cheese Pizza")
	default:
		fmt.Println("not support this type pizza")
		panic("why")
	}
	return
}

const (
	AreaYN int = iota
	AreaChicago
	PTCheese
	PTClam
)

func CreatePizzaStore(t int) PizzaStore {
	switch t {
	case AreaYN:
		return NewNYPizzaStore()
	case AreaChicago:
		return NewChicagoPizzaStore()
	default:
		fmt.Println("this area not support")
		return nil
	}
}

func main() {
	var (
		nyStore      PizzaStore = CreatePizzaStore(AreaYN)
		chicagoStore PizzaStore = CreatePizzaStore(AreaChicago)
		pizza        Pizza
	)

	pizza = nyStore.OrderPizza(PTCheese)
	fmt.Println(pizza)
	fmt.Printf("Ethan ordered a %s\n\n", pizza.Name())

	pizza = nyStore.OrderPizza(PTClam)
	fmt.Println(pizza)
	fmt.Printf("Alex ordered a %s\n\n", pizza.Name())

	pizza = chicagoStore.OrderPizza(PTCheese)
	fmt.Println(pizza)
	fmt.Printf("Joel ordered a %s\n\n", pizza.Name())
}
