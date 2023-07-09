package main

import (
	"example.com/display"
	"example.com/subject"
)

func main() {
	weatherData := subject.NewWeatherData()

	display.NewCurrentConditionsDisplay(weatherData)
	display.NewStatisticsDisplay(weatherData)
	display.NewForecastDisplay(weatherData)
	display.NewHeatIndexDisplay(weatherData)

	weatherData.SetMeasurement(80, 65, 30.4)
	weatherData.SetMeasurement(82, 70, 29.2)
	weatherData.SetMeasurement(78, 90, 29.2)
}
