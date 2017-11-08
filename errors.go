package main

import (
	"time"
	"fmt"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v: %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work!",
	}
}

func main() {
	if e := run(); e != nil {
		fmt.Println(e)
	}
}
