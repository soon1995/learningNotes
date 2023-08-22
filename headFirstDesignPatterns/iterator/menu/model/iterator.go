package model

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
