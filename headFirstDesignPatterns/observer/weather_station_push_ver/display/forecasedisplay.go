package display

import (
	"fmt"

	"example.com/subject"
)

type ForecastDisplay struct {
	currentPressure float64
	lastPressure    float64
}

func NewForecastDisplay(weatherData *subject.WeatherData) *ForecastDisplay {
	display := &ForecastDisplay{29.92, 0}
	weatherData.RegisterObserver(display)
	return display
}

func (d *ForecastDisplay) Update(_, _, pressure float64) {
	d.lastPressure = d.currentPressure
	d.currentPressure = pressure
	d.Display()
}

func (d *ForecastDisplay) Display() {
	fmt.Printf("Forecast: ")
	if d.currentPressure > d.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if d.currentPressure == d.lastPressure {
		fmt.Println("More of the same")
	} else if d.currentPressure < d.lastPressure {
		fmt.Println("Watch out for coolear, rainy weather!")
	}
}
