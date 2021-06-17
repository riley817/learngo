package main

import (
	"bufio"
	"os"
)

func main() {
	var f = os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "STOP" {
			return
		}
	}
}
