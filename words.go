package randomwords

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

	switch getRandSource().Intn(options) {
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

	return l[getRandSource().Intn(len(l)-1)]
}
