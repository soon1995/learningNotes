package main

import (
	"example.com/commands"
	"example.com/commands/cmdimpl"
	"example.com/model"
	"example.com/remotecontrol"
)

func main() {
	light := model.NewLight("Living Room")
	stereo := model.NewStereo("Living Room")

	lightOn := cmdimpl.NewLightOnCommand(light)
	lightOff := cmdimpl.NewLightOnCommand(light)
	stereoOn := cmdimpl.NewStereoOnCommand(stereo)
	stereoOff := cmdimpl.NewStereoOffCommand(stereo)

	partyOn := []commands.Command{lightOn, stereoOn}
	partyOff := []commands.Command{lightOff, stereoOff}

	partyOnMacro := cmdimpl.NewMacroCommand(partyOn)
	partyOffMacro := cmdimpl.NewMacroCommand(partyOff)

	remoteConrtol := remotecontrol.NewRemoteControl()
	remoteConrtol.SetCommand(0, partyOnMacro, partyOffMacro)
	remoteConrtol.OnButtonWasPushed(0)
}
