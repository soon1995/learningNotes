package cmdimpl

import "example.com/model"

type LightOnCommand struct {
	light *model.Light
}

func NewLightOnCommand(light *model.Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}
