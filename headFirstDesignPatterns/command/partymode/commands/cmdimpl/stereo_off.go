package cmdimpl

import "example.com/model"

type SterioOffCommand struct {
	stereo *model.Stereo
}

func NewStereoOffCommand(stereo *model.Stereo) *SterioOffCommand {
	return &SterioOffCommand{
		stereo: stereo,
	}
}

func (c *SterioOffCommand) Execute() {
	c.stereo.Off()
}

func (c *SterioOffCommand) Undo() {
	c.stereo.On()
}
