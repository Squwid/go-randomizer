package randomizer

import (
	"math/rand"
	"sync"
	"time"
)

var rwRand *rand.Rand
var randSourceLock = &sync.RWMutex{}

var randLock = &sync.Mutex{}

func init() {
	NewRandSource()
}

// NewRandSourceFromSource uses a user provided source to set the random source
func NewRandSourceFromSource(source int64) {
	setRandSource(source)
}

// NewRandSource creates a new random source using the current time as the seed
func NewRandSource() {
	setRandSource(time.Now().UnixNano())
}

func setRandSource(source int64) {
	randSourceLock.Lock()
	defer randSourceLock.Unlock()

	rwRand = rand.New(rand.NewSource(source))
}

func getRandSource() *rand.Rand {
	randSourceLock.RLock()
	defer randSourceLock.RUnlock()

	return rwRand
}
