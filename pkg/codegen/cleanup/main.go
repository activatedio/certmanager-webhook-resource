package main

import (
	"os"
)

func main() {
	if err := os.RemoveAll("./pkg/generated"); err != nil {
		panic(err)
	}
}
