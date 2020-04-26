package main

import (
	"os"
)

func appendToFile(filename string, content []byte) int {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writtenLen, err := f.Write(content)
	if err != nil {
		panic(err)
	}
	return writtenLen
}

func main() {
	appendToFile("greatFile.txt", []byte("this is a great file. tremendous file. make it great again :)"))
}
