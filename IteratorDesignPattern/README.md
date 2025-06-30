# Iterator Design Pattern in Go

## What is the Iterator Pattern?

The Iterator is a behavioral design pattern that provides a way to access elements of a collection sequentially without exposing the underlying representation of the collection. It encapsulates the traversal logic and provides a uniform interface for iterating over different types of collections.

### Key Characteristics:
- **Sequential Access**: Provides sequential access to elements without exposing internal structure
- **Uniform Interface**: Same interface for different collection types
- **Multiple Iterators**: Multiple iterators can traverse the same collection independently
- **Encapsulation**: Hides the complexity of traversal from the client

## When to Use Iterator Pattern

### Use this pattern when:

1. **Complex Data Structures**
   - When you have complex data structures and want to hide traversal complexity from clients

2. **Multiple Traversal Methods**
   - When you need different ways to traverse the same collection (forward, backward, filtered)

3. **Uniform Access**
   - When you want to provide a uniform interface for traversing different collection types

4. **Memory Efficiency**
   - When you want to traverse large collections without loading all elements into memory at once

5. **Decoupling**
   - When you want to decouple collection classes from traversal algorithms

### Common Use Cases:
- **Database Result Sets**: Iterating through query results
- **File Systems**: Traversing directory structures
- **Data Processing**: Processing large datasets chunk by chunk
- **UI Components**: Iterating through menu items, list elements
- **Configuration Files**: Processing configuration entries

## Implementation Details

### Core Components:

1. **Iterator Interface**: Defines the contract for iteration (HasNext, Next, Reset)
2. **Collection Interface**: Defines the contract for collections that can be iterated
3. **Concrete Iterator**: Implements specific iteration logic
4. **Concrete Collection**: Implements the collection and creates iterators

### Key Implementation Points:

- **Generic Interface**: Uses Go generics for type safety
- **State Management**: Iterator maintains its own traversal state
- **Multiple Iterator Types**: Forward, reverse, and filtered iterators
- **Independent Iterators**: Multiple iterators can work on the same collection

## Project Structure

```
IteratorDesignPattern/
├── go.mod
├── main.go
├── internal/
│   ├── iterator/
│   │   └── iterator.go              # Iterator and Collection interfaces
│   └── collections/
│       ├── book.go                  # Book entity
│       ├── book_collection.go       # Book collection implementation
│       ├── book_iterator.go         # Forward iterator
│       ├── reverse_book_iterator.go # Reverse iterator
│       └── filtered_iterator.go     # Filtered iterator
└── README.md
```

## How It Works

1. **Collection Creation**: Create a collection and add elements
2. **Iterator Creation**: Collection creates an iterator instance
3. **Traversal**: Use HasNext() and Next() to traverse elements
4. **State Management**: Iterator maintains current position internally
5. **Reset Capability**: Iterator can be reset to start over

## Iterator Types Implemented

### 1. Forward Iterator
- Traverses collection from first to last element
- Standard sequential access pattern

### 2. Reverse Iterator
- Traverses collection from last to first element
- Useful for LIFO processing or reverse chronological order

### 3. Filtered Iterator
- Traverses only elements that match a given filter function
- Allows conditional iteration without modifying the collection

## Running the Demo

```bash
cd gocode/IteratorDesignPattern
go run main.go
```

### Expected Output:
The demo shows:
- Forward iteration through all books
- Reverse iteration through all books
- Filtered iteration (books after 2010, Go-related books)
- Iterator reset functionality
- Multiple independent iterators
- Empty collection handling

## Advantages

- **Encapsulation**: Hides internal collection structure
- **Flexibility**: Multiple traversal methods for same collection
- **Independence**: Multiple iterators can work simultaneously
- **Consistency**: Uniform interface across different collection types
- **Memory Efficiency**: Can implement lazy loading

## Disadvantages

- **Overhead**: Additional abstraction layer
- **Complexity**: More complex than simple for loops for basic cases
- **State Management**: Need to manage iterator state carefully

## Go-Specific Features

### Generics
```go
type Iterator[T any] interface {
    HasNext() bool
    Next() T
    Reset()
}
```

### Function Types for Filtering
```go
type FilterFunc func(*Book) bool
```

### Multiple Iterator Types
- Forward, reverse, and filtered iterators all implement the same interface
- Collections can create different types of iterators

## Real-World Examples

- **Database Drivers**: sql.Rows in Go's database/sql package
- **File Processing**: Iterating through lines in large files
- **API Pagination**: Iterating through paginated API responses
- **Stream Processing**: Processing data streams chunk by chunk
- **Configuration Management**: Iterating through configuration sections

## Best Practices

1. **Always Check HasNext()**: Before calling Next()
2. **Handle Empty Collections**: Gracefully handle empty collections
3. **Reset When Needed**: Use Reset() to reuse iterators
4. **Independent State**: Each iterator should maintain its own state
5. **Error Handling**: Consider how to handle errors during iteration

## Summary

The Iterator pattern provides a clean, consistent way to traverse collections without exposing their internal structure. In Go, it's particularly powerful when combined with generics and function types, allowing for flexible, type-safe iteration over various collection types. The pattern is essential for building maintainable, decoupled code that can handle complex data structures efficiently.