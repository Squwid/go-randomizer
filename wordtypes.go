package randomwords

import "math/rand"

// RandomNoun returns a random noun
func RandomNoun() string {
	return random(jsonContent.Nouns)
}

// RandomAdjective returns a random adjective
func RandomAdjective() string {
	return random(jsonContent.Adjectives)
}

func random(list []string) string {
	if len(list) == 0 {
		return ""
	}

	if len(list) == 1 {
		return list[0]
	}

	i := rand.New(getRandSource()).Intn(len(list) - 1)
	return list[i]
}
