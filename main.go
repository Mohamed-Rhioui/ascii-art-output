package main

import (
	"log"
	"os"
	"strings"

	"ascii-art-fs/programs"
)

func main() {
	if len(os.Args[1:]) > 3 || len(os.Args) == 1 {
		log.Fatalln("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	if len(os.Args[1:]) <= 3 {
		if len(os.Args[1:]) <= 2 && !(strings.HasPrefix(os.Args[1], "--output=")) {
			programs.AsciiArtFs()
		} else if len(os.Args[1:]) == 3 || (len(os.Args[1:]) == 2 && (strings.HasPrefix(os.Args[1], "--output="))) {
			programs.Output()
		} else if len(os.Args[1:]) == 1 && (strings.HasPrefix(os.Args[1], "--output=")) {
			programs.AsciiArtFs()
		}
	}
}
