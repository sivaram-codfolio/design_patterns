package main

import (
	"fmt"

	"github.com/sivaram-codfolio/design_patterns/iterator"
)

// Iterate using a callback function
func main() {
	// use IterateBooks to iterate via a callback function
	iterator.Lib.IterateBooks(myBookCallback)

	// Use IterateBooks to iterate via anonymous function
	iterator.Lib.IterateBooks(func(b iterator.Book) error {
		fmt.Println("Book author:", b.Author)
		return nil
	})

	// create a BookIterator
	iter := iterator.Lib.CreateIterator()
	for iter.HasNext() {
		book := iter.Next()
		fmt.Printf("Book %+v\n", book)
	}
}

// This callback function processes an individual Book object
func myBookCallback(b iterator.Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
