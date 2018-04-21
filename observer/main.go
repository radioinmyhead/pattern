package main

import "fmt"

type Subjecter interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObserver()
	SetMeasurements(float64, float64, float64)
}

type Observer interface {
	Update(temp, humidity, pressure float64)
}

type DisplayElementer interface {
	Display()
}

//
type data struct {
	temp     float64
	humidity float64
	pressure float64
}

type WeatherData struct {
	data
	list []Observer
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.list = append(w.list, o)
}
func (w *WeatherData) RemoveObserver(o Observer) {
	for i, ob := range w.list {
		if ob == o {
			w.list = append(w.list[:i], w.list[i+1:]...)
		}
	}
}
func (w *WeatherData) NotifyObserver() {
	for _, o := range w.list {
		o.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *WeatherData) measurementsChanged() {
	w.NotifyObserver()
}
func (w *WeatherData) SetMeasurements(t, h, p float64) {
	w.temp = t
	w.humidity = h
	w.pressure = p
	w.measurementsChanged()
}

type CurrentConditionsDisplay struct {
	temp        float64
	humidity    float64
	weatherData *WeatherData
}

func (d *CurrentConditionsDisplay) Update(t, h, p float64) {
	d.temp = t
	d.humidity = h
	d.Display()
}
func (d *CurrentConditionsDisplay) Display() {
	fmt.Printf("current conditions: %vF degress and %v%% humidity\n", d.temp, d.humidity)
}

func main() {
	// init
	var weatherData Subjecter = &WeatherData{}
	var currentDisplay DisplayElementer = &CurrentConditionsDisplay{}

	// register observer
	weatherData.RegisterObserver(currentDisplay.(Observer))

	// change
	currentDisplay.Display()
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)

	// remove
	weatherData.RemoveObserver(currentDisplay.(Observer))
	weatherData.SetMeasurements(80, 65, 30.4)
}
