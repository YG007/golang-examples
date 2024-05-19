package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// customer Writer to write contents fof resp
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// getting resp Body from Read method
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Printing response body to terminal
	// io.Copy(os.Stdout, resp.Body)

	// calling io.copy with a customer witer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
