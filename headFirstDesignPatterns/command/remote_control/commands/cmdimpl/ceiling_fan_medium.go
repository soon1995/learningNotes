package cmdimpl

import "example.com/model"

type CeilingFanMediumCommand struct {
	ceilingFan *model.CeilingFan
	prevSpeed  model.CeilingFanSpeed
}

func NewCeilingFanMediumCommand(ceilingFan *model.CeilingFan) *CeilingFanMediumCommand {
	return &CeilingFanMediumCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanMediumCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Medium()
}

func (c *CeilingFanMediumCommand) Undo() {
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
