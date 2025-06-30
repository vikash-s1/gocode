package collections

// BookIterator implements the Iterator interface for BookCollection
type BookIterator struct {
	collection *BookCollection
	index      int
}

// NewBookIterator creates a new BookIterator
func NewBookIterator(collection *BookCollection) *BookIterator {
	return &BookIterator{
		collection: collection,
		index:      0,
	}
}

// HasNext returns true if there are more elements to iterate
func (bi *BookIterator) HasNext() bool {
	return bi.index < bi.collection.Size()
}

// Next returns the next book in the collection
func (bi *BookIterator) Next() *Book {
	if !bi.HasNext() {
		return nil
	}
	book := bi.collection.GetBook(bi.index)
	bi.index++
	return book
}

// Reset resets the iterator to the beginning
func (bi *BookIterator) Reset() {
	bi.index = 0
}