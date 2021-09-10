package uniquenamegenerator

import (
	"strings"
	"testing"

	"github.com/DillonStreator/go-unique-name-generator/dictionaries"
)

func includes(items []string, str string, transform func(item string) string) bool {
	for _, item := range items {
		if transform(item) == str {
			return true
		}
	}
	return false
}

func TestNewUniqueNameGenerator(t *testing.T) {
	t.Run("should use correct default options", func(t *testing.T) {
		g := NewUniqueNameGenerator(UNGOpts{})
		if len(g.Dictionaries) != 3 {
			t.Error("default generator should use 3 dictionaries")
		}
		word := g.Generate()
		if strings.Count(word, "_") != 2 {
			t.Error("should output 3 words separated by 2 underscores")
		}
		words := strings.Split(word, "_")
		if !includes(dictionaries.Adjectives, words[0], strings.ToLower) {
			t.Error("first word should be adjective")
		}
		if !includes(dictionaries.Colors, words[1], strings.ToLower) {
			t.Error("second word should be color")
		}
		if !includes(dictionaries.Names, words[2], strings.ToLower) {
			t.Error("third word should be name")
		}
	})

	t.Run("should respect provided config", func(t *testing.T) {
		g := NewUniqueNameGenerator(UNGOpts{
			Dictionaries: [][]string{
				dictionaries.Colors,
				dictionaries.Colors,
				dictionaries.Adjectives,
				dictionaries.Names,
			},
			Separator: "-",
			Style:     Upper,
		})
		if len(g.Dictionaries) != 4 {
			t.Error("should have 4 dictionaries")
		}
		word := g.Generate()
		if strings.Count(word, "-") != 3 {
			t.Error("should have 4 words separated by 3 dashes")
		}
		words := strings.Split(word, "-")
		if !includes(dictionaries.Colors, words[0], strings.ToUpper) {
			t.Error("first word should be color")
		}
		if !includes(dictionaries.Colors, words[1], strings.ToUpper) {
			t.Error("second word should be color")
		}
		if !includes(dictionaries.Adjectives, words[2], strings.ToUpper) {
			t.Error("third word should be adjective")
		}
		if !includes(dictionaries.Names, words[3], strings.ToUpper) {
			t.Error("fourth word should be name")
		}

		g2 := NewUniqueNameGenerator(UNGOpts{
			Dictionaries: [][]string{
				{"dillon"},
				{"streator"},
			},
			Style: Capital,
		})
		word = g2.Generate()
		if word != "Dillon_Streator" {
			t.Error("it should use Capital casing")
		}
	})

	t.Run("UniquenessCount calculations", func(t *testing.T) {
		dicts := [][]string{
			{"1", "2", "3"},
			{"1", "2", "3"},
		}
		g := NewUniqueNameGenerator(UNGOpts{
			Dictionaries: dicts,
		})
		actual := g.UniquenessCount()
		expected := uint64(9)
		if actual != expected {
			t.Errorf("expected %d combinations with %v but got %d", expected, dicts, actual)
		}
	})

	t.Run("Sanitizes dictionaries", func(t *testing.T) {
		g := NewUniqueNameGenerator(UNGOpts{
			Dictionaries: [][]string{
				{"St. John"},
				{"t t"},
			},
		})
		actual := g.Generate()
		expected := "stjohn_tt"
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}

		g = NewUniqueNameGenerator(UNGOpts{
			Dictionaries: [][]string{
				{"St. John"},
				{"t t"},
			},
			Sanitizer: func(str string) string {
				return strings.Replace(str, " ", "", -1)
			},
		})
		actual = g.Generate()
		expected = "st.john_tt"
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})
}
