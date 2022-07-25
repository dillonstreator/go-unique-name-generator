package uniquenamegenerator

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/DillonStreator/go-unique-name-generator/dictionaries"
)

var alphaNumericReg = regexp.MustCompile("[^a-zA-Z0-9]+")

func sanitizeString(str string) string {
	return alphaNumericReg.ReplaceAllString(str, "")
}

// UniqueNameGenerator is a unique name generator instance
type UniqueNameGenerator struct {
	options *options
}

// NewUniqueNameGenerator creates a new instance of UniqueNameGenerator
func NewUniqueNameGenerator(opts ...option) *UniqueNameGenerator {
	rand.Seed(time.Now().UTC().UnixNano())

	_opts := &options{
		separator: "_",
		dictionaries: [][]string{
			dictionaries.Adjectives,
			dictionaries.Colors,
			dictionaries.Names,
		},
		sanitizer: sanitizeString,
		style:     Lower,
	}

	for _, opt := range opts {
		opt(_opts)
	}

	return &UniqueNameGenerator{
		options: _opts,
	}
}

// Generate generates a new unique name with the configuration
func (ung *UniqueNameGenerator) Generate() string {
	words := make([]string, len(ung.options.dictionaries))
	for i, dict := range ung.options.dictionaries {
		randomIndex := rand.Intn(len(dict))
		word := ung.options.sanitizer(dict[randomIndex])
		switch ung.options.style {
		case Lower:
			word = strings.ToLower(word)
		case Upper:
			word = strings.ToUpper(word)
		case Capital:
			word = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
		words[i] = word
	}
	return strings.Join(words, ung.options.separator)
}

// UniquenessCount returns the number of unique combinations
func (ung *UniqueNameGenerator) UniquenessCount() uint64 {
	if len(ung.options.dictionaries) == 0 {
		return 0
	}
	var count uint64 = 1
	for _, set := range ung.options.dictionaries {
		count *= uint64(len(set))
	}
	return count
}
