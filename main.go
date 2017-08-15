package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		input  = flag.String("i", "", "the `input file` to remove the BOM from")
		output = flag.String("o", "", "the name of the `output file`")
		bom    = []byte("\xef\xbb\xbf")
	)

	flag.Parse()
	if flag.NFlag() < 1 || *input == "" {
		flag.Usage()
		os.Exit(-1)
	}

	fb, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	if !bytes.Equal(bom, fb[0:3]) {
		// short circuit if file doesn't have bom
		return
	}

	// default to overwriting input file, unless output flag was set
	filename := *input
	if *output != "" {
		filename = *output
	}

	err = ioutil.WriteFile(filename, fb[3:], 0666)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
	return
}
