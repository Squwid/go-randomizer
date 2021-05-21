package randomwords

// Number returns a random number between and inclusive of min and max
func Number(min, max int) int {
	return getRandSource().Intn(max-min) + min
}
