package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	defaultEditor        = "vi"
	defaultEditorWindows = "notepad"
)

func main() {
	editor := defaultEditor
	if runtime.GOOS == "windows" {
		editor = defaultEditorWindows
	}
	// todo check if spaces won't mess this up with command execution
	if os.Getenv("EDITOR") != "" {
		editor = os.Getenv("EDITOR")
	}

	dir := os.TempDir()
	path := filepath.Join(dir, "test12345.txt")
	err := ioutil.WriteFile(path, []byte("hello\n"), 0600)
	if err != nil {
		panic(err)
	}
	defer os.Remove(path)

	// todo: on Windows, use file association to figure out what to open the file with
	// using "start", "ftype" and "assoc"
	cmd := exec.Command(editor, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Printf("Running %s\n", cmd)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Contents:")
	fmt.Printf("%s\n", data)
}
