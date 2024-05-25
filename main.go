package main

import (
	"fmt"
	"os"

	"ascii-art-fs/programs"
)

func main() {
	if len(os.Args[1:]) > 3 || len(os.Args) == 1 {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}
	if len(os.Args[1:]) <= 3 {
		programs.Output()
	}
}
