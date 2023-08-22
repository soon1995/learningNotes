package cmdimpl

import "example.com/model"

type CeilingFanOffCommand struct {
	ceilingFan *model.CeilingFan
	prevSpeed  model.CeilingFanSpeed
}

func NewCeilingFanOffCommand(ceilingFan *model.CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanOffCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Off()
}

func (c *CeilingFanOffCommand) Undo() {
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
