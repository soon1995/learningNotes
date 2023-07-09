package cmdimpl

import "example.com/model"

type CeilingFanHighCommand struct {
	ceilingFan *model.CeilingFan
	prevSpeed  model.CeilingFanSpeed
}

func NewCeilingFanHighCommand(ceilingFan *model.CeilingFan) *CeilingFanHighCommand {
	return &CeilingFanHighCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanHighCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.High()
}

func (c *CeilingFanHighCommand) Undo() {
	switch c.prevSpeed {
	case model.High:
		c.ceilingFan.High()
	case model.Medium:
		c.ceilingFan.Medium()
	case model.Low:
		c.ceilingFan.Low()
	case model.Off:
		c.ceilingFan.Off()
	}
}
