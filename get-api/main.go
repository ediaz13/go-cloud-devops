package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	Contents string
}

func (m MySlowReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func main() {
	mySlowReader := MySlowReader{
		Contents: "Hello World",
	}
	out, err := io.ReadAll(mySlowReader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)

}
