package main

import (
	"fmt"
	"github.com/riley817/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first" : "First Word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)

	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
