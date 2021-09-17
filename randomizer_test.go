package randomizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord(t *testing.T) {
	NewRandSourceFromSource(0)

	var expected = []string{
		"hungry",
		"mouse",
		"badger",
		"hand",
		"immense",
		"succulent",
	}

	for i := 0; i < len(expected); i++ {
		word := Word()
		assert.Equal(t, expected[i], word)
	}
}

func TestNumber(t *testing.T) {
	NewRandSourceFromSource(0)

	var expected = []int{
		0,
		0,
		4,
		1,
	}

	for i := 0; i < len(expected); i++ {
		num := Number(0, 9)
		assert.Equal(t, expected[i], num)
	}
}
