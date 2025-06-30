package collections

// FilterFunc defines a function type for filtering books
type FilterFunc func(*Book) bool

// FilteredBookIterator implements filtered iteration for BookCollection
type FilteredBookIterator struct {
	collection *BookCollection
	filter     FilterFunc
	index      int
}

// NewFilteredBookIterator creates a new FilteredBookIterator
func NewFilteredBookIterator(collection *BookCollection, filter FilterFunc) *FilteredBookIterator {
	return &FilteredBookIterator{
		collection: collection,
		filter:     filter,
		index:      0,
	}
}

// HasNext returns true if there are more filtered elements to iterate
func (fbi *FilteredBookIterator) HasNext() bool {
	for fbi.index < fbi.collection.Size() {
		book := fbi.collection.GetBook(fbi.index)
		if fbi.filter(book) {
			return true
		}
		fbi.index++
	}
	return false
}

// Next returns the next book that matches the filter
func (fbi *FilteredBookIterator) Next() *Book {
	if !fbi.HasNext() {
		return nil
	}
	book := fbi.collection.GetBook(fbi.index)
	fbi.index++
	return book
}

// Reset resets the iterator to the beginning
func (fbi *FilteredBookIterator) Reset() {
	fbi.index = 0
}

// CreateFilteredIterator creates a filtered iterator for the book collection
func (bc *BookCollection) CreateFilteredIterator(filter FilterFunc) *FilteredBookIterator {
	return NewFilteredBookIterator(bc, filter)
}