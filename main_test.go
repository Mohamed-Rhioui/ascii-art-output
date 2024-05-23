package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"ascii-art-fs/tools"
)

// function for drawing ascii string
func DrawAsciiArt(elements []string, text string) string {
	var result string
	words := strings.Split(text, `\n`)
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

func AsciiArtFs(text string) string {
	// Read from the file standard

	// we have just one argement we well work with one argement only
	data := tools.Read_Input("Templates/" + infile)

	// Split by newline and after that delete the empty strings to organise the file
	var elements []string
	data = strings.ReplaceAll(string(data[1:]), "\r", "\n")
	elements = strings.Split(string(data[1:]), "\n")
	elements = tools.RemoveEmptyString(elements)
	// Split the argument by new line to check every one
	result := DrawAsciiArt(elements, text)
	// handling the additionnel new line if the arguiment is a bunche of new lines
	if tools.IsAllNl(result) {
		result = result[1:]
	}
	// Printing final result
	return result
}

func LoadTests() []string {
	testfile, err := os.Open("tests_input.txt")
	tools.CheckError(err, "Error opening testfile: \"tests_input.txt\"")
	defer testfile.Close()
	var tests []string
	scanner := bufio.NewScanner(testfile)
	for scanner.Scan() {
		tests = append(tests, scanner.Text())
	}
	tools.CheckError(scanner.Err(), "scanner error")
	return tests
}

const infile = "standard.txt"

func Test_main(t *testing.T) {
	// Read the file test and check it .
	tests := LoadTests()
	for _, test := range tests {
		// run the tests in the main and stock the result from stdout.
		got, err := exec.Command("go", "run", ".", test, infile).Output()
		tools.CheckError(err, "")
		want := AsciiArtFs(test)

		// Compare the result that the main.go give and the test give if they are the same
		if want == string(got) {
			t.Logf(test)
		} else {
			t.Fatal(test)
		}
	}
}
