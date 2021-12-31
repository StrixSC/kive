package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	filetypeZip  = "zip"
	filetypeGZip = "gzip"
)

func main() {
	decompressCmd := flag.NewFlagSet("decompress", flag.ExitOnError)
	compressCmd := flag.NewFlagSet("decompress", flag.ExitOnError)
	autoCmd := flag.NewFlagSet("auto", flag.ExitOnError)

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "decompress":
		input := decompressCmd.String("input", "", "Input file")
		output := decompressCmd.String("output", "", "Output file")
		method := decompressCmd.String("method", "", "Method/Algorithmc")
		decompressCmd.Parse(os.Args[2:])
		handleDecompression(input, output, method)
		break
	case "compress":
		input := compressCmd.String("input", "", "Input file")
		output := compressCmd.String("output", "", "Output file")
		method := compressCmd.String("method", "", "Method/Algorithm")
		compressCmd.Parse(os.Args[2:])
		handleCompression(input, output, method)
		break
	case "auto":
		input := autoCmd.String("input", "", "Input file")
		autoCmd.Parse(os.Args[2:])
		fmt.Println("File: ", *input)
		handleAutomatic(input)
		break
	default:
		printHelp()
		os.Exit(1)
	}
}

func decompressFile(fileType string, filename string, destination string) {
	switch fileType {
	case "zip":
		_, err := unzip(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		break
	}
}

func printHelp() {
	fmt.Println(`
        Kive
    `)
}
