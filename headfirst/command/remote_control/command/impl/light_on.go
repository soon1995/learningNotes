package impl

type LightOnCommand struct {
  light Light
}

func NewLightOnCommand(light Light) *LightOnCommand{
  return &LightOnCommand{
    light: light,
  }
}

func (c *LightOnCommand) Execute() {
  c.light.On()
}
