# go-targrep

Basic tool which searches for strings in tar files

## Usage

Usage:    TARSTREAM | targrep REGEX
Example:  cat test.tar | targrep -r '^test[0-9]+'
  -p string
        Prefix for each line
  -r    Input string is a regex
  -v    Verbose mode
