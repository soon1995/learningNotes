package display

import (
	"fmt"

	"example.com/subject"
)

type StatisticsDisplay struct {
	maxTemp     float64
	minTemp     float64
	tempSum     float64
	numReadings int
	weatherData *subject.WeatherData
}

func NewStatisticsDisplay(weatherData *subject.WeatherData) *StatisticsDisplay {
	display := &StatisticsDisplay{0.0, 200, 0.0, 0, weatherData}
	weatherData.RegisterObserver(display)
	return display
}

func (d *StatisticsDisplay) Update() {
	temperature := d.weatherData.GetTemperature()

	d.tempSum += temperature
	d.numReadings++

	if temperature > d.maxTemp {
		d.maxTemp = temperature
	}

	if temperature < d.minTemp {
		d.minTemp = temperature
	}
	d.Display()
}

func (d *StatisticsDisplay) Display() {
	fmt.Printf("Avg/Max/Min temperature = %.1f / %.1f / %.1f\n", d.tempSum/float64(d.numReadings), d.maxTemp, d.minTemp)
}
