package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "messages.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	var str = ""
	for {
		bytes := make([]byte, 8)
		n, err := file.Read(bytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		str += string(bytes[:n])
	}

	start := 0
	for i, char := range str {
		if char == '\n' {
			printFormatedLine(str[start:i])
			start = i + 1
		}
	}
}

func printFormatedLine(str string) {
	formated_string := fmt.Sprintf("read: %s\n", str)
	fmt.Print(formated_string)
}
