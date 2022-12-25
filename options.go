package uniquenamegenerator

type Sanitizer func(str string) string

type Style int

const (
	Lower Style = iota
	Upper
	Capital
)

type option func(ung *UniqueNameGenerator)

func WithSeparator(separator string) option {
	return func(ung *UniqueNameGenerator) {
		ung.separator = separator
	}
}

func WithDictionaries(dictionaries [][]string) option {
	return func(ung *UniqueNameGenerator) {
		ung.dictionaries = dictionaries
	}
}

func WithStyle(style Style) option {
	return func(ung *UniqueNameGenerator) {
		ung.style = style
	}
}

func WithSanitizer(sanitizer Sanitizer) option {
	return func(ung *UniqueNameGenerator) {
		ung.sanitizer = sanitizer
	}
}
