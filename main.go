package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	env := os.Environ()

	println("Injected DEMO* env variables: ")
	for _, v := range env {
		if strings.HasPrefix(v, "DEMO") {
			println("\t", v)
		}
	}

	println("")
	println("Injected files in /tmp/*.yaml:")
	files, err := ioutil.ReadDir("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), "yaml") {
			println("--- ", file.Name())
			content, err := ioutil.ReadFile("/tmp/" + file.Name())
			if err != nil {
				log.Print("error reading file")
			}
			println("", string(content), "")
		}
	}
}
