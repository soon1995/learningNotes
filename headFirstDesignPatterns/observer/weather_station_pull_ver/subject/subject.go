package subject

import (
	"example.com/observer"
)

type Subject interface {
	registerObserver(observer.Observer)
	removeObserver(observer.Observer)
	notifyObserver()
}
