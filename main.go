package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	gotprint "github.com/andreas-hofmann/gotprint/lib"
)

func main() {
	input := flag.String("i", "", "The input file to read")
	format := flag.String("f", "json", "The Format to parse. Valid options: json")

	flag.Parse()

	if *input == "" {
		log.Fatal("No input file specified.")
	}

	content, err := ioutil.ReadFile(*input)

	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var data interface{}

	switch strings.ToLower(*format) {
	case "json":
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal("Error unmarshalling json: ", err)
		}

	default:
		log.Fatal("Invalid format: ", *format)
	}

	gotprint.Print(data)
}
