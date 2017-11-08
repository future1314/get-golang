package main

import (
	"os"
	"fmt"
)

type Reader interface {
	Read(b []byte) (x int, err error)
}
type Writer interface {
	Write(b []byte) (x int, err error)
}
type ReaderWriter interface {
	Reader
	Writer
}

func main() {
//	var w Writer
	var w ReaderWriter
	w = os.Stdout
	fmt.Fprintln(w, "Hello, world!")
}
