package remotecontrol

import (
	"bytes"
	"fmt"
	"reflect"

	"example.com/commands"
	"example.com/commands/cmdimpl"
)

type SimpleRemoteControl struct {
	slot commands.Command
}

func (c *SimpleRemoteControl) SetCommand(command commands.Command) {
	c.slot = command
}

func (c *SimpleRemoteControl) ButtenWasPressed() {
	c.slot.Execute()
}

type RemoteControl struct {
	onCommands  []commands.Command
	offCommands []commands.Command
	undoCommand commands.Command
}

func NewRemoteControl() *RemoteControl {
	remote := &RemoteControl{
		onCommands:  make([]commands.Command, 7, 7),
		offCommands: make([]commands.Command, 7, 7),
	}
	nocommand := cmdimpl.NewNoCommand()
	for i := 0; i < 7; i++ {
		remote.onCommands[i] = nocommand
		remote.offCommands[i] = nocommand
	}
	remote.undoCommand = nocommand
	return remote
}

func (c *RemoteControl) SetCommand(slot int, onCommand, offCommand commands.Command) {
	c.onCommands[slot] = onCommand
	c.offCommands[slot] = offCommand
}

// we dont do if else check nil, we provide null object when init
func (c *RemoteControl) OnButtonWasPushed(slot int) {
	c.onCommands[slot].Execute()
	c.undoCommand = c.onCommands[slot]
}

func (c *RemoteControl) OffButtonWasPushed(slot int) {
	c.offCommands[slot].Execute()
	c.undoCommand = c.offCommands[slot]
}

func (c *RemoteControl) UndoButtonWasPushed() {
	c.undoCommand.Undo()
}

func (c *RemoteControl) String() string {
	var buf bytes.Buffer
	buf.WriteString("\n---- Remote Control ----\n")
	for i, command := range c.onCommands {
		onName := "none"
		offName := "none"
		if command != nil {
			onName = reflect.TypeOf(command).Elem().Name()
		}
		if c.offCommands[i] != nil {
			offName = reflect.TypeOf(c.offCommands[i]).Elem().Name()
		}
		buf.WriteString(fmt.Sprintf("[slot %d]\t%s\t\t%s\n", i, onName, offName))
	}
	buf.WriteString(fmt.Sprintf("[undo] %s\n", reflect.TypeOf(c.undoCommand).Elem().Name()))
	return buf.String()
}
