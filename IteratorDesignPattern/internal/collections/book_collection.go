package collections

import (
	"iterator/internal/iterator"
)

// BookCollection represents a collection of books
type BookCollection struct {
	books []*Book
}

// NewBookCollection creates a new BookCollection
func NewBookCollection() *BookCollection {
	return &BookCollection{
		books: make([]*Book, 0),
	}
}

// AddBook adds a book to the collection
func (bc *BookCollection) AddBook(book *Book) {
	bc.books = append(bc.books, book)
}

// RemoveBook removes a book at the specified index
func (bc *BookCollection) RemoveBook(index int) bool {
	if index < 0 || index >= len(bc.books) {
		return false
	}
	bc.books = append(bc.books[:index], bc.books[index+1:]...)
	return true
}

// GetBook returns a book at the specified index
func (bc *BookCollection) GetBook(index int) *Book {
	if index < 0 || index >= len(bc.books) {
		return nil
	}
	return bc.books[index]
}

// Size returns the number of books in the collection
func (bc *BookCollection) Size() int {
	return len(bc.books)
}

// IsEmpty returns true if the collection is empty
func (bc *BookCollection) IsEmpty() bool {
	return len(bc.books) == 0
}

// CreateIterator creates a new iterator for the book collection
func (bc *BookCollection) CreateIterator() iterator.Iterator[*Book] {
	return NewBookIterator(bc)
}