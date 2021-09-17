package randomizer

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const routines = 100
const subRoutines = 100

func TestSetRandConcurrency(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			r := getRandSource()
			_ = r

			NewRandSourceFromSource(int64(i))
		}(i)
	}
	wg.Wait()
}

func TestNumberConcurrency(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			defer wg.Done()

			min := 0
			max := 10

			n := Number(min, max)
			assert.LessOrEqual(t, n, max)
			assert.GreaterOrEqual(t, n, min)
		}()
	}
	wg.Wait()
}

func TestWordConcurrency(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < subRoutines; j++ {
				word := Word()
				_ = word
			}
		}()
	}
	wg.Wait()
}

func TestWordsConcurrency(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < subRoutines; j++ {
				words := Words(j)
				_ = words
			}
		}()
	}
	wg.Wait()
}

func TestRandomDate(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			defer wg.Done()

			begin := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC)
			end := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC)

			for j := 0; j < subRoutines; j++ {
				date := Date(begin, end).String()
				_ = date
			}
		}()
	}
	wg.Wait()
}
