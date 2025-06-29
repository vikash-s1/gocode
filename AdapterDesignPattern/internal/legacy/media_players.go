// Package legacy contains old media player systems with incompatible interfaces
package legacy

import (
	"fmt"
	"path/filepath"
	"strings"
)

// MP3Player represents a legacy MP3 player with its own interface
type MP3Player struct {
	currentFile string
	isPlaying   bool
	volume      int
}

// PlayMP3 is the MP3 player's specific method for playing files
func (m *MP3Player) PlayMP3(filename string) error {
	if !strings.HasSuffix(strings.ToLower(filename), ".mp3") {
		return fmt.Errorf("MP3 player can only play .mp3 files, got: %s", filename)
	}
	
	fmt.Printf("   ♪ MP3Player: Playing %s\n", filename)
	m.currentFile = filename
	m.isPlaying = true
	return nil
}

// StopMP3 stops the MP3 playback
func (m *MP3Player) StopMP3() error {
	if !m.isPlaying {
		return fmt.Errorf("MP3 player is not currently playing")
	}
	
	fmt.Printf("   ⏹ MP3Player: Stopped playing %s\n", m.currentFile)
	m.isPlaying = false
	m.currentFile = ""
	return nil
}

// SetMP3Volume sets the volume for MP3 player
func (m *MP3Player) SetMP3Volume(vol int) {
	if vol < 0 {
		vol = 0
	}
	if vol > 100 {
		vol = 100
	}
	m.volume = vol
	fmt.Printf("   🔊 MP3Player: Volume set to %d%%\n", vol)
}

// GetMP3Volume returns the current volume
func (m *MP3Player) GetMP3Volume() int {
	return m.volume
}

// GetMP3Status returns MP3 player status
func (m *MP3Player) GetMP3Status() string {
	if m.isPlaying {
		return fmt.Sprintf("Playing: %s", filepath.Base(m.currentFile))
	}
	return "Stopped"
}

// WAVPlayer represents a legacy WAV player with its own interface
type WAVPlayer struct {
	activeTrack string
	playing     bool
	audioLevel  int
}

// StartWAV is the WAV player's specific method for playing files
func (w *WAVPlayer) StartWAV(filepath string) error {
	if !strings.HasSuffix(strings.ToLower(filepath), ".wav") {
		return fmt.Errorf("WAV player can only play .wav files, got: %s", filepath)
	}
	
	fmt.Printf("   ♪ WAVPlayer: Starting playback of %s\n", filepath)
	w.activeTrack = filepath
	w.playing = true
	return nil
}

// HaltWAV stops the WAV playback
func (w *WAVPlayer) HaltWAV() error {
	if !w.playing {
		return fmt.Errorf("WAV player is not currently playing")
	}
	
	fmt.Printf("   ⏹ WAVPlayer: Halted playback of %s\n", w.activeTrack)
	w.playing = false
	w.activeTrack = ""
	return nil
}

// AdjustWAVLevel adjusts the audio level for WAV player
func (w *WAVPlayer) AdjustWAVLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > 100 {
		level = 100
	}
	w.audioLevel = level
	fmt.Printf("   🔊 WAVPlayer: Audio level adjusted to %d%%\n", level)
}

// GetWAVLevel returns the current audio level
func (w *WAVPlayer) GetWAVLevel() int {
	return w.audioLevel
}

// GetWAVInfo returns WAV player information
func (w *WAVPlayer) GetWAVInfo() string {
	if w.playing {
		return fmt.Sprintf("Active: %s", filepath.Base(w.activeTrack))
	}
	return "Inactive"
}

// FLACPlayer represents a legacy FLAC player with its own interface
type FLACPlayer struct {
	loadedFile   string
	playbackMode bool
	soundLevel   int
}

// LoadAndPlayFLAC is the FLAC player's specific method for playing files
func (f *FLACPlayer) LoadAndPlayFLAC(file string) error {
	if !strings.HasSuffix(strings.ToLower(file), ".flac") {
		return fmt.Errorf("FLAC player can only play .flac files, got: %s", file)
	}
	
	fmt.Printf("   ♪ FLACPlayer: Loading and playing %s\n", file)
	f.loadedFile = file
	f.playbackMode = true
	return nil
}

// TerminateFLAC stops the FLAC playback
func (f *FLACPlayer) TerminateFLAC() error {
	if !f.playbackMode {
		return fmt.Errorf("FLAC player is not in playback mode")
	}
	
	fmt.Printf("   ⏹ FLACPlayer: Terminated playback of %s\n", f.loadedFile)
	f.playbackMode = false
	f.loadedFile = ""
	return nil
}

// ModifySoundLevel modifies the sound level for FLAC player
func (f *FLACPlayer) ModifySoundLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > 100 {
		level = 100
	}
	f.soundLevel = level
	fmt.Printf("   🔊 FLACPlayer: Sound level modified to %d%%\n", level)
}

// GetSoundLevel returns the current sound level
func (f *FLACPlayer) GetSoundLevel() int {
	return f.soundLevel
}

// GetFLACDetails returns FLAC player details
func (f *FLACPlayer) GetFLACDetails() string {
	if f.playbackMode {
		return fmt.Sprintf("Loaded: %s", filepath.Base(f.loadedFile))
	}
	return "No file loaded"
}