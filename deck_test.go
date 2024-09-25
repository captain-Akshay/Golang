package main

import (
	"os"
	"testing"
)

func TestSaveFile(t *testing.T) {
	d := deck{}
	d.generate()
	d.saveToFile("testing_file")
	_, err := os.OpenFile("testing_file", os.O_RDONLY, 0644)
	if err != nil {
		t.Error("Something went wrong")
	}

}
