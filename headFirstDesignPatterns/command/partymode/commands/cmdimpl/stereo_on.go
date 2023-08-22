package cmdimpl

import "example.com/model"

type SterioOnCommand struct {
	stereo *model.Stereo
}

func NewStereoOnCommand(stereo *model.Stereo) *SterioOnCommand {
	return &SterioOnCommand{
		stereo: stereo,
	}
}

func (c *SterioOnCommand) Execute() {
	c.stereo.On()
	c.stereo.SetCD()
	c.stereo.SetVolume(11)
}

func (c *SterioOnCommand) Undo() {
	c.stereo.Off()
}
