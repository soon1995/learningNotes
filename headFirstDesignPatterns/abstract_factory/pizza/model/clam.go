package model

type FreshClams struct{}

func NewFreshClams() *FreshClams {
	return &FreshClams{}
}

type FrozenClams struct{}

func NewFrozenClams() *FrozenClams {
	return &FrozenClams{}
}
