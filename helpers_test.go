package main

import (
	"testing"
)

func TestDoesFileExist(t *testing.T) {
	testFile := "test.txt"

	if doesFileExist(testFile) {
		t.Error("File should not exist")
		t.FailNow()
	}

	createFile(testFile)
	if !doesFileExist(testFile) {
		t.Error("File shoud exist")
		t.FailNow()
	}
}
