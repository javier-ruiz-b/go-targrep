package main

type MatchingStrategy interface {
	Match(string) bool
}
