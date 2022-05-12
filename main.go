package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defaultGitConfig := ".gitconfig"

	createFile(defaultGitConfig)
	writeToFile(defaultGitConfig, "test")
}

func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created")
	defer file.Close()
}

func writeToFile(name string, body string) {
	file, err := os.OpenFile(name, os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(body)
	fmt.Println("File written to")
}
