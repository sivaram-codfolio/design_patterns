package iterator

// The IterableCollection interface defines the CreateIterator
// function, which returns an iterator object
type IterableCollection interface {
	CreateIterator() iterator
}

// The iterator interface contains the hasNext and next functions
// which allow the collection to return items as needed
type iterator interface {
	HasNext() bool
	Next() *Book
}

// BookIterator is a concrete iterator for a Book collection
type BookIterator struct {
	Current int
	Books   []Book
}

func (b *BookIterator) HasNext() bool {
	return b.Current < len(b.Books)
}

func (b *BookIterator) Next() *Book {
	if b.HasNext() {
		bk := b.Books[b.Current]
		b.Current++
		return &bk
	}
	return nil
}
