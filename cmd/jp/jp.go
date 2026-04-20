package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/seggewiss/jp/pkg/jp"
)

const version = "0.0.1"

func main() {
	fmt.Println("jp - ", version)

	pretty := flag.Bool("pretty", false, "Pretty print JSON")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: jp [--pretty] <file> <dotNotationPath>")
		os.Exit(1)
	}

	jsonFilePath := args[0]
	dotNotation := args[1]

	res, err := jp.New().ParseJSON(jsonFilePath, dotNotation)
	if err != nil {
		log.Fatal(err)
	}

	var b []byte
	if *pretty {
		b, err = json.MarshalIndent(res, "", "    ")
	} else {
		b, err = json.Marshal(res)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
