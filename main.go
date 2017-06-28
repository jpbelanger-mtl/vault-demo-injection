package main

import (
	"io/ioutil"
	"log"
	"net/http"
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

	println("Listening to /shutdown call")
	http.HandleFunc("/shutdown", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	println("Shutting down the container")
	os.Exit(0)
}
