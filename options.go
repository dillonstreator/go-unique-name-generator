package uniquenamegenerator

type Sanitizer func(str string) string

type Style int

const (
	Lower Style = iota
	Upper
	Capital
)

type options struct {
	separator    string
	dictionaries [][]string
	style        Style
	sanitizer    Sanitizer
}

type option func(opts *options)

func WithSeparator(separator string) option {
	return func(opts *options) {
		opts.separator = separator
	}
}

func WithDictionaries(dictionaries [][]string) option {
	return func(opts *options) {
		opts.dictionaries = dictionaries
	}
}

func WithStyle(style Style) option {
	return func(opts *options) {
		opts.style = style
	}
}

func WithSanitizer(sanitizer Sanitizer) option {
	return func(opts *options) {
		opts.sanitizer = sanitizer
	}
}
