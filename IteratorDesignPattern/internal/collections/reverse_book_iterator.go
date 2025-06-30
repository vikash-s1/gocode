package collections

// ReverseBookIterator implements reverse iteration for BookCollection
type ReverseBookIterator struct {
	collection *BookCollection
	index      int
}

// NewReverseBookIterator creates a new ReverseBookIterator
func NewReverseBookIterator(collection *BookCollection) *ReverseBookIterator {
	return &ReverseBookIterator{
		collection: collection,
		index:      collection.Size() - 1,
	}
}

// HasNext returns true if there are more elements to iterate (in reverse)
func (rbi *ReverseBookIterator) HasNext() bool {
	return rbi.index >= 0
}

// Next returns the next book in reverse order
func (rbi *ReverseBookIterator) Next() *Book {
	if !rbi.HasNext() {
		return nil
	}
	book := rbi.collection.GetBook(rbi.index)
	rbi.index--
	return book
}

// Reset resets the iterator to the end (for reverse iteration)
func (rbi *ReverseBookIterator) Reset() {
	rbi.index = rbi.collection.Size() - 1
}

// CreateReverseIterator creates a reverse iterator for the book collection
func (bc *BookCollection) CreateReverseIterator() *ReverseBookIterator {
	return NewReverseBookIterator(bc)
}