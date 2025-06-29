// Package modern defines modern interfaces for various systems
package modern

// Database defines the modern database interface
// This represents our target interface for database operations
type Database interface {
	Connect() error
	Disconnect() error
	ExecuteQuery(query string) error
	ExecuteTransaction(queries []string) error
	GetConnectionInfo() string
}

// MediaPlayer defines the modern media player interface
// This represents our target interface for media playback
type MediaPlayer interface {
	Play(filename string) error
	Pause() error
	Stop() error
	SetVolume(volume int)
	GetVolume() int
	GetCurrentTrack() string
	GetDuration() int // in seconds
}

// CloudStorage defines the modern cloud storage interface
type CloudStorage interface {
	Upload(filename string, data []byte) error
	Download(filename string) ([]byte, error)
	Delete(filename string) error
	List(prefix string) ([]string, error)
	GetMetadata(filename string) (map[string]string, error)
}