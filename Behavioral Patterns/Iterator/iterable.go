package main

// Iterable collection (aggregator) interface
type IterableCollection interface {
	createIterator() iterator
}

// Iterator interface
type iterator interface {
	hasNext() bool
	next() *Book
}

// Concrete Iterator
type BookIterator struct {
	current int
	books   []Book
}

func (b *BookIterator) hasNext() bool {
	if b.current < len(b.books) {
		return true
	}
	return false
}

func (b *BookIterator) next() *Book {
	// TODO: implement next()
	if b.hasNext() {
		bk := b.books[b.current]
		b.current++
		return &bk
	}
	return nil
}
