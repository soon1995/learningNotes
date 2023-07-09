package subject

import "example.com/observer"

type WeatherData struct {
	observers   []observer.Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (s *WeatherData) RegisterObserver(o observer.Observer) {
	s.observers = append(s.observers, o)
}

func (s *WeatherData) RemoveObserver(o observer.Observer) {
	for i := 0; i < len(s.observers); i++ {
		if s.observers[i] == o {
			copy(s.observers[i:], s.observers[i+1:])
			s.observers = s.observers[:len(s.observers)-1]
			return
		}
	}
}

func (s *WeatherData) NotifyObservers() {
	for _, v := range s.observers {
		v.Update()
	}
}

func (s *WeatherData) MeasurementsChanged() {
	s.NotifyObservers()
}

func (s *WeatherData) SetMeasurement(temperature, humidity, pressure float64) {
	s.temperature = temperature
	s.humidity = humidity
	s.pressure = pressure
	s.MeasurementsChanged()
}

func (s *WeatherData) GetTemperature() float64 {
	return s.temperature
}

func (s *WeatherData) GetHumidity() float64 {
	return s.humidity
}

func (s *WeatherData) GetPressure() float64 {
	return s.pressure
}
