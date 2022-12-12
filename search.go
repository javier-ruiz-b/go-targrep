package main

import (
	"archive/tar"
	"bufio"
	"fmt"
	"io"
	"os"
)

const readBufferSize int = 512 * 1024
const maxLineLength int = 16 * 1024

var lineBuffer = make([]byte, maxLineLength)

func searchFileForMatches(fileHeader *tar.Header, fileReader io.Reader, matching MatchingStrategy, prefix string) {
	scanner := bufio.NewScanner(fileReader)
	scanner.Buffer(lineBuffer, maxLineLength)

	for scanner.Scan() {
		if matching.Match(scanner.Text()) {
			if len(prefix) > 0 {
				fmt.Printf("%s:", prefix)
			}
			fmt.Printf("%s: %s\n", fileHeader.Name, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err, "File:", fileHeader.Name)
	}
}

func Search(fileReader io.Reader, config *Config) error {
	bufferedReader := bufio.NewReaderSize(fileReader, readBufferSize)
	tarReader := tar.NewReader(bufferedReader)

	var matchingStrategy MatchingStrategy
	if config.IsRegex {
		var err error
		matchingStrategy, err = NewRegexMatchingStrategy(config.InputString)
		if err != nil {
			return err
		}
	} else {
		matchingStrategy = NewStringMatchingStrategy(config.InputString)
	}

	if config.Verbose {
		fmt.Println("Searching for", config.InputString)
	}

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			if config.Verbose {
				fmt.Fprintf(os.Stderr, "File: %-42s  size: %d M\n",
					header.Name,
					header.Size/(1024*1024))
			}
			searchFileForMatches(header, tarReader, matchingStrategy, config.Prefix)
		default:
			fmt.Fprintf(os.Stderr, "%s : %c %s %s\n",
				"Unable to figure out tar type",
				header.Typeflag,
				"in file",
				header.Name,
			)
		}
	}

	return nil
}
