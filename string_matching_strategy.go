package main

import "strings"

type StringMatchingStrategy struct {
	matchingString string
}

func NewStringMatchingStrategy(matchingString string) *StringMatchingStrategy {
	return &StringMatchingStrategy{
		matchingString: matchingString,
	}
}

func (s *StringMatchingStrategy) Match(input string) bool {
	return strings.Contains(input, s.matchingString)
}
