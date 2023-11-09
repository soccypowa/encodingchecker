package main

import (
	"fmt"
	"os"

	"github.com/saintfish/chardet"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("a file path is required")
		os.Exit(1)
	}
	file, err := os.ReadFile(os.Args[1])
	check(err)
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(file)
	check(err)
	fmt.Printf("Detected charset is: %s", result.Charset)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
