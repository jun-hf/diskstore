package diskstore

import (
	"os"
	"sync"
)

const (
	defaultBasePath = "diskstore"
	defaultFilePerm os.FileMode = 0o666
	defaultPathPerm os.FileMode = 0o777
)

// PathKey represents the path 
// and file name of a content being store
type PathKey struct {
	Path []string
	FileName string
	originalKey string
}

var (
	defaultAdvancedTransform = func(s string) *PathKey { return &PathKey{Path: []string{}, FileName: s}}
	
)

// Options define a set properties that determine Diskstore behavior
// All values are optional
type Options struct {
	BasePath string
	Transform TransformFunc
	AdvancedTransform AdvancedTransformFunc
	InverseTransform InverseTransformFunc
	CacheSizeMax uint64 
	PathPerm os.FileMode
	FilePerm os.FileMode

	// If TempDir is set, it will enable filesystem atomic writes 
	// by writing temporary files to that location before being moved
	// to BasePath.
	TempDir string

	Index Index
	IndexLess LessFunction

	Compression Compression
}

type Diskstore struct {
	Options
	mu sync.RWMutex
	cache map[string][]byte
	cacheSize uint64
}