package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var (
		i = flag.String("i", "", "The input file.")
	)

	flag.Parse()

	fb, err := ioutil.ReadFile(*i)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	bom := []byte("\xef\xbb\xbf")

	if !bytes.Equal(bom, fb[0:3]) {
		// short circuit
		return
	}
	err = ioutil.WriteFile(*i, fb[3:], 0666)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
	return
}
