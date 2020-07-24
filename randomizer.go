package randomwords

import (
	"math/rand"
	"sync"
	"time"
)

var randSource rand.Source
var randSourceLock = &sync.RWMutex{}

func init() {
	setRandSource(time.Now().Unix())
}

// NewRandSourceFromSource uses a user provided source to set the random source
func NewRandSourceFromSource(source int64) {
	setRandSource(source)
}

// NewRandSource creates a new random source using the current time as the seed
func NewRandSource() {
	setRandSource(time.Now().Unix())
}

func setRandSource(source int64) {
	randSourceLock.Lock()
	defer randSourceLock.Unlock()

	randSource = rand.NewSource(source)
}

func getRandSource() rand.Source {
	randSourceLock.RLock()
	defer randSourceLock.RUnlock()

	return randSource
}
