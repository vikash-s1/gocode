// Package adapters contains media player adapter implementations
package adapters

import (
	"fmt"

	"github.com/example/adapter-pattern/internal/legacy"
	"github.com/example/adapter-pattern/internal/modern"
)

// MP3Adapter adapts the legacy MP3 player to the modern media player interface
type MP3Adapter struct {
	player *legacy.MP3Player
}

// NewMP3Adapter creates a new MP3 adapter
func NewMP3Adapter(player *legacy.MP3Player) *MP3Adapter {
	return &MP3Adapter{player: player}
}

// Play adapts MP3Player's PlayMP3 method to our standard interface
func (m *MP3Adapter) Play(filename string) error {
	return m.player.PlayMP3(filename)
}

// Pause simulates pause functionality (MP3 player doesn't have native pause)
func (m *MP3Adapter) Pause() error {
	// Since legacy MP3Player doesn't have pause, we simulate it
	fmt.Printf("   ⏸ MP3Adapter: Simulating pause for %s\n", m.player.GetMP3Status())
	return nil
}

// Stop adapts MP3Player's StopMP3 method to our standard interface
func (m *MP3Adapter) Stop() error {
	return m.player.StopMP3()
}

// SetVolume adapts MP3Player's SetMP3Volume method to our standard interface
func (m *MP3Adapter) SetVolume(volume int) {
	m.player.SetMP3Volume(volume)
}

// GetVolume adapts MP3Player's GetMP3Volume method to our standard interface
func (m *MP3Adapter) GetVolume() int {
	return m.player.GetMP3Volume()
}

// GetCurrentTrack adapts MP3Player's status method to our standard interface
func (m *MP3Adapter) GetCurrentTrack() string {
	return m.player.GetMP3Status()
}

// GetDuration returns simulated duration (MP3 player doesn't provide this)
func (m *MP3Adapter) GetDuration() int {
	// Simulate duration since legacy player doesn't provide it
	return 180 // 3 minutes
}

// WAVAdapter adapts the legacy WAV player to the modern media player interface
type WAVAdapter struct {
	player *legacy.WAVPlayer
}

// NewWAVAdapter creates a new WAV adapter
func NewWAVAdapter(player *legacy.WAVPlayer) *WAVAdapter {
	return &WAVAdapter{player: player}
}

// Play adapts WAVPlayer's StartWAV method to our standard interface
func (w *WAVAdapter) Play(filename string) error {
	return w.player.StartWAV(filename)
}

// Pause simulates pause functionality (WAV player doesn't have native pause)
func (w *WAVAdapter) Pause() error {
	// Since legacy WAVPlayer doesn't have pause, we simulate it
	fmt.Printf("   ⏸ WAVAdapter: Simulating pause for %s\n", w.player.GetWAVInfo())
	return nil
}

// Stop adapts WAVPlayer's HaltWAV method to our standard interface
func (w *WAVAdapter) Stop() error {
	return w.player.HaltWAV()
}

// SetVolume adapts WAVPlayer's AdjustWAVLevel method to our standard interface
func (w *WAVAdapter) SetVolume(volume int) {
	w.player.AdjustWAVLevel(volume)
}

// GetVolume adapts WAVPlayer's GetWAVLevel method to our standard interface
func (w *WAVAdapter) GetVolume() int {
	return w.player.GetWAVLevel()
}

// GetCurrentTrack adapts WAVPlayer's info method to our standard interface
func (w *WAVAdapter) GetCurrentTrack() string {
	return w.player.GetWAVInfo()
}

// GetDuration returns simulated duration (WAV player doesn't provide this)
func (w *WAVAdapter) GetDuration() int {
	// Simulate duration since legacy player doesn't provide it
	return 240 // 4 minutes
}

// FLACAdapter adapts the legacy FLAC player to the modern media player interface
type FLACAdapter struct {
	player *legacy.FLACPlayer
}

// NewFLACAdapter creates a new FLAC adapter
func NewFLACAdapter(player *legacy.FLACPlayer) *FLACAdapter {
	return &FLACAdapter{player: player}
}

// Play adapts FLACPlayer's LoadAndPlayFLAC method to our standard interface
func (f *FLACAdapter) Play(filename string) error {
	return f.player.LoadAndPlayFLAC(filename)
}

// Pause simulates pause functionality (FLAC player doesn't have native pause)
func (f *FLACAdapter) Pause() error {
	// Since legacy FLACPlayer doesn't have pause, we simulate it
	fmt.Printf("   ⏸ FLACAdapter: Simulating pause for %s\n", f.player.GetFLACDetails())
	return nil
}

// Stop adapts FLACPlayer's TerminateFLAC method to our standard interface
func (f *FLACAdapter) Stop() error {
	return f.player.TerminateFLAC()
}

// SetVolume adapts FLACPlayer's ModifySoundLevel method to our standard interface
func (f *FLACAdapter) SetVolume(volume int) {
	f.player.ModifySoundLevel(volume)
}

// GetVolume adapts FLACPlayer's GetSoundLevel method to our standard interface
func (f *FLACAdapter) GetVolume() int {
	return f.player.GetSoundLevel()
}

// GetCurrentTrack adapts FLACPlayer's details method to our standard interface
func (f *FLACAdapter) GetCurrentTrack() string {
	return f.player.GetFLACDetails()
}

// GetDuration returns simulated duration (FLAC player doesn't provide this)
func (f *FLACAdapter) GetDuration() int {
	// Simulate duration since legacy player doesn't provide it
	return 300 // 5 minutes
}

// MediaAdapterFactory creates media player adapters based on file type
type MediaAdapterFactory struct{}

// CreateMediaAdapter creates the appropriate media adapter based on file extension
func (f *MediaAdapterFactory) CreateMediaAdapter(filename string) (modern.MediaPlayer, error) {
	switch {
	case len(filename) >= 4 && filename[len(filename)-4:] == ".mp3":
		return NewMP3Adapter(&legacy.MP3Player{}), nil
	case len(filename) >= 4 && filename[len(filename)-4:] == ".wav":
		return NewWAVAdapter(&legacy.WAVPlayer{}), nil
	case len(filename) >= 5 && filename[len(filename)-5:] == ".flac":
		return NewFLACAdapter(&legacy.FLACPlayer{}), nil
	default:
		return nil, fmt.Errorf("unsupported media format for file: %s", filename)
	}
}

// UniversalMediaAdapter can handle multiple formats through composition
type UniversalMediaAdapter struct {
	mp3Player  *MP3Adapter
	wavPlayer  *WAVAdapter
	flacPlayer *FLACAdapter
	currentPlayer modern.MediaPlayer
}

// NewUniversalMediaAdapter creates a universal media adapter
func NewUniversalMediaAdapter() *UniversalMediaAdapter {
	return &UniversalMediaAdapter{
		mp3Player:  NewMP3Adapter(&legacy.MP3Player{}),
		wavPlayer:  NewWAVAdapter(&legacy.WAVPlayer{}),
		flacPlayer: NewFLACAdapter(&legacy.FLACPlayer{}),
	}
}

// Play automatically selects the appropriate player based on file extension
func (u *UniversalMediaAdapter) Play(filename string) error {
	factory := &MediaAdapterFactory{}
	player, err := factory.CreateMediaAdapter(filename)
	if err != nil {
		return err
	}
	
	u.currentPlayer = player
	return player.Play(filename)
}

// Pause delegates to the current player
func (u *UniversalMediaAdapter) Pause() error {
	if u.currentPlayer == nil {
		return fmt.Errorf("no media currently playing")
	}
	return u.currentPlayer.Pause()
}

// Stop delegates to the current player
func (u *UniversalMediaAdapter) Stop() error {
	if u.currentPlayer == nil {
		return fmt.Errorf("no media currently playing")
	}
	return u.currentPlayer.Stop()
}

// SetVolume delegates to the current player
func (u *UniversalMediaAdapter) SetVolume(volume int) {
	if u.currentPlayer != nil {
		u.currentPlayer.SetVolume(volume)
	}
}

// GetVolume delegates to the current player
func (u *UniversalMediaAdapter) GetVolume() int {
	if u.currentPlayer != nil {
		return u.currentPlayer.GetVolume()
	}
	return 0
}

// GetCurrentTrack delegates to the current player
func (u *UniversalMediaAdapter) GetCurrentTrack() string {
	if u.currentPlayer != nil {
		return u.currentPlayer.GetCurrentTrack()
	}
	return "No media loaded"
}

// GetDuration delegates to the current player
func (u *UniversalMediaAdapter) GetDuration() int {
	if u.currentPlayer != nil {
		return u.currentPlayer.GetDuration()
	}
	return 0
}