package main

import (
	"fmt"
	"log"
	"os"

	"github.com/seggewiss/jp/pkg/jp"
)

const version = "0.0.1"

func main() {
	fmt.Println("jp - ", version)

	if len(os.Args) < 3 {
		fmt.Println("Usage: jp <file> <dotNotationPath>")
		os.Exit(1)
	}

	jsonFilePath := os.Args[1]
	dotNotation := os.Args[2]

	res, err := jp.New().ParseJSON(jsonFilePath, dotNotation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
