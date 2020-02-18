# What is this?

A demonstration of using command line execution in Golang to open a file with `$EDITOR`, and then
read back the contents once the file has been closed by the editor.

Examples of where this is useful is for interactively editing application configuration files. Kubernetes
does this with the `kubectl edit` command, for example.

Brief demonstation below:

<a href="https://asciinema.org/a/bUMdA0HebM3oB04gOsiRZfG9u" target="_blank"><img src="https://asciinema.org/a/bUMdA0HebM3oB04gOsiRZfG9u.svg" /></a>

## Usage

`go get github.com/halkyon/go-open-with-editor`

`$ go-open-with-editor`

This will create a temporary file `test12345.txt` in your system temp path.

