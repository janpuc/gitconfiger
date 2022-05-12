package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	testFile := "test.txt"

	createFile(testFile)
	_, err := os.Stat(testFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = os.Remove(testFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestWriteToFile(t *testing.T) {
	testFile := "test.txt"
	testBody := "test"

	createFile(testFile)
	_, err := os.Stat(testFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	writeToFile(testFile, testBody)
	writtenBody, err := os.ReadFile(testFile)
	if err != nil {
		t.Error(err)
	}

	if testBody != string(writtenBody) {
		t.Errorf("Supposed to write: %s, written: %s", testBody, writtenBody)
	}

	err = os.Remove(testFile)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
