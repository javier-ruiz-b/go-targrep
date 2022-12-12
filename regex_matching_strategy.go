package main

import "regexp"

type RegexMatchingStrategy struct {
	regex *regexp.Regexp
}

func NewRegexMatchingStrategy(inputString string) (*RegexMatchingStrategy, error) {
	regex, err := regexp.Compile(inputString)

	return &RegexMatchingStrategy{
		regex: regex,
	}, err
}

func (r *RegexMatchingStrategy) Match(input string) bool {
	return r.regex.MatchString(input)
}
