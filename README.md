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
}
```

