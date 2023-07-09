package main

import (
	"fmt"

	"example.com/commands/cmdimpl"
	"example.com/model"
	"example.com/remotecontrol"
)

func main() {
	// control := &remotecontrol.SimpleRemoteControl{}
	// light := &model.Light{}
	// lightOnCmd := cmdimpl.NewLightOnCommand(light)

	// control.SetCommand(lightOnCmd)
	// control.ButtenWasPressed()

	// control := remotecontrol.NewRemoteControl()
	// livingRoomLight := model.NewLight("Living Room")
	// kitchenLight := model.NewLight("Kitchen")
	// garageDoor := model.NewGarageDoor("Garage")
	// stereo := model.NewStereo("Living Room")

	// livingRoomLightOn := cmdimpl.NewLightOnCommand(livingRoomLight)
	// livingRoomLightOff := cmdimpl.NewLightOffCommand(livingRoomLight)
	// kitchenLightOn := cmdimpl.NewLightOnCommand(kitchenLight)
	// kitchenLightOff := cmdimpl.NewLightOffCommand(kitchenLight)

	// garageDoorOn := cmdimpl.NewGarageDoorOnCommand(garageDoor)
	// garageDoorOff := cmdimpl.NewGarageDoorOffCommand(garageDoor)

	// stereoOn := cmdimpl.NewStereoOnCommand(stereo)
	// stereoOff := cmdimpl.NewStereoOffCommand(stereo)

	// control.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	// control.SetCommand(1, kitchenLightOn, kitchenLightOff)
	// control.SetCommand(2, garageDoorOn, garageDoorOff)
	// control.SetCommand(3, stereoOn, stereoOff)

	// fmt.Println(control)

	// control.OnButtonWasPushed(0)
	// control.OffButtonWasPushed(0)
	// control.OnButtonWasPushed(1)
	// control.OffButtonWasPushed(1)
	// control.OnButtonWasPushed(2)
	// control.OffButtonWasPushed(2)
	// control.OnButtonWasPushed(3)
	// control.OffButtonWasPushed(3)
	// control.OnButtonWasPushed(5)

	// control := remotecontrol.NewRemoteControl()
	// light := &model.Light{}
	// lightOnCmd := cmdimpl.NewLightOnCommand(light)
	// lightOffCmd := cmdimpl.NewLightOffCommand(light)

	// control.SetCommand(0, lightOnCmd, lightOffCmd)

	// control.OnButtonWasPushed(0)
	// control.OffButtonWasPushed(0)
	// fmt.Println(control)
	// control.UndoButtonWasPushed()
	// control.OffButtonWasPushed(0)
	// control.OnButtonWasPushed(0)
	// fmt.Println(control)
	// control.UndoButtonWasPushed()
	control := remotecontrol.NewRemoteControl()
	ceilingFan := model.NewCeilingFan("Living Room")
	ceilingFanHighCmd := cmdimpl.NewCeilingFanHighCommand(ceilingFan)
	ceilingFanMediumCmd := cmdimpl.NewCeilingFanMediumCommand(ceilingFan)
	ceilingFanOffCmd := cmdimpl.NewCeilingFanOffCommand(ceilingFan)

	control.SetCommand(0, ceilingFanMediumCmd, ceilingFanOffCmd)
	control.SetCommand(1, ceilingFanHighCmd, ceilingFanOffCmd)

	control.OnButtonWasPushed(0)
	control.OffButtonWasPushed(0)
	fmt.Println(control)
	control.UndoButtonWasPushed()

	control.OnButtonWasPushed(1)
	fmt.Println(control)
	control.UndoButtonWasPushed()
}
