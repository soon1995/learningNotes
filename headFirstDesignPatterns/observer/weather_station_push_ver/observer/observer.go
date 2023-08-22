package observer

type Observer interface {
	Update(temp, humidity, pressure float64)
}
