package display

import (
	"fmt"

	"example.com/subject"
)

type HeatIndexDisplay struct {
	heatindex   float64
	weatherData *subject.WeatherData
}

func NewHeatIndexDisplay(weatherData *subject.WeatherData) *HeatIndexDisplay {
	display := &HeatIndexDisplay{0, weatherData}
	weatherData.RegisterObserver(display)
	return display
}

func (d *HeatIndexDisplay) Update() {
	d.heatindex = d.computeHeatIndex(d.weatherData.GetTemperature(), d.weatherData.GetHumidity())
	d.Display()
}

func (d *HeatIndexDisplay) Display() {
	fmt.Printf("Heat Index is %.5f\n", d.heatindex)
}

func (d *HeatIndexDisplay) computeHeatIndex(t, rh float64) float64 {
	index := (16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) +
		(0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) +
		(0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) +
		(0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 *
		(rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) +
		(0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) +
		0.000000000843296*(t*t*rh*rh*rh)) -
		(0.0000000000481975 * (t * t * t * rh * rh * rh))
	return index
}
