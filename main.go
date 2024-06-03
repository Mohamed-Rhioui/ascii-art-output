package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art-fs/programs"
)

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) == 0 {
		log.Fatalln("\n          Usage: go run . [OPTION] [STRING] [BANNER]\n          EX: go run . --output=<fileName.txt> something standard")
	}
	outputToFile := false
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			outputToFile = true
			break
		}
	}
	if outputToFile {
		fmt.Println(args)
		if len(args) == 3 || (len(args) == 2 && strings.HasPrefix(args[1], "--output=")) { 
			programs.AsciiArt(true)
		} else {
			log.Fatalln("\n          Usage: go run . [OPTION] [STRING] [BANNER]\n          EX: go run . --output=<fileName.txt> something standard")
		}
	} else {
		if len(args) == 1 || len(args) == 2 { 
			programs.AsciiArt(false)
		} else {
			log.Fatalln("\n          Usage: go run . [OPTION] [STRING] [BANNER]\n          EX: go run . --output=<fileName.txt> something standard")
		}
	}
}
