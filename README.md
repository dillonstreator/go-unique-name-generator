# go-unique-name-generator <img src="https://img.icons8.com/external-kiranshastry-gradient-kiranshastry/64/000000/external-fingerprint-law-and-crime-kiranshastry-gradient-kiranshastry.png" height="25" width="25"/> <img src="https://img.icons8.com/color/48/000000/golang.png" height="25" width="25"/>

> unique name generator for go with 308,769,760 unique combinations by default

## Installation
```sh
go get github.com/DillonStreator/go-unique-name-generator
```

## Basic Usage

```go
package main

import (
    "fmt"

    ung "github.com/DillonStreator/go-unique-name-generator"
)

func main() {
    generator := ung.NewUniqueNameGenerator()

    fmt.Println(generator.Generate())
    // "{adjective}_{color}_{name}"

    fmt.Println(generator.UniquenessCount())
    // 308769760
}
```

## Usage continued ([examples](./example/main.go))
