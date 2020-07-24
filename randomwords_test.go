package randomwords

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRandomWords(t *testing.T) {
	fmt.Println("*** Printing Nouns ***")
	for i := 0; i < 1000; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(strings.Title(fmt.Sprintf("%s %s\n", RandomAdjective(), RandomNoun())))
	}
}
