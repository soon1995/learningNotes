package display

import (
	"fmt"

	"example.com/subject"
)

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData subject.WeatherData
}

func NewCurrentConditionsDisplay(weatherData *subject.WeatherData) *CurrentConditionsDisplay {
	display := &CurrentConditionsDisplay{0, 0, *weatherData}
	weatherData.RegisterObserver(display)
	return display
}

func (d *CurrentConditionsDisplay) Update(temperature, humidity, _ float64) {
	d.temperature = temperature
	d.humidity = humidity
	d.Display()
}

func (d *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.1fF degrees and %.1f%% humidity\n", d.temperature, d.humidity)
}
