// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"testing"

// 	"ascii-art-fs/tools"
// )

// func DrawAsciiArt(elements []string) string {
// 	var result string
// 	words := strings.Split(os.Args[1], `\n`)
// 	for _, word := range words {
// 		// replace empty string by new line
// 		if word != "" {
// 			for j := 0; j < 8; j++ {
// 				for _, char := range word {
// 					if char < 32 || char > 126 {
// 						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
// 					} else {
// 						// detect the line from where we should start reading
// 						start := int(char-32)*8 + j
// 						result += (elements[start])
// 					}
// 				}
// 				result += "\n"
// 			}
// 		} else {
// 			result += "\n"
// 		}
// 	}
// 	return result
// }

// func DrawAsciiArtO(elements []string, arg2 string) string {
// 	var result string
// 	words := strings.Split(arg2, `\n`)
// 	for _, word := range words {
// 		// replace empty string by new line
// 		if word != "" {
// 			for j := 0; j < 8; j++ {
// 				for _, char := range word {
// 					if char < 32 || char > 126 {
// 						log.Fatalln("Error: please provide printable characters!!\nhelp: man ascii")
// 					} else {
// 						// detect the line from where we should start reading
// 						start := int(char-32)*8 + j
// 						result += (elements[start])
// 					}
// 				}
// 				result += "\n"
// 			}
// 		} else {
// 			result += "\n"
// 		}
// 	}
// 	return result
// }

// func AsciiArtFs() {
// 	// Read from the file standard
// 	var data string
// 	if len(os.Args[1:]) == 2 {
// 		// if have to argements we well have choice to work with any template
// 		data = tools.CHeckTemplate()
// 	} else {
// 		// we have just one argement we well work with one argement only
// 		data = tools.Read_Input("Templates/standard.txt")
// 	}
// 	// Split by newline and after that delete the empty strings to organise the file
// 	var elements []string
// 	data = strings.ReplaceAll(string(data[1:]), "\r", "\n")
// 	elements = strings.Split(string(data[1:]), "\n")
// 	elements = tools.RemoveEmptyString(elements)
// 	// Split the argument by new line to check every one
// 	result := DrawAsciiArt(elements)
// 	// handling the additionnel new line if the arguiment is a bunche of new lines
// 	if tools.IsAllNl(result) {
// 		result = result[1:]
// 	}
// 	// Printing final result
// 	fmt.Print(result)
// }

// func Output() string {
// 	var data string
// 	if len(os.Args[1:]) == 3 {
// 		// if have to argements we well have choice to work with any template
// 		data = tools.CHeckTemplateO()
// 	} else {
// 		// we have just one argement we well work with one argement only
// 		data = tools.Read_Input("Templates/standard.txt")
// 	}
// 	// Split by newline and after that delete the empty strings to organise the file
// 	var elements []string
// 	data = strings.ReplaceAll(string(data[1:]), "\r", "\n")
// 	elements = strings.Split(string(data[1:]), "\n")
// 	elements = tools.RemoveEmptyString(elements)
// 	// Split the argument by new line to check every one
// 	result := DrawAsciiArtO(elements, os.Args[2])
// 	// handling the additionnel new line if the arguiment is a bunche of new lines
// 	if tools.IsAllNl(result) {
// 		result = result[1:]
// 	}
// 	// Printing final result
// 	return result
// }

// func LoadTests() []string {
// 	data, err := os.ReadFile("tests_input.txt")
// 	if err != nil {
// 		log.Fatalln("Error :", err)
// 	}
// 	text := strings.Split(string(data), "\n")
// 	return text
// }

// const InputFile string = "thinkertoy.txt"

// // const InputFile string = "standard.txt"
// // const InputFile string = "shadow.txt"
//
//	func Test_main(t *testing.T) {
//		var want string
//		tests := LoadTests()
//		for _, test := range tests {
//			if test == "" {
//				continue
//			}
//			// run the tests in the main and stock the result from stdout.
//			got, err := exec.Command("go", "run", ".", test, InputFile).Output()
//			if err != nil {
//				log.Fatalln(err)
//			}
//			want = Output(test)
//			// Compare the result that the main.go give and the test give if they are the same
//			if want == string(got) {
//				t.Logf(test)
//			} else {
//				t.Fatal(test)
//			}
//		}
//	}
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

// Function to handle ASCII art drawing, either printing to console or saving to file
func AsciiArt(isOutput bool) {
	args := os.Args[1:]
	input := args[0]
	var template string

	if isOutput {
		template = args[2]
	} else {
		template = args[1]
	}

	// Read template data
	data := tools.CHeckTemplate(template)
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
