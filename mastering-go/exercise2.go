package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var result float64 = 0
	var total float64 = 0

	for i := 0; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[i], 64)
		total += n
	}

	result = total / float64(len(args))
	fmt.Println("Result : ", result)
}
