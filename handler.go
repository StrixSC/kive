package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/h2non/filetype"
)

func handleCompression(input string, output string, method string) (bool, error) {
	return false, nil
}

func handleDecompression(input string, output string, method string) (bool, error) {
	return false, nil
}

type FileHandler map[string]func(string) (bool, error)

var KiveFileHandler = FileHandler{
	"zip": unzip,
}

func handleAutomatic(input string) (bool, error) {
	buf, _ := ioutil.ReadFile(input)
	if filetype.IsArchive(buf) {
		fType, err := filetype.Match(buf)
		if err != nil {
		} else {
			t := fType.MIME.Subtype
			fmt.Printf("Detected %s file. Working...\n", t)
			if com, exists := KiveFileHandler[t]; !exists {
				return false, errors.New(unhandledFiletype)
			} else {
				return com(input)
			}
		}
	} else {
		return false, errors.New(unhandledFiletype)
	}

	return true, nil
}
