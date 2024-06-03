package programs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art-fs/tools"
)

// Function to draw ASCII art from elements based on input string
func DrawAsciiArt(elements []string, input string) string {
	var result string
	lines := strings.Split(input, `\n`)
	for _, line := range lines {
		if line != "" {
			for j := 0; j < 8; j++ {
				for _, char := range line {
					if char < 32 || char > 126 {
						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
					} else {
						start := int(char-32)*8 + j
						result += elements[start]
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}

func AsciiArt(isOutput bool) {
	args := os.Args[1:]
	input := args[0]
	var template string

	if isOutput {
		if len(args) == 3 {
			template = args[3]
		}else{
			template = "standard"
		}
	} else {
		if len(args) == 2 {
			template = args[1]
		} else {
			template = "standard"
		}
	}

	// Read template data
	data := tools.CheckTemplate(template)
	data = strings.ReplaceAll(data, "\r", "\n")
	elements := strings.Split(data, "\n")
	elements = tools.RemoveEmptyStrings(elements)

	// Draw ASCII art
	result := DrawAsciiArt(elements, input)

	// Handle additional new lines
	if tools.IsAllNl(result) {
		result = result[1:]
	}

	if isOutput {
		outputFile := args[1]
		tools.StoreResult(outputFile, result)
	} else {
		fmt.Print(result)
	}
}
