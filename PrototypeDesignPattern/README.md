# Prototype Design Pattern in Go

## Overview

The Prototype Design Pattern is a creational design pattern that allows you to create new objects by cloning existing instances (prototypes) rather than creating them from scratch. This pattern is particularly useful when object creation is expensive or complex, and you need to create many similar objects with slight variations.

## Problem It Solves

Without the Prototype pattern, creating complex objects can be expensive and cumbersome:

```go
// BAD: Expensive object creation every time
func CreateComplexDocument() *Document {
    doc := &Document{}
    
    // Expensive initialization
    doc.LoadTemplate()           // File I/O operation
    doc.InitializeFormatting()   // Complex setup
    doc.SetupDefaultStyles()     // Resource-intensive
    doc.ConfigureMetadata()      // Database queries
    doc.LoadPlugins()           // External dependencies
    
    return doc
}

// Creating 100 similar documents = 100x expensive operations!
for i := 0; i < 100; i++ {
    doc := CreateComplexDocument() // Repeats all expensive operations
    doc.SetTitle(fmt.Sprintf("Document %d", i))
}
```

This leads to:
- **Performance Issues**: Repeated expensive initialization
- **Resource Waste**: Unnecessary I/O, database calls, memory allocation
- **Complex Construction**: Intricate object setup logic scattered throughout code
- **Inflexibility**: Hard to create variations of similar objects

## Solution

The Prototype pattern creates objects by cloning pre-configured prototypes, avoiding expensive initialization:

```go
// GOOD: Clone from prototype
prototype := CreateComplexDocumentPrototype() // Expensive operation done once

// Creating 100 similar documents = 1 expensive operation + 100 cheap clones
for i := 0; i < 100; i++ {
    doc := prototype.Clone()                    // Fast cloning operation
    doc.SetTitle(fmt.Sprintf("Document %d", i)) // Customize as needed
}
```

## Implementation Structure

### Core Components

1. **Prototype Interface** (`Prototype`)
   - Defines the `Clone()` method that all prototypes must implement
   - Ensures consistent cloning behavior across different object types

2. **Concrete Prototypes** (Document Types)
   - `TextDocument`: Simple text documents with encoding and word count
   - `PDFDocument`: PDF documents with pages, bookmarks, and annotations
   - `SpreadsheetDocument`: Spreadsheets with sheets, formulas, and charts
   - `PresentationDocument`: Presentations with slides, themes, and animations

3. **Prototype Registry** (`DocumentRegistry`)
   - Manages a collection of prototype instances
   - Provides methods to register, retrieve, and clone prototypes
   - Acts as a central repository for reusable prototypes

4. **Client/Factory** (`DocumentFactory`)
   - Uses prototypes to create new objects
   - Provides convenient methods for document creation
   - Manages default prototypes and custom templates

## Document Management Example

Our implementation simulates a document management system with different document types:

```
┌─────────────────┐
│    Prototype    │
│   (Interface)   │
│                 │
│ + Clone()       │
│ + GetInfo()     │
│ + GetType()     │
└─────────────────┘
         ▲
    ┌────┼────┐
    │    │    │
┌──────────┐ ┌─────────────┐ ┌──────────────────┐
│TextDoc   │ │PDFDocument  │ │SpreadsheetDoc    │
│          │ │             │ │                  │
│+ Clone() │ │+ Clone()    │ │+ Clone()         │
└──────────┘ └─────────────┘ └──────────────────┘

┌─────────────────────┐
│ DocumentRegistry    │
│                     │
│ - prototypes map    │
│ + RegisterPrototype()│
│ + GetPrototype()    │
│ + ListPrototypes()  │
└─────────────────────┘
         ▲
         │ uses
┌─────────────────────┐
│ DocumentFactory     │
│                     │
│ - registry          │
│ + CreateDocument()  │
│ + CreateCustom()    │
└─────────────────────┘
```

## Step-by-Step Implementation

### Step 1: Define the Prototype Interface

```go
type Prototype interface {
    Clone() Prototype
    GetInfo() string
    GetType() string
}

type DocumentPrototype interface {
    Prototype
    SetTitle(title string)
    SetContent(content string)
    SetAuthor(author string)
    // ... other document-specific methods
}
```

**Purpose**: Establishes the contract for cloneable objects, ensuring all prototypes can be cloned consistently.

### Step 2: Implement Concrete Prototypes

```go
type TextDocument struct {
    *BaseDocument
    Encoding    string
    LineEndings string
    WordCount   int
}

func (td *TextDocument) Clone() Prototype {
    // Create new BaseDocument with copied values
    clonedBase := &BaseDocument{
        ID:        generateID(), // New ID for clone
        Title:     td.Title,
        Content:   td.Content,
        Author:    td.Author,
        CreatedAt: time.Now(),   // New creation time
        // ... copy other fields
    }
    
    // Deep copy complex fields
    clonedBase.Tags = make([]string, len(td.Tags))
    copy(clonedBase.Tags, td.Tags)
    
    // Create cloned TextDocument
    return &TextDocument{
        BaseDocument: clonedBase,
        Encoding:     td.Encoding,
        LineEndings:  td.LineEndings,
        WordCount:    td.WordCount,
    }
}
```

**Key Features**:
- **Deep Copying**: Creates independent copies of complex fields (slices, maps)
- **New Identity**: Generates new ID and creation timestamp for clones
- **Preserved State**: Maintains all configuration and content from prototype

### Step 3: Create Prototype Registry

```go
type DocumentRegistry struct {
    prototypes map[string]DocumentPrototype
}

func (dr *DocumentRegistry) RegisterPrototype(key string, prototype DocumentPrototype) {
    dr.prototypes[key] = prototype
}

func (dr *DocumentRegistry) GetPrototype(key string) (DocumentPrototype, error) {
    prototype, exists := dr.prototypes[key]
    if !exists {
        return nil, fmt.Errorf("prototype '%s' not found", key)
    }
    
    cloned := prototype.Clone()
    return cloned.(DocumentPrototype), nil
}
```

**Benefits**:
- **Centralized Management**: Single location for all prototypes
- **Easy Access**: Retrieve prototypes by name/key
- **Automatic Cloning**: Returns cloned instances, not originals

### Step 4: Implement Document Factory

```go
type DocumentFactory struct {
    registry *DocumentRegistry
}

func (df *DocumentFactory) CreateDocument(prototypeKey, title, author string) (DocumentPrototype, error) {
    document, err := df.registry.GetPrototype(prototypeKey)
    if err != nil {
        return nil, err
    }
    
    // Customize the cloned document
    document.SetTitle(title)
    document.SetAuthor(author)
    
    return document, nil
}
```

**Features**:
- **Convenient Interface**: Simple methods for document creation
- **Automatic Customization**: Sets common properties after cloning
- **Error Handling**: Proper error management for missing prototypes

## Key Benefits Demonstrated

### 1. **Performance Optimization**
```go
// Expensive prototype creation (done once)
complexPrototype := CreateComplexPrototype() // 100ms

// Fast cloning (done many times)
for i := 0; i < 1000; i++ {
    clone := complexPrototype.Clone() // 1ms each
    // Total: 100ms + (1000 × 1ms) = 1.1 seconds
    // vs Direct creation: 1000 × 100ms = 100 seconds
}
```

### 2. **Deep vs Shallow Copying**
```go
original := NewPDFDocument("Contract", "Terms...", "Legal", 5)
original.AddBookmark("Section 1")
original.AddAnnotation(1, "highlight", "Important", "Reviewer")

clone := original.Clone().(*PDFDocument)
clone.SetTitle("Modified Contract")

// Modify original
original.AddBookmark("Section 2")

// Clone remains independent
fmt.Printf("Original bookmarks: %d\n", len(original.Bookmarks)) // 2
fmt.Printf("Clone bookmarks: %d\n", len(clone.Bookmarks))       // 1
```

### 3. **Registry-based Management**
```go
// Register prototypes once
registry.RegisterPrototype("contract-template", contractPrototype)
registry.RegisterPrototype("report-template", reportPrototype)

// Use anywhere in application
contract, _ := registry.GetPrototype("contract-template")
report, _ := registry.GetPrototype("report-template")
```

### 4. **Easy Variations**
```go
basePresentation, _ := factory.CreateDocument("business-template", "Base", "Author")

// Create variations
salesPitch := basePresentation.Clone()
salesPitch.SetTitle("Sales Pitch")
salesPitch.SetTheme("Corporate")

trainingSession := basePresentation.Clone()
trainingSession.SetTitle("Training Session")
trainingSession.SetTheme("Educational")
```

## Running the Example

### Build and Run
```bash
cd PrototypeDesignPattern
go run main.go
```

### Interactive Mode Features
1. **List Available Prototypes**: View all registered document templates
2. **Create from Prototype**: Clone existing prototypes with custom titles/authors
3. **Create Custom Document**: Build new documents from scratch
4. **Clone Existing Document**: Clone documents created during the session
5. **Register New Prototype**: Add custom prototypes to the registry
6. **Compare Documents**: Side-by-side comparison of document properties
7. **Performance Test**: Benchmark prototype cloning vs direct creation

### Automated Demos
The program demonstrates:
1. **Basic Prototype Cloning**: Creating different document types from prototypes
2. **Deep Cloning**: Independence between original and cloned objects
3. **Registry Management**: Centralized prototype storage and retrieval
4. **Performance Comparison**: Speed benefits of prototype pattern
5. **Complex Document Structures**: Cloning documents with nested data
6. **Prototype Variations**: Creating multiple variations from single prototype

## Document Types Comparison

| Document Type | Complexity | Special Features | Clone Performance |
|---------------|------------|------------------|-------------------|
| TextDocument | Low | Encoding, word count | Very Fast |
| PDFDocument | Medium | Pages, bookmarks, annotations, encryption | Fast |
| SpreadsheetDocument | High | Multiple sheets, formulas, charts | Medium |
| PresentationDocument | High | Slides, themes, animations, transitions | Medium |

## When to Use Prototype Pattern

✅ **Use When**:
- Object creation is expensive (I/O, database, network operations)
- You need many similar objects with slight variations
- Object initialization is complex with many steps
- You want to avoid subclassing for object creation
- Runtime object composition is needed
- You need to create objects without knowing their exact types

❌ **Avoid When**:
- Object creation is simple and fast
- Objects don't have complex initialization
- You rarely need multiple similar objects
- Deep copying is expensive (large object graphs)
- Objects have circular references (complicates cloning)

## Real-World Applications

### Game Development
```go
// Character prototypes for different classes
type CharacterPrototype interface {
    Clone() CharacterPrototype
    SetName(name string)
    SetLevel(level int)
}

// Warrior, Mage, Archer prototypes
warriorPrototype := NewWarrior("Template Warrior", defaultWarriorStats)
registry.RegisterPrototype("warrior", warriorPrototype)

// Create player characters
player1 := registry.GetPrototype("warrior")
player1.SetName("Conan")
player1.SetLevel(10)
```

### UI Component Libraries
```go
// Widget prototypes with default styling
type WidgetPrototype interface {
    Clone() WidgetPrototype
    SetPosition(x, y int)
    SetSize(width, height int)
}

// Button, TextField, Panel prototypes
buttonPrototype := NewButton("Default Button", defaultButtonStyle)
registry.RegisterPrototype("primary-button", buttonPrototype)

// Create UI elements
loginButton := registry.GetPrototype("primary-button")
loginButton.SetText("Login")
loginButton.SetPosition(100, 200)
```

### Configuration Management
```go
// Server configuration prototypes
type ServerConfigPrototype interface {
    Clone() ServerConfigPrototype
    SetEnvironment(env string)
    SetPort(port int)
}

// Development, staging, production prototypes
devConfig := NewServerConfig("development", defaultDevSettings)
registry.RegisterPrototype("dev-server", devConfig)

// Create environment-specific configs
webServer := registry.GetPrototype("dev-server")
webServer.SetPort(8080)
webServer.SetEnvironment("development")
```

### Document Templates
```go
// Business document prototypes
invoiceTemplate := NewInvoiceDocument(defaultInvoiceFormat)
contractTemplate := NewContractDocument(defaultContractTerms)
reportTemplate := NewReportDocument(defaultReportStructure)

// Create specific documents
customerInvoice := invoiceTemplate.Clone()
customerInvoice.SetCustomer("ABC Corp")
customerInvoice.SetAmount(5000.00)
```

## Comparison with Other Patterns

### vs Factory Method Pattern
- **Prototype**: Creates objects by cloning existing instances
- **Factory Method**: Creates objects using construction logic in subclasses

### vs Builder Pattern
- **Prototype**: Clones pre-configured objects, then customizes
- **Builder**: Constructs objects step-by-step with fluent interface

### vs Singleton Pattern
- **Prototype**: Creates multiple instances through cloning
- **Singleton**: Ensures only one instance exists

### vs Abstract Factory Pattern
- **Prototype**: Uses cloning to create object families
- **Abstract Factory**: Uses factory methods to create related objects

## Advanced Considerations

### Thread Safety
For concurrent environments:
```go
import "sync"

type ThreadSafeRegistry struct {
    mu         sync.RWMutex
    prototypes map[string]DocumentPrototype
}

func (tsr *ThreadSafeRegistry) GetPrototype(key string) (DocumentPrototype, error) {
    tsr.mu.RLock()
    prototype, exists := tsr.prototypes[key]
    tsr.mu.RUnlock()
    
    if !exists {
        return nil, fmt.Errorf("prototype not found")
    }
    
    return prototype.Clone().(DocumentPrototype), nil
}
```

### Circular References
Handle circular references in cloning:
```go
type CloneContext struct {
    cloned map[interface{}]interface{}
}

func (obj *ComplexObject) CloneWithContext(ctx *CloneContext) Prototype {
    if cloned, exists := ctx.cloned[obj]; exists {
        return cloned.(Prototype)
    }
    
    clone := &ComplexObject{}
    ctx.cloned[obj] = clone
    
    // Clone fields, using context to handle circular references
    clone.Reference = obj.Reference.CloneWithContext(ctx)
    
    return clone
}
```

### Lazy Cloning
Optimize cloning with copy-on-write:
```go
type LazyClone struct {
    original *Document
    modified map[string]interface{}
    cloned   bool
}

func (lc *LazyClone) SetTitle(title string) {
    if !lc.cloned {
        lc.performClone()
    }
    lc.modified["title"] = title
}

func (lc *LazyClone) performClone() {
    // Perform actual cloning only when modification occurs
    lc.cloned = true
    // ... cloning logic
}
```

### Prototype Versioning
Manage prototype versions:
```go
type VersionedPrototype struct {
    prototype DocumentPrototype
    version   string
    createdAt time.Time
}

type VersionedRegistry struct {
    prototypes map[string]map[string]*VersionedPrototype
}

func (vr *VersionedRegistry) RegisterPrototype(key, version string, prototype DocumentPrototype) {
    if vr.prototypes[key] == nil {
        vr.prototypes[key] = make(map[string]*VersionedPrototype)
    }
    
    vr.prototypes[key][version] = &VersionedPrototype{
        prototype: prototype,
        version:   version,
        createdAt: time.Now(),
    }
}
```

### Memory Management
Prevent memory leaks in prototype registries:
```go
type WeakPrototypeRegistry struct {
    prototypes map[string]func() DocumentPrototype // Factory functions instead of instances
}

func (wpr *WeakPrototypeRegistry) RegisterPrototype(key string, factory func() DocumentPrototype) {
    wpr.prototypes[key] = factory
}

func (wpr *WeakPrototypeRegistry) GetPrototype(key string) (DocumentPrototype, error) {
    factory, exists := wpr.prototypes[key]
    if !exists {
        return nil, fmt.Errorf("prototype not found")
    }
    
    return factory(), nil // Create fresh instance each time
}
```

This implementation demonstrates how the Prototype pattern creates efficient, flexible object creation systems by leveraging cloning instead of construction, providing significant performance benefits for complex objects while maintaining clean separation of concerns and easy extensibility.