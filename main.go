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
		str := fmt.Sprintf("read: %s\n", bytes[:n])
		fmt.Print(str)
	}
}
