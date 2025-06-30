package receivers

import "fmt"

// Light represents a smart light device
type Light struct {
	Location string
	IsOn     bool
	Brightness int // 0-100
}

func NewLight(location string) *Light {
	return &Light{
		Location:   location,
		IsOn:       false,
		Brightness: 0,
	}
}

func (l *Light) TurnOn() {
	l.IsOn = true
	l.Brightness = 100
	fmt.Printf("💡 %s light is ON (Brightness: %d%%)\n", l.Location, l.Brightness)
}

func (l *Light) TurnOff() {
	l.IsOn = false
	l.Brightness = 0
	fmt.Printf("💡 %s light is OFF\n", l.Location)
}

func (l *Light) SetBrightness(level int) {
	if level < 0 {
		level = 0
	} else if level > 100 {
		level = 100
	}
	
	l.Brightness = level
	if level > 0 {
		l.IsOn = true
	} else {
		l.IsOn = false
	}
	
	fmt.Printf("💡 %s light brightness set to %d%%\n", l.Location, l.Brightness)
}

func (l *Light) GetStatus() string {
	if l.IsOn {
		return fmt.Sprintf("%s light: ON (Brightness: %d%%)", l.Location, l.Brightness)
	}
	return fmt.Sprintf("%s light: OFF", l.Location)
}

// Fan represents a ceiling fan device
type Fan struct {
	Location string
	IsOn     bool
	Speed    int // 0-5
}

func NewFan(location string) *Fan {
	return &Fan{
		Location: location,
		IsOn:     false,
		Speed:    0,
	}
}

func (f *Fan) TurnOn() {
	f.IsOn = true
	f.Speed = 1
	fmt.Printf("🌀 %s fan is ON (Speed: %d)\n", f.Location, f.Speed)
}

func (f *Fan) TurnOff() {
	f.IsOn = false
	f.Speed = 0
	fmt.Printf("🌀 %s fan is OFF\n", f.Location)
}

func (f *Fan) SetSpeed(speed int) {
	if speed < 0 {
		speed = 0
	} else if speed > 5 {
		speed = 5
	}
	
	f.Speed = speed
	if speed > 0 {
		f.IsOn = true
	} else {
		f.IsOn = false
	}
	
	fmt.Printf("🌀 %s fan speed set to %d\n", f.Location, f.Speed)
}

func (f *Fan) GetStatus() string {
	if f.IsOn {
		return fmt.Sprintf("%s fan: ON (Speed: %d)", f.Location, f.Speed)
	}
	return fmt.Sprintf("%s fan: OFF", f.Location)
}

// Stereo represents a stereo system
type Stereo struct {
	Location string
	IsOn     bool
	Volume   int    // 0-100
	Station  string
}

func NewStereo(location string) *Stereo {
	return &Stereo{
		Location: location,
		IsOn:     false,
		Volume:   0,
		Station:  "",
	}
}

func (s *Stereo) TurnOn() {
	s.IsOn = true
	s.Volume = 50
	s.Station = "FM 101.5"
	fmt.Printf("🎵 %s stereo is ON (Volume: %d, Station: %s)\n", s.Location, s.Volume, s.Station)
}

func (s *Stereo) TurnOff() {
	s.IsOn = false
	s.Volume = 0
	s.Station = ""
	fmt.Printf("🎵 %s stereo is OFF\n", s.Location)
}

func (s *Stereo) SetVolume(volume int) {
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}
	
	s.Volume = volume
	fmt.Printf("🎵 %s stereo volume set to %d\n", s.Location, s.Volume)
}

func (s *Stereo) SetStation(station string) {
	s.Station = station
	fmt.Printf("🎵 %s stereo tuned to %s\n", s.Location, s.Station)
}

func (s *Stereo) GetStatus() string {
	if s.IsOn {
		return fmt.Sprintf("%s stereo: ON (Volume: %d, Station: %s)", s.Location, s.Volume, s.Station)
	}
	return fmt.Sprintf("%s stereo: OFF", s.Location)
}

// GarageDoor represents a garage door
type GarageDoor struct {
	IsOpen bool
}

func NewGarageDoor() *GarageDoor {
	return &GarageDoor{
		IsOpen: false,
	}
}

func (g *GarageDoor) Open() {
	if !g.IsOpen {
		g.IsOpen = true
		fmt.Println("🚪 Garage door is OPENING...")
		fmt.Println("🚪 Garage door is OPEN")
	} else {
		fmt.Println("🚪 Garage door is already OPEN")
	}
}

func (g *GarageDoor) Close() {
	if g.IsOpen {
		g.IsOpen = false
		fmt.Println("🚪 Garage door is CLOSING...")
		fmt.Println("🚪 Garage door is CLOSED")
	} else {
		fmt.Println("🚪 Garage door is already CLOSED")
	}
}

func (g *GarageDoor) GetStatus() string {
	if g.IsOpen {
		return "Garage door: OPEN"
	}
	return "Garage door: CLOSED"
}