package commands

import (
	"commandpattern/internal/command"
	"commandpattern/internal/receivers"
	"fmt"
)

// LightOnCommand turns on a light
type LightOnCommand struct {
	light *receivers.Light
	previousState bool
	previousBrightness int
}

func NewLightOnCommand(light *receivers.Light) command.Command {
	return &LightOnCommand{
		light: light,
	}
}

func (c *LightOnCommand) Execute() error {
	// Store previous state for undo
	c.previousState = c.light.IsOn
	c.previousBrightness = c.light.Brightness
	
	c.light.TurnOn()
	return nil
}

func (c *LightOnCommand) Undo() error {
	if c.previousState {
		c.light.SetBrightness(c.previousBrightness)
	} else {
		c.light.TurnOff()
	}
	return nil
}

func (c *LightOnCommand) GetDescription() string {
	return fmt.Sprintf("Turn ON %s light", c.light.Location)
}

func (c *LightOnCommand) CanUndo() bool {
	return true
}

// LightOffCommand turns off a light
type LightOffCommand struct {
	light *receivers.Light
	previousState bool
	previousBrightness int
}

func NewLightOffCommand(light *receivers.Light) command.Command {
	return &LightOffCommand{
		light: light,
	}
}

func (c *LightOffCommand) Execute() error {
	// Store previous state for undo
	c.previousState = c.light.IsOn
	c.previousBrightness = c.light.Brightness
	
	c.light.TurnOff()
	return nil
}

func (c *LightOffCommand) Undo() error {
	if c.previousState {
		c.light.SetBrightness(c.previousBrightness)
	}
	return nil
}

func (c *LightOffCommand) GetDescription() string {
	return fmt.Sprintf("Turn OFF %s light", c.light.Location)
}

func (c *LightOffCommand) CanUndo() bool {
	return true
}

// LightDimCommand sets light brightness
type LightDimCommand struct {
	light *receivers.Light
	brightness int
	previousBrightness int
}

func NewLightDimCommand(light *receivers.Light, brightness int) command.Command {
	return &LightDimCommand{
		light: light,
		brightness: brightness,
	}
}

func (c *LightDimCommand) Execute() error {
	c.previousBrightness = c.light.Brightness
	c.light.SetBrightness(c.brightness)
	return nil
}

func (c *LightDimCommand) Undo() error {
	c.light.SetBrightness(c.previousBrightness)
	return nil
}

func (c *LightDimCommand) GetDescription() string {
	return fmt.Sprintf("Set %s light brightness to %d%%", c.light.Location, c.brightness)
}

func (c *LightDimCommand) CanUndo() bool {
	return true
}