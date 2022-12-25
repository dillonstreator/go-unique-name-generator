package uniquenamegenerator

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/dillonstreator/go-unique-name-generator/dictionaries"
)

var alphaNumericReg = regexp.MustCompile("[^a-zA-Z0-9]+")

func sanitizeString(str string) string {
	return alphaNumericReg.ReplaceAllString(str, "")
}

// UniqueNameGenerator is a unique name generator instance
type UniqueNameGenerator struct {
	separator    string
	dictionaries [][]string
	style        Style
	sanitizer    Sanitizer

	rnd *rand.Rand
}

// NewUniqueNameGenerator creates a new instance of UniqueNameGenerator
func NewUniqueNameGenerator(opts ...option) *UniqueNameGenerator {
	ung := &UniqueNameGenerator{
		separator: "_",
		dictionaries: [][]string{
			dictionaries.Adjectives,
			dictionaries.Colors,
			dictionaries.Names,
		},
		sanitizer: sanitizeString,
		style:     Lower,
		rnd:       rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}

	for _, opt := range opts {
		opt(ung)
	}

	return ung
}

// Generate generates a new unique name with the configuration
func (ung *UniqueNameGenerator) Generate() string {
	words := make([]string, len(ung.dictionaries))
	for i, dict := range ung.dictionaries {
		randomIndex := ung.rnd.Intn(len(dict))
		word := ung.sanitizer(dict[randomIndex])
		switch ung.style {
		case Lower:
			word = strings.ToLower(word)
		case Upper:
			word = strings.ToUpper(word)
		case Capital:
			word = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
		words[i] = word
	}
	return strings.Join(words, ung.separator)
}

// UniquenessCount returns the number of unique combinations
func (ung *UniqueNameGenerator) UniquenessCount() uint64 {
	if len(ung.dictionaries) == 0 {
		return 0
	}
	var count uint64 = 1
	for _, set := range ung.dictionaries {
		count *= uint64(len(set))
	}
	return count
}
