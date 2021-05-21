package randomizer

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRandomSource(t *testing.T) {
	NewRandSourceFromSource(0)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s %s\n", Noun(), Adjective())
	}
}

func TestRandomWords(t *testing.T) {
	// NewRandSourceFromSource(0)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", Word())
	}
}

func TestRandomDate(t *testing.T) {
	begin := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		fmt.Println(Date(begin, end).String())
	}
}

func TestRandomTitle(t *testing.T) {
	title := strings.Title(strings.Join(Words(Number(5, 10)), " "))
	fmt.Println(title)
}
