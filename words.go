package randomizer

// word options
const (
	adjective int = iota
	noun
)

// Noun returns a random noun
func Noun() string {
	return randomize(jsonContent.Nouns...)
}

// Adjective returns a random adjective
func Adjective() string {
	return randomize(jsonContent.Adjectives...)
}

// Word returns a random word, either an adjective or a noun
func Word() string {
	const options = 2 // # of enum options

	s := getRandSource()

	randLock.Lock()
	n := s.Intn(options)
	randLock.Unlock()

	switch n {
	case adjective:
		return Adjective()
	default:
		return Noun()
	}
}

// Words returns a slice of words of count n
func Words(n int) []string {
	var words = make([]string, n)

	i := 0
	for i < n {
		words[i] = Word()
		i++
	}
	return words
}

func randomize(l ...string) string {
	if len(l) == 0 {
		return ""
	}

	if len(l) == 1 {
		return l[0]
	}

	s := getRandSource()

	randLock.Lock()
	n := s.Intn(len(l) - 1)
	randLock.Unlock()

	return l[n]
}
