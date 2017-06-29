package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if os.Args[1] == "" {
		log.Fatalln("please provide file path")
	}
	fb, err := ioutil.ReadFile()
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
