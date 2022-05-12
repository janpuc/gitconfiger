package main

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	testFile := "test"

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
