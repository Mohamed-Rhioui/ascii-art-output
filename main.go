package main

import (
	"os"
	"strings"

	"ascii-art-fs/programs"
)

func main() {
	if len(os.Args[1:]) <= 3 {
		if len(os.Args[1:]) == 3 || (len(os.Args[1:]) == 2 && (strings.HasPrefix(os.Args[1], "--output="))) {
			programs.Output()
		}else if len(os.Args[1:]) <= 2 && !(strings.HasPrefix(os.Args[1], "--output=")){
			programs.AsciiArtFs()
		}
	}
}
