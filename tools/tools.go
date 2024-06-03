package tools

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Read_Input(infile_name string) string {
	buffer, err := os.ReadFile(infile_name)
	CheckError(err, "Error: failed to read infile: \""+infile_name+"\"!!")
	if len(buffer) == 0 {
		log.Fatalln("Error: infile is empty: \"" + infile_name + "\"!!")
	}
	return string(buffer)
}

func RemoveEmptyString(slice []string) []string {
	var result []string
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			result = append(result, slice[i])
		}
	}
	return result
}

func CHeckTemplate() string {
	var data string
	if os.Args[2] == "standard" {
		data = Read_Input("Templates/standard.txt")
	} else if os.Args[2] == "shadow" {
		data = Read_Input("Templates/shadow.txt")
	} else if os.Args[2] == "thinkertoy" {
		data = Read_Input("Templates/thinkertoy.txt")
	} else {
		log.Fatalln("Usage: go run . [STRING] [BANNER] \nEX: go run . something standard")
	}
	return data
}

func CHeckTemplateO() string {
	var data string
	if os.Args[3] == "standard" || os.Args[3] == "" {
		data = Read_Input("Templates/standard.txt")
	} else if os.Args[3] == "shadow" {
		data = Read_Input("Templates/shadow.txt")
	} else if os.Args[3] == "thinkertoy" {
		data = Read_Input("Templates/thinkertoy.txt")
	} else {
		log.Fatalln("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	return data
}

func StoreResult(filename, content string) {
	if strings.HasPrefix(filename, "--output=") {
		filename := filename[len("--output="):]
		if filename == "main.go" {
			fmt.Println("Can not storing result in main.go")
			return
		}
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Result stored in", filename)
	} else {
		if filename == "main.go" {
			fmt.Println("Can not storing result in main.go")
			return
		}
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Result stored in", filename)
	}
}

func IsAllNl(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg+"\n")
	}
}
