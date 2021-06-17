package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	var result int = 0

	for i := 0; i < len(args); i++ {
		tmp, _ := strconv.Atoi(args[i])
		result += tmp
	}

	fmt.Println("Result : ", result)

}