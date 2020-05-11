package main

import (
	"flag"
	"os"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "game.csv", "Provide a file a csv file in the format of 'question,answer' eg '3+5,8'")

	flag.Parse()

	_, err = os.Open(file)
}
