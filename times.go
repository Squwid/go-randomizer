package randomizer

import (
	"math/rand"
	"time"
)

func Date(begin, end time.Time) time.Time {
	delta := end.Unix() - begin.Unix()
	return time.Unix(rand.New(getRandSource()).Int63n(delta)+begin.Unix(), 0)
}
