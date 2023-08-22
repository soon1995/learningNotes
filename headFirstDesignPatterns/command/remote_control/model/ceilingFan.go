package model

import "fmt"

type CeilingFanSpeed int

const (
	Off CeilingFanSpeed = iota
	Low
	Medium
	High
)

type CeilingFan struct {
	loc   string
	speed CeilingFanSpeed
}

func NewCeilingFan(loc string) *CeilingFan {
	return &CeilingFan{
		loc:   loc,
		speed: Off,
	}
}

func (c *CeilingFan) High() {
	fmt.Printf("%s ceiling fan is on high\n", c.loc)
	c.speed = High
}

func (c *CeilingFan) Medium() {
	fmt.Printf("%s ceiling fan is on medium\n", c.loc)
	c.speed = Medium
}

func (c *CeilingFan) Low() {
	fmt.Printf("%s ceiling fan is on low\n", c.loc)
	c.speed = Low
}

func (c *CeilingFan) Off() {
	fmt.Printf("%s ceiling fan is off\n", c.loc)
	c.speed = Off
}

func (c *CeilingFan) GetSpeed() CeilingFanSpeed {
	return c.speed
}
