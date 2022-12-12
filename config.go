package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Verbose     bool
	IsRegex     bool
	Prefix      string
	InputString string
}

func LoadConfig() (*Config, error) {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\n%s\n\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:    TARSTREAM | targrep REGEX\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Example:  cat test.tar | targrep '^test[0-9]+'\n")

		flag.PrintDefaults()
	}
	verbose := flag.Bool("v", false, "Verbose mode")
	isRegex := flag.Bool("r", false, "Input string is a regex")
	prefix := flag.String("p", "", "Prefix for each line")
	flag.Parse()

	if flag.NArg() == 0 {
		return nil, errors.New("missing input string")
	}

	if !stdinComesFromPipe() {
		return nil, errors.New("missing stdin stream")
	}

	return &Config{
		Verbose:     *verbose,
		IsRegex:     *isRegex,
		Prefix:      *prefix,
		InputString: strings.Join(flag.Args(), " "),
	}, nil
}

func stdinComesFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return (fi.Mode() & os.ModeCharDevice) == 0
}
