package cmdimpl

import "example.com/model"

type LightOffCommand struct {
	light *model.Light
}

func NewLightOffCommand(light *model.Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}
