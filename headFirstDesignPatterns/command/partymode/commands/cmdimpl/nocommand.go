package cmdimpl

type NoCommand struct{}

func NewNoCommand() *NoCommand {
	return &NoCommand{}
}

func (c *NoCommand) Execute() {}

func (c *NoCommand) Undo() {}
