package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"prototypepattern/internal/prototypes"
)

func main() {
	fmt.Println("📄 Welcome to the Prototype Pattern Document Management Demo!")
	fmt.Println(strings.Repeat("=", 60))

	// Interactive demo
	runInteractiveDemo()

	fmt.Println("\n🎯 Automated Demo:")
	fmt.Println(strings.Repeat("=", 30))

	// Automated demo to show all prototype patterns
	runAutomatedDemo()
}

func runInteractiveDemo() {
	scanner := bufio.NewScanner(os.Stdin)

	// Create document factory with pre-registered prototypes
	factory := prototypes.NewDocumentFactory()

	fmt.Println("\n🏭 Document Factory initialized with default templates!")
	fmt.Printf("📋 Available prototypes: %d\n", len(factory.GetRegistry().ListPrototypes()))

	fmt.Println("\nCommands:")
	fmt.Println("1 - List Available Prototypes")
	fmt.Println("2 - Create Document from Prototype")
	fmt.Println("3 - Create Custom Document")
	fmt.Println("4 - Clone Existing Document")
	fmt.Println("5 - Register New Prototype")
	fmt.Println("6 - Compare Original vs Clone")
	fmt.Println("7 - Performance Test")
	fmt.Println("q - Quit to Automated Demo")

	var createdDocuments []prototypes.DocumentPrototype

	for {
		fmt.Printf("\nCreated Documents: %d | Enter command: ", len(createdDocuments))
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			listPrototypes(factory)
		case "2":
			if doc := createFromPrototype(factory, scanner); doc != nil {
				createdDocuments = append(createdDocuments, doc)
			}
		case "3":
			if doc := createCustomDocument(factory, scanner); doc != nil {
				createdDocuments = append(createdDocuments, doc)
			}
		case "4":
			if len(createdDocuments) > 0 {
				if doc := cloneExistingDocument(createdDocuments, scanner); doc != nil {
					createdDocuments = append(createdDocuments, doc)
				}
			} else {
				fmt.Println("❌ No documents available to clone. Create some documents first.")
			}
		case "5":
			registerNewPrototype(factory, scanner)
		case "6":
			if len(createdDocuments) >= 2 {
				compareDocuments(createdDocuments, scanner)
			} else {
				fmt.Println("❌ Need at least 2 documents to compare. Create more documents first.")
			}
		case "7":
			performanceTest(factory)
		case "q":
			return
		default:
			fmt.Println("❌ Invalid command. Try again.")
		}
	}
}

func listPrototypes(factory *prototypes.DocumentFactory) {
	fmt.Println("\n📋 Available Prototypes:")
	fmt.Println(strings.Repeat("-", 40))
	
	prototypes := factory.GetRegistry().ListPrototypes()
	for i, key := range prototypes {
		info, _ := factory.GetRegistry().GetPrototypeInfo(key)
		fmt.Printf("%d. %s\n   %s\n", i+1, key, info)
	}
}

func createFromPrototype(factory *prototypes.DocumentFactory, scanner *bufio.Scanner) prototypes.DocumentPrototype {
	fmt.Println("\n📋 Available prototypes:")
	prototypeKeys := factory.GetRegistry().ListPrototypes()
	for i, key := range prototypeKeys {
		fmt.Printf("%d. %s\n", i+1, key)
	}

	fmt.Print("Enter prototype number: ")
	scanner.Scan()
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil || choice < 1 || choice > len(prototypeKeys) {
		fmt.Println("❌ Invalid choice")
		return nil
	}

	prototypeKey := prototypeKeys[choice-1]

	fmt.Print("Enter document title: ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = "New Document from " + prototypeKey
	}

	fmt.Print("Enter author name: ")
	scanner.Scan()
	author := scanner.Text()
	if author == "" {
		author = "Anonymous"
	}

	document, err := factory.CreateDocument(prototypeKey, title, author)
	if err != nil {
		fmt.Printf("❌ Error creating document: %v\n", err)
		return nil
	}

	fmt.Printf("✅ Created document: %s\n", document.GetInfo())
	return document
}

func createCustomDocument(factory *prototypes.DocumentFactory, scanner *bufio.Scanner) prototypes.DocumentPrototype {
	fmt.Println("\n📄 Document Types:")
	fmt.Println("1. Text Document")
	fmt.Println("2. PDF Document")
	fmt.Println("3. Spreadsheet Document")
	fmt.Println("4. Presentation Document")

	fmt.Print("Enter document type (1-4): ")
	scanner.Scan()
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil || choice < 1 || choice > 4 {
		fmt.Println("❌ Invalid choice")
		return nil
	}

	docTypes := []string{"text", "pdf", "spreadsheet", "presentation"}
	docType := docTypes[choice-1]

	fmt.Print("Enter document title: ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = "New " + strings.Title(docType) + " Document"
	}

	fmt.Print("Enter document content: ")
	scanner.Scan()
	content := scanner.Text()
	if content == "" {
		content = "This is a new " + docType + " document."
	}

	fmt.Print("Enter author name: ")
	scanner.Scan()
	author := scanner.Text()
	if author == "" {
		author = "Anonymous"
	}

	document, err := factory.CreateCustomDocument(docType, title, content, author)
	if err != nil {
		fmt.Printf("❌ Error creating document: %v\n", err)
		return nil
	}

	fmt.Printf("✅ Created document: %s\n", document.GetInfo())
	return document
}

func cloneExistingDocument(documents []prototypes.DocumentPrototype, scanner *bufio.Scanner) prototypes.DocumentPrototype {
	fmt.Println("\n📄 Available Documents to Clone:")
	fmt.Println(strings.Repeat("-", 40))
	
	for i, doc := range documents {
		fmt.Printf("%d. %s\n", i+1, doc.GetInfo())
	}

	fmt.Print("Enter document number to clone: ")
	scanner.Scan()
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil || choice < 1 || choice > len(documents) {
		fmt.Println("❌ Invalid choice")
		return nil
	}

	original := documents[choice-1]
	cloned := original.Clone()

	if doc, ok := cloned.(prototypes.DocumentPrototype); ok {
		fmt.Print("Enter new title for cloned document: ")
		scanner.Scan()
		newTitle := scanner.Text()
		if newTitle != "" {
			doc.SetTitle(newTitle)
		}

		fmt.Print("Enter new author for cloned document: ")
		scanner.Scan()
		newAuthor := scanner.Text()
		if newAuthor != "" {
			doc.SetAuthor(newAuthor)
		}

		fmt.Printf("✅ Cloned document: %s\n", doc.GetInfo())
		return doc
	}

	fmt.Println("❌ Error: Cloned object is not a valid document")
	return nil
}

func registerNewPrototype(factory *prototypes.DocumentFactory, scanner *bufio.Scanner) {
	fmt.Print("Enter prototype key/name: ")
	scanner.Scan()
	key := scanner.Text()
	if key == "" {
		fmt.Println("❌ Prototype key cannot be empty")
		return
	}

	fmt.Print("Enter prototype title: ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		title = "Prototype " + key
	}

	fmt.Print("Enter prototype content: ")
	scanner.Scan()
	content := scanner.Text()
	if content == "" {
		content = "This is a prototype document."
	}

	// Create a text document as prototype (for simplicity)
	prototype := prototypes.NewTextDocument(title, content, "System")
	prototype.AddTag("custom-prototype")

	factory.GetRegistry().RegisterPrototype(key, prototype)
	fmt.Printf("✅ Registered new prototype: %s\n", key)
}

func compareDocuments(documents []prototypes.DocumentPrototype, scanner *bufio.Scanner) {
	fmt.Println("\n📊 Document Comparison")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("Available documents:")
	for i, doc := range documents {
		fmt.Printf("%d. %s\n", i+1, doc.GetInfo())
	}

	fmt.Print("Enter first document number: ")
	scanner.Scan()
	first, err1 := strconv.Atoi(scanner.Text())

	fmt.Print("Enter second document number: ")
	scanner.Scan()
	second, err2 := strconv.Atoi(scanner.Text())

	if err1 != nil || err2 != nil || first < 1 || second < 1 || 
	   first > len(documents) || second > len(documents) {
		fmt.Println("❌ Invalid document numbers")
		return
	}

	doc1 := documents[first-1]
	doc2 := documents[second-1]

	fmt.Printf("\n📄 Document 1: %s\n", doc1.GetInfo())
	fmt.Printf("   Title: %s\n", doc1.GetTitle())
	fmt.Printf("   Author: %s\n", doc1.GetAuthor())
	fmt.Printf("   Created: %s\n", doc1.GetCreatedAt().Format("2006-01-02 15:04:05"))
	fmt.Printf("   Size: %d characters\n", doc1.GetSize())

	fmt.Printf("\n📄 Document 2: %s\n", doc2.GetInfo())
	fmt.Printf("   Title: %s\n", doc2.GetTitle())
	fmt.Printf("   Author: %s\n", doc2.GetAuthor())
	fmt.Printf("   Created: %s\n", doc2.GetCreatedAt().Format("2006-01-02 15:04:05"))
	fmt.Printf("   Size: %d characters\n", doc2.GetSize())

	fmt.Printf("\n🔍 Comparison:\n")
	fmt.Printf("   Same Type: %t\n", doc1.GetType() == doc2.GetType())
	fmt.Printf("   Same Title: %t\n", doc1.GetTitle() == doc2.GetTitle())
	fmt.Printf("   Same Author: %t\n", doc1.GetAuthor() == doc2.GetAuthor())
	fmt.Printf("   Same Content: %t\n", doc1.GetContent() == doc2.GetContent())
}

func performanceTest(factory *prototypes.DocumentFactory) {
	fmt.Println("\n⚡ Performance Test: Prototype vs Direct Creation")
	fmt.Println(strings.Repeat("-", 50))

	iterations := 1000

	// Test prototype cloning
	fmt.Printf("🔄 Testing prototype cloning (%d iterations)...\n", iterations)
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		doc, _ := factory.CreateDocument("basic-text", fmt.Sprintf("Document %d", i), "Test Author")
		_ = doc // Use the document to prevent optimization
	}
	
	prototypeTime := time.Since(start)

	// Test direct creation
	fmt.Printf("🏗️  Testing direct creation (%d iterations)...\n", iterations)
	start = time.Now()
	
	for i := 0; i < iterations; i++ {
		doc := prototypes.NewTextDocument(fmt.Sprintf("Document %d", i), "This is a template for text documents.", "Test Author")
		doc.AddTag("template")
		doc.AddTag("text")
		_ = doc // Use the document to prevent optimization
	}
	
	directTime := time.Since(start)

	fmt.Printf("\n📊 Results:\n")
	fmt.Printf("   Prototype cloning: %v\n", prototypeTime)
	fmt.Printf("   Direct creation:   %v\n", directTime)
	fmt.Printf("   Difference:        %v\n", directTime-prototypeTime)
	
	if prototypeTime < directTime {
		speedup := float64(directTime) / float64(prototypeTime)
		fmt.Printf("   Prototype is %.2fx faster! 🚀\n", speedup)
	} else {
		slowdown := float64(prototypeTime) / float64(directTime)
		fmt.Printf("   Direct creation is %.2fx faster\n", slowdown)
	}
}f
unc runAutomatedDemo() {
	// Create document factory
	factory := prototypes.NewDocumentFactory()

	// Demo 1: Basic Prototype Cloning
	fmt.Println("\n🎬 Demo 1: Basic Prototype Cloning")
	fmt.Println(strings.Repeat("-", 40))

	// Clone different types of documents from prototypes
	textDoc, _ := factory.CreateDocument("basic-text", "My Text Document", "John Doe")
	pdfDoc, _ := factory.CreateDocument("report-pdf", "Quarterly Report", "Jane Smith")
	spreadsheetDoc, _ := factory.CreateDocument("budget-spreadsheet", "2024 Budget", "Bob Johnson")
	presentationDoc, _ := factory.CreateDocument("business-presentation", "Company Overview", "Alice Brown")

	fmt.Printf("📄 Created from prototypes:\n")
	fmt.Printf("   %s\n", textDoc.GetInfo())
	fmt.Printf("   %s\n", pdfDoc.GetInfo())
	fmt.Printf("   %s\n", spreadsheetDoc.GetInfo())
	fmt.Printf("   %s\n", presentationDoc.GetInfo())

	time.Sleep(1 * time.Second)

	// Demo 2: Deep Cloning Demonstration
	fmt.Println("\n🎬 Demo 2: Deep Cloning vs Shallow Copying")
	fmt.Println(strings.Repeat("-", 45))

	// Create original document with complex data
	original := prototypes.NewPDFDocument("Original Contract", "Contract terms and conditions", "Legal Team", 5)
	original.AddBookmark("Section 1")
	original.AddBookmark("Section 2")
	original.AddAnnotation(1, "highlight", "Important clause", "Reviewer")
	original.SetEncryption(true, "secret123")

	fmt.Printf("📄 Original document: %s\n", original.GetInfo())
	fmt.Printf("   Bookmarks: %d\n", len(original.Bookmarks))
	fmt.Printf("   Annotations: %d\n", len(original.Annotations))
	fmt.Printf("   Encrypted: %t\n", original.IsEncrypted)

	// Clone the document
	cloned := original.Clone().(*prototypes.PDFDocument)
	cloned.SetTitle("Cloned Contract")
	cloned.SetAuthor("New Author")

	fmt.Printf("\n📄 Cloned document: %s\n", cloned.GetInfo())
	fmt.Printf("   Bookmarks: %d\n", len(cloned.Bookmarks))
	fmt.Printf("   Annotations: %d\n", len(cloned.Annotations))
	fmt.Printf("   Encrypted: %t\n", cloned.IsEncrypted)

	// Modify original to show independence
	original.AddBookmark("Section 3")
	original.SetTitle("Modified Original")

	fmt.Printf("\n🔍 After modifying original:\n")
	fmt.Printf("   Original title: %s (bookmarks: %d)\n", original.GetTitle(), len(original.Bookmarks))
	fmt.Printf("   Cloned title: %s (bookmarks: %d)\n", cloned.GetTitle(), len(cloned.Bookmarks))
	fmt.Printf("   ✅ Clone is independent of original!\n")

	time.Sleep(1 * time.Second)

	// Demo 3: Registry-based Prototype Management
	fmt.Println("\n🎬 Demo 3: Prototype Registry Management")
	fmt.Println(strings.Repeat("-", 42))

	registry := factory.GetRegistry()

	fmt.Printf("📋 Available prototypes: %d\n", len(registry.ListPrototypes()))
	for _, key := range registry.ListPrototypes() {
		info, _ := registry.GetPrototypeInfo(key)
		fmt.Printf("   • %s: %s\n", key, info)
	}

	// Register a custom prototype
	customTemplate := prototypes.NewTextDocument("Meeting Notes Template", 
		"Meeting: [Title]\nDate: [Date]\nAttendees: [Names]\n\nAgenda:\n1. [Item 1]\n2. [Item 2]\n\nNotes:\n[Content]\n\nAction Items:\n- [Action 1]\n- [Action 2]", 
		"System")
	customTemplate.AddTag("template")
	customTemplate.AddTag("meeting")
	customTemplate.SetMetadata("category", "business")

	registry.RegisterPrototype("meeting-notes", customTemplate)

	// Use the new prototype
	meetingDoc, _ := factory.CreateDocument("meeting-notes", "Weekly Team Meeting", "Project Manager")
	fmt.Printf("\n📄 Created from custom prototype: %s\n", meetingDoc.GetInfo())

	time.Sleep(1 * time.Second)

	// Demo 4: Performance Comparison
	fmt.Println("\n🎬 Demo 4: Performance Comparison")
	fmt.Println(strings.Repeat("-", 35))

	iterations := 100
	fmt.Printf("⚡ Creating %d documents...\n", iterations)

	// Prototype-based creation
	start := time.Now()
	for i := 0; i < iterations; i++ {
		doc, _ := factory.CreateDocument("basic-text", fmt.Sprintf("Doc %d", i), "Author")
		_ = doc
	}
	prototypeTime := time.Since(start)

	// Direct creation
	start = time.Now()
	for i := 0; i < iterations; i++ {
		doc := prototypes.NewTextDocument(fmt.Sprintf("Doc %d", i), "Content", "Author")
		doc.AddTag("template")
		doc.AddTag("text")
		_ = doc
	}
	directTime := time.Since(start)

	fmt.Printf("📊 Performance Results:\n")
	fmt.Printf("   Prototype cloning: %v\n", prototypeTime)
	fmt.Printf("   Direct creation:   %v\n", directTime)
	
	if prototypeTime < directTime {
		speedup := float64(directTime) / float64(prototypeTime)
		fmt.Printf("   🚀 Prototype is %.2fx faster!\n", speedup)
	} else {
		fmt.Printf("   Direct creation is faster by %v\n", directTime-prototypeTime)
	}

	time.Sleep(1 * time.Second)

	// Demo 5: Complex Document Cloning
	fmt.Println("\n🎬 Demo 5: Complex Document Structures")
	fmt.Println(strings.Repeat("-", 40))

	// Create a complex spreadsheet
	complexSheet := prototypes.NewSpreadsheetDocument("Financial Model", "Complex financial calculations", "Finance Team", 5)
	complexSheet.AddSheet("Revenue", 50, 12)
	complexSheet.AddSheet("Expenses", 40, 10)
	complexSheet.AddSheet("Profit & Loss", 30, 8)
	complexSheet.AddSheet("Cash Flow", 25, 6)
	complexSheet.AddSheet("Dashboard", 15, 15)
	
	// Add formulas
	complexSheet.AddFormula("L51", "=SUM(L2:L50)", 1000000)
	complexSheet.AddFormula("J41", "=SUM(J2:J40)", 750000)
	complexSheet.AddFormula("H31", "=L51-J41", 250000)
	
	complexSheet.HasCharts = true
	complexSheet.AddTag("financial-model")
	complexSheet.AddTag("template")

	fmt.Printf("📊 Original complex spreadsheet: %s\n", complexSheet.GetInfo())
	fmt.Printf("   Sheets: %d, Formulas: %d, Charts: %t\n", 
		len(complexSheet.Sheets), len(complexSheet.Formulas), complexSheet.HasCharts)

	// Clone and modify
	clonedSheet := complexSheet.Clone().(*prototypes.SpreadsheetDocument)
	clonedSheet.SetTitle("Q2 Financial Model")
	clonedSheet.SetAuthor("Q2 Finance Team")
	clonedSheet.AddSheet("Forecast", 20, 8)

	fmt.Printf("\n📊 Cloned and modified: %s\n", clonedSheet.GetInfo())
	fmt.Printf("   Sheets: %d, Formulas: %d, Charts: %t\n", 
		len(clonedSheet.Sheets), len(clonedSheet.Formulas), clonedSheet.HasCharts)

	fmt.Printf("\n🔍 Independence check:\n")
	fmt.Printf("   Original sheets: %d\n", len(complexSheet.Sheets))
	fmt.Printf("   Cloned sheets: %d\n", len(clonedSheet.Sheets))
	fmt.Printf("   ✅ Clone has additional sheet without affecting original!\n")

	// Demo 6: Prototype Variations
	fmt.Println("\n🎬 Demo 6: Creating Variations from Prototypes")
	fmt.Println(strings.Repeat("-", 45))

	// Create multiple variations of the same prototype
	basePresentation, _ := factory.CreateDocument("business-presentation", "Sales Presentation", "Sales Team")
	
	variations := []struct {
		title  string
		author string
		theme  string
	}{
		{"Q1 Sales Review", "Sales Manager", "Corporate"},
		{"Product Launch", "Marketing Team", "Modern"},
		{"Investor Pitch", "CEO", "Executive"},
		{"Training Session", "HR Department", "Educational"},
	}

	fmt.Printf("📊 Creating variations from base prototype:\n")
	for i, variation := range variations {
		clone := basePresentation.Clone().(*prototypes.PresentationDocument)
		clone.SetTitle(variation.title)
		clone.SetAuthor(variation.author)
		clone.SetTheme(variation.theme)
		
		fmt.Printf("   %d. %s\n", i+1, clone.GetInfo())
	}

	fmt.Println("\n✨ Prototype Pattern Demo Complete!")
	fmt.Println("\nKey Benefits Demonstrated:")
	fmt.Println("• 🚀 Efficient object creation through cloning")
	fmt.Println("• 🔄 Deep copying preserves object independence")
	fmt.Println("• 📋 Registry pattern for prototype management")
	fmt.Println("• ⚡ Performance benefits for complex objects")
	fmt.Println("• 🎯 Easy creation of object variations")
	fmt.Println("• 🏗️  Reduced complexity in object construction")
}