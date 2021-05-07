package mydict

import "errors"

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exsits")
var errCantupdate = errors.New("Cant't update non-exsiting word")
// Dictionary Type
type Dictionary map[string]string

// Search for a word
// Maps are reference types, so they are always passed by reference. You don't need a pointer.
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantupdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}