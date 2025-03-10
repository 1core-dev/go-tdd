package dictionary

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")

// Dictionary store definitions to words.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
