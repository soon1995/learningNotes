package cmdimpl

import "example.com/model"

type GarageDoorOnCommand struct {
	garageDoor *model.GarageDoor
}

func NewGarageDoorOnCommand(garageDoor *model.GarageDoor) *GarageDoorOnCommand {
	return &GarageDoorOnCommand{
		garageDoor: garageDoor,
	}
}

func (c *GarageDoorOnCommand) Execute() {
	c.garageDoor.On()
}

func (c *GarageDoorOnCommand) Undo() {
	c.garageDoor.Off()
}
