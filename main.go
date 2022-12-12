package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		flag.Usage()
		return
	}

	err = Search(os.Stdin, config)

	if err != nil {
		fmt.Print("Error reading tar stream:", err)
	}
}
