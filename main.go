package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	createFile("test")
}

func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created")
	defer file.Close()
}
