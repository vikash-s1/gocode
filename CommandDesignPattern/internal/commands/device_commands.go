package commands

import (
	"commandpattern/internal/command"
	"commandpattern/internal/receivers"
	"fmt"
)

// FanOnCommand turns on a fan
type FanOnCommand struct {
	fan *receivers.Fan
	previousState bool
	previousSpeed int
}

func NewFanOnCommand(fan *receivers.Fan) command.Command {
	return &FanOnCommand{
		fan: fan,
	}
}

func (c *FanOnCommand) Execute() error {
	c.previousState = c.fan.IsOn
	c.previousSpeed = c.fan.Speed
	c.fan.TurnOn()
	return nil
}

func (c *FanOnCommand) Undo() error {
	if c.previousState {
		c.fan.SetSpeed(c.previousSpeed)
	} else {
		c.fan.TurnOff()
	}
	return nil
}

func (c *FanOnCommand) GetDescription() string {
	return fmt.Sprintf("Turn ON %s fan", c.fan.Location)
}

func (c *FanOnCommand) CanUndo() bool {
	return true
}

// FanOffCommand turns off a fan
type FanOffCommand struct {
	fan *receivers.Fan
	previousState bool
	previousSpeed int
}

func NewFanOffCommand(fan *receivers.Fan) command.Command {
	return &FanOffCommand{
		fan: fan,
	}
}

func (c *FanOffCommand) Execute() error {
	c.previousState = c.fan.IsOn
	c.previousSpeed = c.fan.Speed
	c.fan.TurnOff()
	return nil
}

func (c *FanOffCommand) Undo() error {
	if c.previousState {
		c.fan.SetSpeed(c.previousSpeed)
	}
	return nil
}

func (c *FanOffCommand) GetDescription() string {
	return fmt.Sprintf("Turn OFF %s fan", c.fan.Location)
}

func (c *FanOffCommand) CanUndo() bool {
	return true
}

// StereoOnCommand turns on stereo with default settings
type StereoOnCommand struct {
	stereo *receivers.Stereo
	previousState bool
	previousVolume int
	previousStation string
}

func NewStereoOnCommand(stereo *receivers.Stereo) command.Command {
	return &StereoOnCommand{
		stereo: stereo,
	}
}

func (c *StereoOnCommand) Execute() error {
	c.previousState = c.stereo.IsOn
	c.previousVolume = c.stereo.Volume
	c.previousStation = c.stereo.Station
	c.stereo.TurnOn()
	return nil
}

func (c *StereoOnCommand) Undo() error {
	if c.previousState {
		c.stereo.SetVolume(c.previousVolume)
		c.stereo.SetStation(c.previousStation)
	} else {
		c.stereo.TurnOff()
	}
	return nil
}

func (c *StereoOnCommand) GetDescription() string {
	return fmt.Sprintf("Turn ON %s stereo", c.stereo.Location)
}

func (c *StereoOnCommand) CanUndo() bool {
	return true
}

// StereoOffCommand turns off stereo
type StereoOffCommand struct {
	stereo *receivers.Stereo
	previousState bool
	previousVolume int
	previousStation string
}

func NewStereoOffCommand(stereo *receivers.Stereo) command.Command {
	return &StereoOffCommand{
		stereo: stereo,
	}
}

func (c *StereoOffCommand) Execute() error {
	c.previousState = c.stereo.IsOn
	c.previousVolume = c.stereo.Volume
	c.previousStation = c.stereo.Station
	c.stereo.TurnOff()
	return nil
}

func (c *StereoOffCommand) Undo() error {
	if c.previousState {
		c.stereo.TurnOn()
		c.stereo.SetVolume(c.previousVolume)
		c.stereo.SetStation(c.previousStation)
	}
	return nil
}

func (c *StereoOffCommand) GetDescription() string {
	return fmt.Sprintf("Turn OFF %s stereo", c.stereo.Location)
}

func (c *StereoOffCommand) CanUndo() bool {
	return true
}

// GarageDoorOpenCommand opens garage door
type GarageDoorOpenCommand struct {
	garageDoor *receivers.GarageDoor
	previousState bool
}

func NewGarageDoorOpenCommand(garageDoor *receivers.GarageDoor) command.Command {
	return &GarageDoorOpenCommand{
		garageDoor: garageDoor,
	}
}

func (c *GarageDoorOpenCommand) Execute() error {
	c.previousState = c.garageDoor.IsOpen
	c.garageDoor.Open()
	return nil
}

func (c *GarageDoorOpenCommand) Undo() error {
	if !c.previousState {
		c.garageDoor.Close()
	}
	return nil
}

func (c *GarageDoorOpenCommand) GetDescription() string {
	return "Open garage door"
}

func (c *GarageDoorOpenCommand) CanUndo() bool {
	return true
}

// GarageDoorCloseCommand closes garage door
type GarageDoorCloseCommand struct {
	garageDoor *receivers.GarageDoor
	previousState bool
}

func NewGarageDoorCloseCommand(garageDoor *receivers.GarageDoor) command.Command {
	return &GarageDoorCloseCommand{
		garageDoor: garageDoor,
	}
}

func (c *GarageDoorCloseCommand) Execute() error {
	c.previousState = c.garageDoor.IsOpen
	c.garageDoor.Close()
	return nil
}

func (c *GarageDoorCloseCommand) Undo() error {
	if c.previousState {
		c.garageDoor.Open()
	}
	return nil
}

func (c *GarageDoorCloseCommand) GetDescription() string {
	return "Close garage door"
}

func (c *GarageDoorCloseCommand) CanUndo() bool {
	return true
}