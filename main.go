package main

import (
	"fmt"

	"github.com/halkyon/go-editor-capture-input/pkg/editor"
)

func main() {
	editor := editor.New([]byte("hello\n"), "test12345.txt")
	output, err := editor.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", output)
}
