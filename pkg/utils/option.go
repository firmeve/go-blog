package utils

type Option interface{}

type OptionFunc func(option Option)

func ApplyOption(option Option, options ...OptionFunc) Option {
	for _, opt := range options {
		opt(option)
	}

	return option
}
