package cmdimpl

import "example.com/commands"

type MacroCommand struct {
	cmds []commands.Command
}

func NewMacroCommand(cmds []commands.Command) *MacroCommand {
	return &MacroCommand{cmds}
}

func (c *MacroCommand) Execute() {
	for _, cmd := range c.cmds {
		cmd.Execute()
	}
}

func (c *MacroCommand) Undo() {
  for i := len(c.cmds); i >= 0; i-- {
    c.cmds[i].Undo()
  }
}
