package main

import (
	"commandpattern/internal/commands"
	"commandpattern/internal/command"
	"commandpattern/internal/invoker"
	"commandpattern/internal/receivers"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Command Design Pattern Demo ===")
	fmt.Println("Smart Home Automation System")
	fmt.Println(strings.Repeat("=", 50))

	// Create devices (receivers)
	livingRoomLight := receivers.NewLight("Living Room")
	kitchenLight := receivers.NewLight("Kitchen")
	bedroomFan := receivers.NewFan("Bedroom")
	livingRoomStereo := receivers.NewStereo("Living Room")
	garageDoor := receivers.NewGarageDoor()

	// Create remote control (invoker)
	remote := invoker.NewRemoteControl()

	fmt.Println("\n1. Setting up Remote Control Commands")
	fmt.Println(strings.Repeat("-", 40))

	// Create commands
	livingRoomLightOn := commands.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := commands.NewLightOffCommand(livingRoomLight)
	kitchenLightOn := commands.NewLightOnCommand(kitchenLight)
	kitchenLightOff := commands.NewLightOffCommand(kitchenLight)
	fanOn := commands.NewFanOnCommand(bedroomFan)
	fanOff := commands.NewFanOffCommand(bedroomFan)
	stereoOn := commands.NewStereoOnCommand(livingRoomStereo)
	stereoOff := commands.NewStereoOffCommand(livingRoomStereo)
	garageOpen := commands.NewGarageDoorOpenCommand(garageDoor)
	garageClose := commands.NewGarageDoorCloseCommand(garageDoor)

	// Set up remote control slots
	remote.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remote.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remote.SetCommand(2, fanOn, fanOff)
	remote.SetCommand(3, stereoOn, stereoOff)
	remote.SetCommand(4, garageOpen, garageClose)

	// Show remote configuration
	remote.ShowConfiguration()

	fmt.Println("\n2. Testing Basic Commands")
	fmt.Println(strings.Repeat("-", 40))

	// Test basic commands
	remote.OnButtonPressed(0)  // Living room light on
	remote.OnButtonPressed(1)  // Kitchen light on
	remote.OnButtonPressed(2)  // Fan on
	remote.OnButtonPressed(3)  // Stereo on

	fmt.Println("\n3. Testing Undo Functionality")
	fmt.Println(strings.Repeat("-", 40))

	remote.UndoButtonPressed() // Undo stereo on
	remote.OffButtonPressed(2) // Fan off
	remote.UndoButtonPressed() // Undo fan off

	fmt.Println("\n4. Testing Light Dimming")
	fmt.Println(strings.Repeat("-", 40))

	dimCommand := commands.NewLightDimCommand(livingRoomLight, 30)
	remote.ExecuteCommand(dimCommand)
	remote.UndoButtonPressed() // Undo dim

	fmt.Println("\n5. Testing Macro Commands")
	fmt.Println(strings.Repeat("-", 40))

	// Create a "Party Mode" macro
	partyCommands := []command.Command{
		commands.NewLightOnCommand(livingRoomLight),
		commands.NewLightOnCommand(kitchenLight),
		commands.NewLightDimCommand(livingRoomLight, 70),
		commands.NewStereoOnCommand(livingRoomStereo),
	}
	partyMacro := commands.NewMacroCommand("Party Mode", partyCommands)

	// Create a "Sleep Mode" macro
	sleepCommands := []command.Command{
		commands.NewLightOffCommand(livingRoomLight),
		commands.NewLightOffCommand(kitchenLight),
		commands.NewFanOffCommand(bedroomFan),
		commands.NewStereoOffCommand(livingRoomStereo),
		commands.NewGarageDoorCloseCommand(garageDoor),
	}
	sleepMacro := commands.NewMacroCommand("Sleep Mode", sleepCommands)

	// Execute macros
	remote.ExecuteCommand(partyMacro)
	
	fmt.Println("\n" + strings.Repeat("-", 30))
	
	remote.ExecuteCommand(sleepMacro)
	
	fmt.Println("\n6. Testing Macro Undo")
	fmt.Println(strings.Repeat("-", 40))
	
	remote.UndoButtonPressed() // Undo sleep mode

	fmt.Println("\n7. Command History")
	fmt.Println(strings.Repeat("-", 40))
	
	remote.ShowHistory()

	fmt.Println("\n8. Device Status Check")
	fmt.Println(strings.Repeat("-", 40))
	
	fmt.Println("📊 Current Device Status:")
	fmt.Printf("- %s\n", livingRoomLight.GetStatus())
	fmt.Printf("- %s\n", kitchenLight.GetStatus())
	fmt.Printf("- %s\n", bedroomFan.GetStatus())
	fmt.Printf("- %s\n", livingRoomStereo.GetStatus())
	fmt.Printf("- %s\n", garageDoor.GetStatus())

	fmt.Println("\n9. Testing Error Handling")
	fmt.Println(strings.Repeat("-", 40))
	
	// Test invalid slot
	remote.OnButtonPressed(10) // Invalid slot
	remote.OffButtonPressed(-1) // Invalid slot

	fmt.Println("\n=== Demo Complete ===")
	fmt.Println("The Command pattern provides:")
	fmt.Println("✅ Decoupling between invoker and receiver")
	fmt.Println("✅ Undo/Redo functionality")
	fmt.Println("✅ Macro commands (composite operations)")
	fmt.Println("✅ Command queuing and logging")
	fmt.Println("✅ Flexible and extensible design")
}