package programs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art-fs/tools"
)

func DrawAsciiArtOutput(elements []string, arg2 string) string {
	var result string
	words := strings.Split(arg2, `\n`)
	for _, word := range words {
		// replace empty string by new line
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
					} else {
						// detect the line from where we should start reading
						start := int(char-32)*8 + j
						result += (elements[start])
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

func Output() {
	arg1, arg2, _ := tools.ReplaceArgsO()
	var data string
	if len(os.Args[1:]) == 3 {
		// if have to argements we well have choice to work with any template
		data = tools.CHeckTemplateO()
	} else if len(os.Args[1:]) == 2 {
		// if have to argements we well have choice to work with any template
		data = tools.CHeckTemplate()
	} else {
		// we have just one argement we well work with one argement only
		data = tools.Read_Input("Templates/standard.txt")
	}
	// Split by newline and after that delete the empty strings to organise the file
	var elements []string
	data = strings.ReplaceAll(string(data[1:]), "\r", "\n")
	elements = strings.Split(string(data[1:]), "\n")
	elements = tools.RemoveEmptyString(elements)
	// Split the argument by new line to check every one
	result := DrawAsciiArtOutput(elements, arg2)
	// handling the additionnel new line if the arguiment is a bunche of new lines
	if tools.IsAllNl(result) {
		result = result[1:]
	}
	// Printing final result
	if len(os.Args[1:]) <= 2 {
		fmt.Print(result)
	} else {
		tools.StoreResult(arg1, result)
	}
}
