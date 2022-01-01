package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	dstSuffix = "_kive"
)

func unzip(filename string) (bool, error) {

	input := filename
	outputFile := input + dstSuffix

	if split := strings.Split(input, "."); len(split) != 0 {
		outputFile = split[0] + dstSuffix
	}

	f, err := zip.OpenReader(input)
	if err != nil {
		return false, err
	}
	defer f.Close()

	for _, f := range f.File {
		filePath := filepath.Join(outputFile, f.Name)
		fmt.Printf("Unzipping %s\n", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(outputFile)+string(os.PathSeparator)) {
			fmt.Println("Invalid path")
			return false, err
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return false, err
		}

		outputDst, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return false, err
		}

		archFile, err := f.Open()
		if err != nil {
			return false, err
		}

		if _, err := io.Copy(outputDst, archFile); err != nil {
			return false, err
		}

		fmt.Printf("File unzipped successfully. Location: %s/\n", outputFile)
		outputDst.Close()
		archFile.Close()
	}

	return true, nil
}
