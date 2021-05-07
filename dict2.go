package main

import (
	"fmt"
	"github.com/riley817/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first" : "First Word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "Frist")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)

	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(word)
}
