# go-randomwords

Randomwords is a library that randomizes different types of words, such as

* Adjectives
* Nouns

## Installation

`go get github.com/Squwid/go-randomwords`

## How to Use

```go
package main

import (
    "fmt"

    "github.com/Squwid/go-randomwords"
)

func main() {
    // Print a snake case Adjective - Noun 
    // ex. dizzy_horse
    snakeCase := fmt.Sprintf("%s_%s", randomwords.Adjective(), randomwords.Noun())
    fmt.Println(snakeCase)

    // Print a title case Adjective - Noun
    // ex. Pretty Eagle
    titleCase := strings.Title(fmt.Sprintf("%s %s", randomwords.Adjective(), randomwords.Noun()))
    fmt.Println(titleCase)

    // Print a random time in 2020
    // ex. 2020-06-22 13:34:06 -0400 EDT
    begin := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC)
    date := randomwords.Date(begin, end)
    fmt.Println(date.String())
}
```

