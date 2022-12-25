package uniquenamegenerator

type Transformer func(s string) string

type option func(ung *UniqueNameGenerator)

// WithSeparator specifies the separator to be used to separate each selected dictionary word in the resulting name
func WithSeparator(separator string) option {
	return func(ung *UniqueNameGenerator) {
		ung.separator = separator
	}
}

// WithDictionaries sets the order specific dictionaries to be used in name generation
func WithDictionaries(dictionaries [][]string) option {
	return func(ung *UniqueNameGenerator) {
		ung.dictionaries = dictionaries
	}
}

// WithTransformer specifies a function to be applied against dictionary words selected for the unique name
// this is useful if the dictionaries are retrieved from uncontrolled sources
func WithTransformer(transformer Transformer) option {
	return func(ung *UniqueNameGenerator) {
		ung.transformer = transformer
	}
}
