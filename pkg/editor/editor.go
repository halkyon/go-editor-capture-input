package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Editor struct {
	content  []byte
	filename string
	shell    string
	editor   string
}

func New(content []byte, filename string) *Editor {
	shell := defaultShell
	if os.Getenv("SHELL") != "" {
		shell = os.Getenv("SHELL")
	}
	editor := defaultEditor
	if os.Getenv("EDITOR") != "" {
		editor = os.Getenv("EDITOR")
	}

	return &Editor{
		content:  content,
		filename: filename,
		shell:    shell,
		editor:   editor,
	}
}

func (e *Editor) openFile() (*os.File, error) {
	if e.filename != "" {
		return os.OpenFile(filepath.Join(os.TempDir(), e.filename), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	} else {
		return ioutil.TempFile(os.TempDir(), "*")
	}
}

func (e *Editor) Run() ([]byte, error) {
	file, err := e.openFile()
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())

	_, err = file.Write(e.content)
	if err != nil {
		return nil, err
	}

	args := []string{e.shell, shellCommandFlag, fmt.Sprintf("%s %s", e.editor, file.Name())}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	return content, nil
}
