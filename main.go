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
	filetypeTar  = "tar"
	filetypeBZip = "bzip2"
)

func main() {
	decompressCmd := flag.NewFlagSet("decompress", flag.ExitOnError)
	decompressInput := decompressCmd.String("input", "", "Input file")
	decompressOutput := decompressCmd.String("output", "", "Output file")
	decompressMethod := decompressCmd.String("method", "", "Method/Algorithm")

	compressCmd := flag.NewFlagSet("compress", flag.ExitOnError)
	compressInput := compressCmd.String("input", "", "Input file")
	compressOutput := compressCmd.String("output", "", "Output file")
	compressMethod := compressCmd.String("method", "", "Method/Algorithm")

	autoCmd := flag.NewFlagSet("auto", flag.ExitOnError)
	autoInput := autoCmd.String("input", "", "Input file")

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {

	case "decompress":
		decompressCmd.Parse(os.Args[2:])
		_, err := handleDecompression(*decompressInput, *decompressOutput, *decompressMethod)
		if err != nil {
			log.Fatal(err)
		}

	case "compress":
		compressCmd.Parse(os.Args[2:])
		_, err := handleCompression(*compressInput, *compressOutput, *compressMethod)
		if err != nil {
			log.Fatal(err)
		}

	case "auto":
		autoCmd.Parse(os.Args[2:])
		fmt.Println("File: ", *autoInput)
		_, err := handleAutomatic(*autoInput)
		if err != nil {
			log.Fatal(err)
		}

	default:
		printHelp()
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
         _  ___  _   _  ___
        | |/ / || \ / || __|
        |   <| | \ V / | _|
        |_|\_\_|  \_/  |___|

        /kē'vā/
        Command line automatic multi-format compression/decompression & archiving/unarchiving utility.
    `)

	flag.PrintDefaults()
}
