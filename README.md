# go-randomizer

[![GoDoc][1]][2] [![Release][3]][4] [![Go Report Card][5]][6] [![Coverage][7]][8]

[1]: https://godoc.org/github.com/Squwid/go-randomizer?status.svg
[2]: https://godoc.org/github.com/Squwid/go-randomizer
[3]: https://img.shields.io/github/v/release/Squwid/go-randomizer.svg
[4]: https://github.com/Squwid/go-randomizer/releases/latest
[5]: https://goreportcard.com/badge/github.com/Squwid/go-randomizer
[6]: https://goreportcard.com/report/github.com/Squwid/go-randomizer
[7]: coverage.png
[8]: https://github.com/Squwid/go-randomizer



Randomizer is a library that randomizes different types of words, such as Adjectives and Nouns, or just Words and handles all of Go's randomizing concurrency for you. Also can pick a time between a start and end date.

## Installation

`go get github.com/Squwid/go-randomizer`

## How to Use

The math/rand seed is set as `time.Now().UnixNano()` when the package is initialized. A custom seed can also be used by calling `randomizer.NewRandSourceFromSource(source int64)`.

```go
package main

import (
    "fmt"

    "github.com/Squwid/go-randomizer"
)

func main() {
    // Print a snake case Adjective - Noun 
    // ex. dizzy_horse
    snakeCase := fmt.Sprintf("%s_%s", randomizer.Adjective(), randomizer.Noun())
    fmt.Println(snakeCase)

    // Print a title case Adjective - Noun
    // ex. Pretty Eagle
    titleCase := strings.Title(fmt.Sprintf("%s %s", randomizer.Adjective(), randomizer.Noun()))
    fmt.Println(titleCase)

    // Print a fake title between 5 and 10 words
    // ex. Astonishing Tiger Massive Glorious Dress Walrus Swanky
    title := strings.Title(strings.Join(randomizer.Words(randomizer.Number(5, 10)), " "))
    fmt.Println(title)

    // Print a random time in 2020
    // ex. 2020-06-22 13:34:06 -0400 EDT
    begin := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC)
    date := randomizer.Date(begin, end)
    fmt.Println(date.String())
}
```

## Set Custom Rand Seed

Setting a custom seed is effective for testing, since having the same seed will guarantee that randomize functions will have the same randomized order.

```go
package main

import (
    "github.com/Squwid/go-randomizer"
)

func main() {
    randomizer.NewRandSourceFromSource(0) // Extremely effective for testing, randomize will always have same order
}
```