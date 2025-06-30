package iterator

// Iterator defines the interface for iterating over a collection
type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Reset()
}

// Collection defines the interface for collections that can be iterated
type Collection[T any] interface {
	CreateIterator() Iterator[T]
	Size() int
	IsEmpty() bool
}