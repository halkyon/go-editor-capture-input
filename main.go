package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	dir := os.TempDir()
	path := filepath.Join(dir, "test12345.txt")
	defer os.Remove(path)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		panic(err)
	}
	f.Write([]byte("hello\n"))
	// close here immediately, otherwise the subsequent execution won't succeed as the file is already open.
	f.Close()

	// todo: file on windows with correct file type association, see "ftype" and "assoc" commands
	args := []string{}
	if runtime.GOOS == "windows" {
		args = append(args, "cmd", "/C")
	} else {
		args = append(args, "edit")
	}
	args = append(args, path)

	cmd := exec.Command(args[0], args[1:]...)
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

	fmt.Printf("%s\n", data)
}
