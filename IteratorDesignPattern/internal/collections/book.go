package collections

import "fmt"

// Book represents a book entity
type Book struct {
	Title  string
	Author string
	ISBN   string
	Year   int
}

// NewBook creates a new Book instance
func NewBook(title, author, isbn string, year int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		ISBN:   isbn,
		Year:   year,
	}
}

// String returns a string representation of the book
func (b *Book) String() string {
	return fmt.Sprintf("%s by %s (%d) - ISBN: %s", b.Title, b.Author, b.Year, b.ISBN)
}