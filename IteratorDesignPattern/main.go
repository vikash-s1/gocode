package main

import (
	"fmt"
	"iterator/internal/collections"
)

func main() {
	fmt.Println("=== Iterator Design Pattern Demo ===\n")

	// Create a book collection
	bookCollection := collections.NewBookCollection()

	// Add some books
	books := []*collections.Book{
		collections.NewBook("The Go Programming Language", "Alan Donovan", "978-0134190440", 2015),
		collections.NewBook("Clean Code", "Robert Martin", "978-0132350884", 2008),
		collections.NewBook("Design Patterns", "Gang of Four", "978-0201633612", 1994),
		collections.NewBook("Effective Go", "Go Team", "978-1234567890", 2020),
		collections.NewBook("Concurrency in Go", "Katherine Cox-Buday", "978-1491941195", 2017),
		collections.NewBook("Go in Action", "William Kennedy", "978-1617291784", 2015),
	}

	for _, book := range books {
		bookCollection.AddBook(book)
	}

	fmt.Printf("Collection size: %d books\n\n", bookCollection.Size())

	// Demo 1: Basic Forward Iterator
	fmt.Println("=== Forward Iterator ===")
	iterator := bookCollection.CreateIterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Printf("📖 %s\n", book.String())
	}
	fmt.Println()

	// Demo 2: Reverse Iterator
	fmt.Println("=== Reverse Iterator ===")
	reverseIterator := bookCollection.CreateReverseIterator()
	for reverseIterator.HasNext() {
		book := reverseIterator.Next()
		fmt.Printf("📖 %s\n", book.String())
	}
	fmt.Println()

	// Demo 3: Filtered Iterator - Books published after 2010
	fmt.Println("=== Filtered Iterator (Books after 2010) ===")
	modernBooksFilter := func(book *collections.Book) bool {
		return book.Year > 2010
	}
	filteredIterator := bookCollection.CreateFilteredIterator(modernBooksFilter)
	for filteredIterator.HasNext() {
		book := filteredIterator.Next()
		fmt.Printf("📖 %s\n", book.String())
	}
	fmt.Println()

	// Demo 4: Filtered Iterator - Go books
	fmt.Println("=== Filtered Iterator (Go Books) ===")
	goBooksFilter := func(book *collections.Book) bool {
		title := book.Title
		return len(title) >= 2 && (title[:2] == "Go" || title[len(title)-2:] == "Go")
	}
	goIterator := bookCollection.CreateFilteredIterator(goBooksFilter)
	for goIterator.HasNext() {
		book := goIterator.Next()
		fmt.Printf("📖 %s\n", book.String())
	}
	fmt.Println()

	// Demo 5: Reset and reuse iterator
	fmt.Println("=== Iterator Reset Demo ===")
	fmt.Println("First iteration (first 3 books):")
	iterator.Reset()
	count := 0
	for iterator.HasNext() && count < 3 {
		book := iterator.Next()
		fmt.Printf("📖 %s\n", book.String())
		count++
	}

	fmt.Println("\nAfter reset (all books):")
	iterator.Reset()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Printf("📖 %s\n", book.String())
	}
	fmt.Println()

	// Demo 6: Multiple iterators on same collection
	fmt.Println("=== Multiple Iterators Demo ===")
	iter1 := bookCollection.CreateIterator()
	iter2 := bookCollection.CreateIterator()

	fmt.Println("Iterator 1 - First book:")
	if iter1.HasNext() {
		fmt.Printf("📖 %s\n", iter1.Next().String())
	}

	fmt.Println("Iterator 2 - First book:")
	if iter2.HasNext() {
		fmt.Printf("📖 %s\n", iter2.Next().String())
	}

	fmt.Println("Iterator 1 - Second book:")
	if iter1.HasNext() {
		fmt.Printf("📖 %s\n", iter1.Next().String())
	}
	fmt.Println()

	// Demo 7: Empty collection handling
	fmt.Println("=== Empty Collection Demo ===")
	emptyCollection := collections.NewBookCollection()
	emptyIterator := emptyCollection.CreateIterator()
	
	fmt.Printf("Empty collection size: %d\n", emptyCollection.Size())
	fmt.Printf("Empty collection is empty: %t\n", emptyCollection.IsEmpty())
	fmt.Printf("Empty iterator has next: %t\n", emptyIterator.HasNext())
	
	if !emptyIterator.HasNext() {
		fmt.Println("No books to iterate over in empty collection")
	}
}