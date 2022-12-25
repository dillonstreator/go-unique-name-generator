package uniquenamegenerator

import (
	"regexp"
	"strings"
	"testing"

	"github.com/dillonstreator/go-unique-name-generator/dictionaries"
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
		g := NewUniqueNameGenerator(
			WithTransformer(strings.ToLower),
		)
		if len(g.dictionaries) != 3 {
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
		g := NewUniqueNameGenerator(
			WithDictionaries([][]string{
				dictionaries.Colors,
				dictionaries.Colors,
				dictionaries.Adjectives,
				dictionaries.Names,
			}),
			WithSeparator("-"),
			WithTransformer(strings.ToUpper),
		)
		if len(g.dictionaries) != 4 {
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

		g2 := NewUniqueNameGenerator(WithDictionaries([][]string{
			{"dillon"},
			{"streator"},
		}), WithTransformer(func(s string) string { return strings.ToUpper(s) }))
		word = g2.Generate()
		if word != "DILLON_STREATOR" {
			t.Error("it should transform to upper")
		}
	})

	t.Run("UniquenessCount calculations", func(t *testing.T) {
		dicts := [][]string{
			{"1", "2", "3"},
			{"1", "2", "3"},
		}
		g := NewUniqueNameGenerator(WithDictionaries(dicts))
		actual := g.UniquenessCount()
		expected := uint64(9)
		if actual != expected {
			t.Errorf("expected %d combinations with %v but got %d", expected, dicts, actual)
		}

		g = NewUniqueNameGenerator(WithDictionaries([][]string{}))
		actual = g.UniquenessCount()
		expected = uint64(0)
		if actual != expected {
			t.Errorf("expected %d combinations with %v but got %d", expected, dicts, actual)
		}

		g = NewUniqueNameGenerator(WithDictionaries([][]string{{}}))
		actual = g.UniquenessCount()
		expected = uint64(0)
		if actual != expected {
			t.Errorf("expected %d combinations with %v but got %d", expected, dicts, actual)
		}
	})

	t.Run("Sanitizes dictionaries", func(t *testing.T) {
		re := regexp.MustCompile("[. ]")
		g := NewUniqueNameGenerator(
			WithDictionaries([][]string{
				{"St. John"},
				{"t t"},
			}),
			WithTransformer(func(s string) string {
				return strings.ToLower(re.ReplaceAllString(s, ""))
			}),
		)
		actual := g.Generate()
		expected := "stjohn_tt"
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}

		g = NewUniqueNameGenerator(WithDictionaries([][]string{
			{"St. John"},
			{"t t"},
		}), WithTransformer(func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", ""))
		}))
		actual = g.Generate()
		expected = "st.john_tt"
		if actual != expected {
			t.Errorf("expected %s but got %s", expected, actual)
		}
	})
}

func BenchmarkUniqueNameGenerator_Generate(b *testing.B) {
	ung := NewUniqueNameGenerator()

	for i := 0; i < b.N; i++ {
		ung.Generate()
	}

}
