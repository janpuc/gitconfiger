package main

import (
	"errors"
	"log"
	"os"
)

func doesFileExist(name string) bool {
	_, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}
