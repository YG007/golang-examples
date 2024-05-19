package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// get filename from arguments
	filename := os.Args[1]

	// Open the particular file using os package
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Read contents from the file and print contents onto the terminal
	io.Copy(os.Stdout, f)

}
