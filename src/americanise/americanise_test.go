package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	inFilename       = "input.txt"
	expectedFilename = "expected.txt"
	actualFilename   = "actual.txt"
)

func TestAmericanize(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST americanize")

	_, caller, _, _ := runtime.Caller(0)
	path, _ := filepath.Split(caller)

	var inFile, outFile *os.File
	var err error

	inFilename := filepath.Join(path, inFilename)
	if inFile, err = os.Open(inFilename); err != nil {
		t.Fatal(err)
	}
	defer inFile.Close()

	outFileName := filepath.Join(path, actualFilename)
	if outFile, err = os.Create(outFileName); err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	defer os.Remove(outFileName)

	if err := americanise(inFile, outFile); err != nil {
		t.Fatal(err)
	}

	compare(outFileName, filepath.Join(path, expectedFilename), t)
}

func compare(actual, expected string, t *testing.T) {
	if actualBytes, err := ioutil.ReadFile(actual); err != nil {
		t.Fatal(err)
	} else if expectedBytes, err := ioutil.ReadFile(expected); err != nil {
		t.Fatal(err)
	} else {
		if bytes.Compare(actualBytes, expectedBytes) != 0 {
			t.Fatal("actual != expected")
		}
	}
}
