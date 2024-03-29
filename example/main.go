package main

import (
	"fmt"
	"strings"

	ung "github.com/dillonstreator/go-unique-name-generator"
	"github.com/dillonstreator/go-unique-name-generator/dictionaries"
)

func main() {
	defaultGenerator := ung.NewUniqueNameGenerator()
	generator1 := ung.NewUniqueNameGenerator(
		ung.WithDictionaries([][]string{}),
		ung.WithSeparator("."),
	)
	generator2 := ung.NewUniqueNameGenerator(
		ung.WithDictionaries(
			[][]string{
				dictionaries.Colors,
				dictionaries.Animals,
				dictionaries.Names,
			},
		),
		ung.WithSeparator("-"),
	)
	generator3 := ung.NewUniqueNameGenerator(
		ung.WithDictionaries([][]string{
			dictionaries.Colors,
			dictionaries.Adjectives,
			dictionaries.Drinks,
		}),
		ung.WithTransformer(func(str string) string {
			return strings.Replace(str, " ", "", -1)
		}),
	)

	fmt.Printf("defaultGenerator possible unique names: %d\n", defaultGenerator.UniquenessCount())
	fmt.Printf("defaultGenerator name: %s\n", defaultGenerator.Generate())
	fmt.Printf("generator1 possible unique names: %d\n", generator1.UniquenessCount())
	fmt.Printf("generator1 name: %s\n", generator1.Generate())
	fmt.Printf("generator2 possible unique names: %d\n", generator2.UniquenessCount())
	fmt.Printf("generator2 name: %s\n", generator2.Generate())

	fmt.Printf("generator3 name: %s\n", generator3.Generate())
}
