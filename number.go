package randomizer

// Number returns a random number between and inclusive of min and max
func Number(min, max int) int {
	r := getRandSource()
	randLock.Lock()
	n := r.Intn(max-min) + min
	randLock.Unlock()
	return n
}
