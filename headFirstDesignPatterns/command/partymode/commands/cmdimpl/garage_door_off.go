package cmdimpl

import "example.com/model"

type GarageDoorOffCommand struct {
	garageDoor *model.GarageDoor
}

func NewGarageDoorOffCommand(garageDoor *model.GarageDoor) *GarageDoorOffCommand {
	return &GarageDoorOffCommand{
		garageDoor: garageDoor,
	}
}

func (c *GarageDoorOffCommand) Execute() {
	c.garageDoor.Off()
}

func (c *GarageDoorOffCommand) Undo() {
	c.garageDoor.On()
}
