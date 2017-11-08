package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	len1, err := rot.r.Read(b)
	if err != nil {
		return len1, err
	}

	for i := range b {
		switch {
		case b[i] <= 'M' && b[i] >= 'A':
			fallthrough
		case b[i] <= 'm' && b[i] >= 'a':
			b[i] += 13
		case b[i] <= 'Z' && b[i] >= 'N':
			fallthrough
		case b[i] <= 'z' && b[i] >= 'n':
			b[i] -= 13
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
