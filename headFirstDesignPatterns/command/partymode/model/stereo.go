package model

import "fmt"

type Stereo struct {
	name string
}

func NewStereo(name string) *Stereo {
	return &Stereo{
		name,
	}
}

func (m *Stereo) On() {
	fmt.Printf("Stereo %s is On\n", m.name)
}

func (m *Stereo) SetCD() {
	fmt.Printf("Stereo %s CD is set\n", m.name)
}

func (m *Stereo) SetVolume(v int) {
	fmt.Printf("Stereo %s volume is set to %d\n", m.name, v)
}

func (m *Stereo) Off() {
	fmt.Printf("Stereo %s is Off\n", m.name)
}
