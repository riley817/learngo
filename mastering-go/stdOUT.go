package main

import (
	"io"
	"os"
)

func main() {
	str := ""
	arg := os.Args
	if len(arg) == 1 {
		str = "Please give me one argument!"
	} else {
		str = arg[1]
	}

	io.WriteString(os.Stdout, str)
	io.WriteString(os.Stdout, "\n")

}