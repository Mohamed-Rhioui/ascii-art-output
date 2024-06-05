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
func AsciiArt( input string, template string) string {
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

// LoadTests reads the test inputs from a file.
func LoadTests() []string {
	data, err := os.ReadFile("tests_input.txt")
	if err != nil {
		log.Fatalln("Error reading test inputs:", err)
	}
	text := strings.Split(string(data), "\n")
	return text
}

func TestMain(t *testing.T) {
	template := "standard"
	tests := LoadTests()
	for _, test := range tests {
		if test == "" {
			continue
		}
		cmd := exec.Command("go", "run", ".","--output=banner.txt", test, template)
		err := cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}

		got, err := os.ReadFile("banner.txt")
		if err != nil {
			log.Fatalln(err)
		}

		// want := AsciiArt(true, test, template)

		want := AsciiArt( test, template)

		// Compare the result that the main.go give and the test give if they are the same
		if want == string(got) {
			t.Logf(test)
		} else {
			t.Fatal(test)
		}
	}
}
