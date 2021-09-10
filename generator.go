package uniquenamegenerator

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/DillonStreator/go-unique-name-generator/dictionaries"
)

type Style int

const (
	Lower Style = iota
	Upper
	Capital
)

var alphaNumericReg = regexp.MustCompile("[^a-zA-Z0-9]+")

func sanitizeString(str string) string {
	return alphaNumericReg.ReplaceAllString(str, "")
}

// UniqueNameGenerator is a unique name generator instance
type UniqueNameGenerator struct {
	Separator    string
	Dictionaries [][]string
	Style        Style
	Sanitizer    Sanitizer
}

type Sanitizer func(str string) string

// UNGOpts are the options for creating a new UniqueNameGenerator
type UNGOpts struct {
	Separator    string
	Dictionaries [][]string
	Style        Style
	Sanitizer    Sanitizer
}

// NewUniqueNameGenerator creates a new instance of UniqueNameGenerator
func NewUniqueNameGenerator(opts UNGOpts) *UniqueNameGenerator {
	rand.Seed(time.Now().UTC().UnixNano())
	separator := opts.Separator
	if separator == "" {
		separator = "_"
	}
	dicts := opts.Dictionaries
	if len(dicts) == 0 {
		dicts = [][]string{
			dictionaries.Adjectives,
			dictionaries.Colors,
			dictionaries.Names,
		}
	}
	sanitizer := opts.Sanitizer
	if sanitizer == nil {
		sanitizer = sanitizeString
	}
	return &UniqueNameGenerator{
		Separator:    separator,
		Dictionaries: dicts,
		Style:        opts.Style,
		Sanitizer:    sanitizer,
	}
}

// Generate generates a new unique name with the configuration
func (ung *UniqueNameGenerator) Generate() string {
	words := make([]string, len(ung.Dictionaries))
	for i, dict := range ung.Dictionaries {
		randomIndex := rand.Intn(len(dict))
		word := ung.Sanitizer(dict[randomIndex])
		switch ung.Style {
		case Lower:
			word = strings.ToLower(word)
		case Upper:
			word = strings.ToUpper(word)
		case Capital:
			word = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
		words[i] = word
	}
	return strings.Join(words, ung.Separator)
}

// UniquenessCount returns the number of unique combinations
func (ung *UniqueNameGenerator) UniquenessCount() uint64 {
	var count uint64 = 1
	for _, set := range ung.Dictionaries {
		count *= uint64(len(set))
	}
	return count
}
