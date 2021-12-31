package main

import (
	"fmt"
	"github.com/h2non/filetype"
	"io/ioutil"
)

type KiveFileHandler func(input string) (bool error)

type FileHandler struct {
	zip  KiveFileHandler
	gzip KiveFileHandler
	tar2 KiveFileHandler
	bzip KiveFileHandler
}

func handleCompression(input *string, output *string, method *string) {
	return
}

func handleDecompression(input *string, output *string, method *string) {
	return
}

func handleAutomatic(input *string) (bool, error) {
	buf, _ := ioutil.ReadFile(*input)
	err := nil
	result := false

	if filetype.IsArchive(buf) {
		fType, err := filetype.Match(buf)
		if err != nil {
			return false, err
		} else {
			fmt.Printf("Detected %s file. Working...", fType.MIME.Subtype)
			handler := FileHandler{
				zip: unzip(input),

				bzip: func(input string) (bool error) {
					return false, nil
				},

				gzip: func(input string) (bool error) {
					return false, nil
				},

				tar2: func(input string) (bool error) {
					return false, nil
				},
			}
		}
	} else {
		fmt.Println("Filetype does not correspond to a valid decompressable file")
	}

	return result, err
}
