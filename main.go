package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// defaultGitConfig := ".gitconfig"
	// var input string

	// createFile(defaultGitConfig)
	// writeToFile(defaultGitConfig, "test")
	// fileExists := doesFileExist("test")
	// fmt.Println(fileExists)

	// if doesFileExist(defaultGitConfig) {
	// 	log.Fatal("not implemented")
	// }

	// fmt.Print("Do you want to create a multi-account config? (y/n): ")
	// fmt.Scan(&input)
	// // fmt.Print("\n")
	// fmt.Println(input)

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
