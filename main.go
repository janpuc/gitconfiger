package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	createFile(".gitconfig")
}

func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created")
	defer file.Close()
}
