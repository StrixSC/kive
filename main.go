package main

import (
	"fmt"
    "flag"
	"os"
)

func main() {
    decompressCommand := flag.NewFlagSet("decompress", flag.ExitOnError)
	compressCommand := flag.NewFlagSet("decompress", flag.ExitOnError)
	autoCommand := flag.NewFlagSet("auto", flag.ExitOnError)
	
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	
	switch os.Args[1] {
	case "decompress":
		input := decompressCommand.String("input", "", "Input file")
		output := decompressCommand.String("output", "", "Output file")
		method := decompressCommand.String("method", "", "Method/Algorithm")	
		flag.Parse()
		handleDecompression(input, output, method)
		break
	case "compress":
		input := compressCommand.String("input", "", "Input file")
		output := compressCommand.String("output", "", "Output file")
		method := compressCommand.String("method", "", "Method/Algorithm")	
		flag.Parse()
		handleCompression(input, output, method)
		break
	case "auto":
		input := autoCommand.String("input", "", "Input file")
		flag.Parse()
		handleAutomatic(input)
		break
	default:
		printHelp()
		os.Exit(1)
	}
}

func decompressFile(fileType string, buf []byte, destination string) {
	switch fileType {
	case "zip":
		unzip(buf, destination)
		break
	}
}

func unzip(buffer []byte, destination string) {
    
}

func printHelp() {
    fmt.Println(`
        Kive
    `)
}