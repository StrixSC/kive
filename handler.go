package main

import (
	"fmt"
	"io/ioutil"
	"github.com/h2non/filetype"
)

func handleCompression(input *string, output *string, method *string) {
	fmt.Printf("Compressing file %s into %s using %s", input, output, method)
}

func handleDecompression(input *string, output *string, method *string) {
	fmt.Println("Decompressing file %s into %s using %s", input, output, method)
}

func handleAutomatic(input *string) {
	buf, _ := ioutil.ReadFile(*input)

	if filetype.IsArchive(buf) {
		fType, err := filetype.Match(buf)
		if err != nil {
			fmt.Println("Error while trying to obtain the filetype")
		} else {
			fmt.Printf("Detected %s archive file. Decompressing...", fType.MIME.Subtype)
		}

	} else {
		fmt.Println("Filetype does not correspond to a valid decompressable file")
	}
}