package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"ascii-art-output/tools"
)

// Function to draw ascii
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

// Function to read input and template, and printing or storing result
func AsciiArt(input string, isOutput bool) string {
	args := os.Args[1:]
	var template string

	if isOutput {
		input = args[1]
		if len(args) == 3 {
			template = args[2]
		} else {
			template = "standard"
		}
	} else {
		input = args[0]
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

	return result
}

func LoadTests() []string {
	data, err := os.ReadFile("tests_input.txt")
	if err != nil {
		log.Fatalln("Error :", err)
	}
	text := strings.Split(string(data), "\n")
	return text
}

const InputFile string = "Templates/thinkertoy.txt"

// const InputFile string = "standard.txt"
// const InputFile string = "shadow.txt"

func Test_main(t *testing.T) {
	var want string
	tests := LoadTests()
	for _, test := range tests {
		if test == "" {
			continue
		}
		// run the tests in the main and stock the result from stdout.
		got, err := exec.Command("go", "run", ".", "--output=banner.txt", test, InputFile).Output()
		if err != nil {
			log.Fatalln(err)
		}
		want = AsciiArt(test, true)
		// Compare the result that the main.go give and the test give if they are the same
		if want == string(got) {
			t.Logf(test)
		} else {
			t.Fatal(test)
		}
	}
}
