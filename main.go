package main

import (
	"bytes"
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
		data := make([]byte, 8)
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			printFormatedLine(str)
			data = data[i+1:]
			str = ""
		}
		str += string(data)
	}
}

func printFormatedLine(str string) {
	formated_string := fmt.Sprintf("read: %s\n", str)
	fmt.Print(formated_string)
}
