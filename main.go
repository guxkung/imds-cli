package main

import (
	"os"
	"guxkung.com/curl-in-go"
)

func ErrorHandling() {
}

func main() {
	args := os.Args

	if len(args) < 2 {
		os.Exit(1)
	}
	code := curl.RequestV1(args[1])
}
