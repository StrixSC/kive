package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/h2non/filetype"
)

const (
	unhandledFiletype = "UnhandledFiletype"
)

type FileHandlerCommands struct {
	name     string
	commands []string
}

type FileHandler map[string]FileHandlerCommands

var KiveFileHandler = FileHandler{
	"zip":   FileHandlerCommands{name: "unzip", commands: []string{}},
	"x-tar": FileHandlerCommands{name: "tar", commands: []string{"xvf"}},
}

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

	if len(os.Args) <= 2 {
		printHelp()
		return
	}

	switch os.Args[1] {

	case "decompress":
		decompressCmd.Parse(os.Args[2:])
		err := handleDecompression(*decompressInput, *decompressOutput, *decompressMethod)
		if err != nil {
			log.Fatal(err)
		}

	case "compress":
		compressCmd.Parse(os.Args[2:])
		err := handleCompression(*compressInput, *compressOutput, *compressMethod)
		if err != nil {
			log.Fatal(err)
		}

	case "auto":
		autoCmd.Parse(os.Args[2:])
		err := checkFileExists(*autoInput)
		if err != nil {
			log.Fatal(err)
		}
		err = handleAutomatic(*autoInput)
		if err != nil {
			log.Fatal(err)
		}

	default:
		printHelp()
	}
}

func checkFileExists(filename string) error {
	_, err := os.Stat(filename)
	return err
}

func handleCompression(input string, output string, method string) error {
	return nil
}

func handleDecompression(input string, output string, method string) error {
	return nil
}

func handleAutomatic(filename string) error {
	buf, _ := ioutil.ReadFile(filename)
	if filetype.IsArchive(buf) {
		fType, err := filetype.Match(buf)
		if err != nil {
		} else {
			t := fType.MIME.Subtype
			fmt.Printf("Detected %s file. Working...\n", t)
			if kiveHandler, exists := KiveFileHandler[t]; !exists {
				return errors.New(unhandledFiletype)
			} else {
				return handle(filename, kiveHandler)
			}
		}
	} else {
		return errors.New(unhandledFiletype)
	}
	return nil
}

func handle(filename string, handler FileHandlerCommands) error {
	commands := append(handler.commands, filename)
	output, err := exec.Command(handler.name, commands...).Output()
	if err != nil {
		return err
	}

	fmt.Println(string(output[:]))
	return nil
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
