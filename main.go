package main

import (
	"log"
	"os"
	"strings"

	"ascii-art-output/programs"
)

func main() {
	args := os.Args[1:]
	// Protection
	if len(args) > 3 || len(args) == 0 {
		log.Fatalln("\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}
	outputToFile := false
	if strings.HasPrefix(args[0], "--output=") {
		outputToFile = true
	}

	// Check if the output file already exists
	if outputToFile {
		if len(args) == 3 || (len(args) == 2 && strings.HasPrefix(args[0], "--output=")) {
			programs.AsciiArt(true)
		} else if len(args) == 1 && strings.HasPrefix(args[0], "--output=") {
			programs.AsciiArt(false )
		} else {
			log.Fatalln("\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		}
	} else {
		if len(args) == 1 || len(args) == 2 {
			if strings.Contains(args[0],"--output") {
				log.Fatalln("\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			}
			programs.AsciiArt(false)
		} else {
			log.Fatalln("\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		}
	}
}
