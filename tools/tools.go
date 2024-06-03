package tools

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadInput reads the content of the specified file.
func ReadInput(filename string) string {
	buffer, err := os.ReadFile(filename)
	CheckError(err, "Error: failed to read infile: \""+filename+"\"!!")
	if len(buffer) == 0 {
		log.Fatalln("Error: infile is empty: \"" + filename + "\"!!")
	}
	return string(buffer)
}

// RemoveEmptyStrings removes empty strings from a slice of strings.
func RemoveEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// CheckTemplate returns the content of the selected template file based on the argument.
func CheckTemplate(arg string) string {
	var data string
	switch arg {
	case "standard":
		data = ReadInput("Templates/standard.txt")
	case "shadow":
		data = ReadInput("Templates/shadow.txt")
	case "thinkertoy":
		data = ReadInput("Templates/thinkertoy.txt")
	default:
		log.Fatalln("\n          Usage: go run . [OPTION] [STRING] [BANNER]\n          EX: go run . --output=<fileName.txt> something standard")
	}
	return data
}

// StoreResult stores the result string into the specified file.
func StoreResult(filename, content string) {
	if strings.HasPrefix(filename, "--output=") {
		filename = strings.TrimPrefix(filename, "--output=")
	}
	if filename == "main.go" {
		fmt.Println("Cannot store result in main.go")
		return
	}
	file, err := os.Create(filename)
	CheckError(err, "Error creating file: "+filename)
	defer file.Close()
	_, err = file.WriteString(content)
	CheckError(err, "Error writing to file: "+filename)
	fmt.Println("Result stored in", filename)
}

// IsAllNl checks if a string consists solely of newline characters.
func IsAllNl(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}

// CheckError logs a fatal error if the provided error is not nil.
func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
