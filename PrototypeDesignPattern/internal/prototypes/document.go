package prototypes

import (
	"fmt"
	"time"
)

// Prototype interface defines the contract for cloneable objects
type Prototype interface {
	Clone() Prototype
	GetInfo() string
	GetType() string
}

// DocumentPrototype is the base interface for all document types
type DocumentPrototype interface {
	Prototype
	SetTitle(title string)
	SetContent(content string)
	SetAuthor(author string)
	GetTitle() string
	GetContent() string
	GetAuthor() string
	GetCreatedAt() time.Time
	GetSize() int
}

// BaseDocument provides common functionality for all document types
type BaseDocument struct {
	ID        string
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
	Tags      []string
	Metadata  map[string]interface{}
}

// NewBaseDocument creates a new base document
func NewBaseDocument(title, content, author string) *BaseDocument {
	return &BaseDocument{
		ID:        generateID(),
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
		Tags:      make([]string, 0),
		Metadata:  make(map[string]interface{}),
	}
}

// Common methods for BaseDocument
func (bd *BaseDocument) SetTitle(title string) {
	bd.Title = title
	bd.UpdatedAt = time.Now()
	bd.Version++
}

func (bd *BaseDocument) SetContent(content string) {
	bd.Content = content
	bd.UpdatedAt = time.Now()
	bd.Version++
}

func (bd *BaseDocument) SetAuthor(author string) {
	bd.Author = author
	bd.UpdatedAt = time.Now()
}

func (bd *BaseDocument) GetTitle() string {
	return bd.Title
}

func (bd *BaseDocument) GetContent() string {
	return bd.Content
}

func (bd *BaseDocument) GetAuthor() string {
	return bd.Author
}

func (bd *BaseDocument) GetCreatedAt() time.Time {
	return bd.CreatedAt
}

func (bd *BaseDocument) GetSize() int {
	return len(bd.Content)
}

func (bd *BaseDocument) AddTag(tag string) {
	bd.Tags = append(bd.Tags, tag)
}

func (bd *BaseDocument) SetMetadata(key string, value interface{}) {
	bd.Metadata[key] = value
}

// Helper function to generate unique IDs
func generateID() string {
	return fmt.Sprintf("doc_%d", time.Now().UnixNano())
}// 
TextDocument represents a simple text document
type TextDocument struct {
	*BaseDocument
	Encoding    string
	LineEndings string
	WordCount   int
}

// NewTextDocument creates a new text document
func NewTextDocument(title, content, author string) *TextDocument {
	return &TextDocument{
		BaseDocument: NewBaseDocument(title, content, author),
		Encoding:     "UTF-8",
		LineEndings:  "LF",
		WordCount:    countWords(content),
	}
}

// Clone creates a deep copy of the text document
func (td *TextDocument) Clone() Prototype {
	// Create new BaseDocument with copied values
	clonedBase := &BaseDocument{
		ID:        generateID(), // New ID for the clone
		Title:     td.Title,
		Content:   td.Content,
		Author:    td.Author,
		CreatedAt: time.Now(), // New creation time for clone
		UpdatedAt: time.Now(),
		Version:   1, // Reset version for clone
		Tags:      make([]string, len(td.Tags)),
		Metadata:  make(map[string]interface{}),
	}
	
	// Deep copy tags
	copy(clonedBase.Tags, td.Tags)
	
	// Deep copy metadata
	for k, v := range td.Metadata {
		clonedBase.Metadata[k] = v
	}
	
	// Create cloned TextDocument
	cloned := &TextDocument{
		BaseDocument: clonedBase,
		Encoding:     td.Encoding,
		LineEndings:  td.LineEndings,
		WordCount:    td.WordCount,
	}
	
	return cloned
}

func (td *TextDocument) GetInfo() string {
	return fmt.Sprintf("Text Document: '%s' by %s (%d words, %s encoding)",
		td.Title, td.Author, td.WordCount, td.Encoding)
}

func (td *TextDocument) GetType() string {
	return "TextDocument"
}

func (td *TextDocument) SetContent(content string) {
	td.BaseDocument.SetContent(content)
	td.WordCount = countWords(content)
}

// PDFDocument represents a PDF document with additional properties
type PDFDocument struct {
	*BaseDocument
	PageCount    int
	IsEncrypted  bool
	Password     string
	Compression  string
	Bookmarks    []string
	Annotations  []Annotation
}

type Annotation struct {
	Page    int
	Type    string
	Content string
	Author  string
}

// NewPDFDocument creates a new PDF document
func NewPDFDocument(title, content, author string, pageCount int) *PDFDocument {
	return &PDFDocument{
		BaseDocument: NewBaseDocument(title, content, author),
		PageCount:    pageCount,
		IsEncrypted:  false,
		Compression:  "ZIP",
		Bookmarks:    make([]string, 0),
		Annotations:  make([]Annotation, 0),
	}
}

// Clone creates a deep copy of the PDF document
func (pd *PDFDocument) Clone() Prototype {
	// Create new BaseDocument with copied values
	clonedBase := &BaseDocument{
		ID:        generateID(),
		Title:     pd.Title,
		Content:   pd.Content,
		Author:    pd.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
		Tags:      make([]string, len(pd.Tags)),
		Metadata:  make(map[string]interface{}),
	}
	
	// Deep copy tags
	copy(clonedBase.Tags, pd.Tags)
	
	// Deep copy metadata
	for k, v := range pd.Metadata {
		clonedBase.Metadata[k] = v
	}
	
	// Deep copy bookmarks
	clonedBookmarks := make([]string, len(pd.Bookmarks))
	copy(clonedBookmarks, pd.Bookmarks)
	
	// Deep copy annotations
	clonedAnnotations := make([]Annotation, len(pd.Annotations))
	copy(clonedAnnotations, pd.Annotations)
	
	// Create cloned PDFDocument
	cloned := &PDFDocument{
		BaseDocument: clonedBase,
		PageCount:    pd.PageCount,
		IsEncrypted:  pd.IsEncrypted,
		Password:     pd.Password,
		Compression:  pd.Compression,
		Bookmarks:    clonedBookmarks,
		Annotations:  clonedAnnotations,
	}
	
	return cloned
}

func (pd *PDFDocument) GetInfo() string {
	encryptedStatus := ""
	if pd.IsEncrypted {
		encryptedStatus = " (Encrypted)"
	}
	return fmt.Sprintf("PDF Document: '%s' by %s (%d pages, %s compression)%s",
		pd.Title, pd.Author, pd.PageCount, pd.Compression, encryptedStatus)
}

func (pd *PDFDocument) GetType() string {
	return "PDFDocument"
}

func (pd *PDFDocument) AddBookmark(bookmark string) {
	pd.Bookmarks = append(pd.Bookmarks, bookmark)
}

func (pd *PDFDocument) AddAnnotation(page int, annotationType, content, author string) {
	annotation := Annotation{
		Page:    page,
		Type:    annotationType,
		Content: content,
		Author:  author,
	}
	pd.Annotations = append(pd.Annotations, annotation)
}

func (pd *PDFDocument) SetEncryption(encrypted bool, password string) {
	pd.IsEncrypted = encrypted
	pd.Password = password
}

// SpreadsheetDocument represents a spreadsheet document
type SpreadsheetDocument struct {
	*BaseDocument
	SheetCount   int
	CellCount    int
	HasFormulas  bool
	HasCharts    bool
	Sheets       []Sheet
	Formulas     []Formula
}

type Sheet struct {
	Name     string
	RowCount int
	ColCount int
	IsHidden bool
}

type Formula struct {
	Cell    string
	Formula string
	Result  interface{}
}

// NewSpreadsheetDocument creates a new spreadsheet document
func NewSpreadsheetDocument(title, content, author string, sheetCount int) *SpreadsheetDocument {
	return &SpreadsheetDocument{
		BaseDocument: NewBaseDocument(title, content, author),
		SheetCount:   sheetCount,
		CellCount:    0,
		HasFormulas:  false,
		HasCharts:    false,
		Sheets:       make([]Sheet, 0),
		Formulas:     make([]Formula, 0),
	}
}

// Clone creates a deep copy of the spreadsheet document
func (sd *SpreadsheetDocument) Clone() Prototype {
	// Create new BaseDocument with copied values
	clonedBase := &BaseDocument{
		ID:        generateID(),
		Title:     sd.Title,
		Content:   sd.Content,
		Author:    sd.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
		Tags:      make([]string, len(sd.Tags)),
		Metadata:  make(map[string]interface{}),
	}
	
	// Deep copy tags
	copy(clonedBase.Tags, sd.Tags)
	
	// Deep copy metadata
	for k, v := range sd.Metadata {
		clonedBase.Metadata[k] = v
	}
	
	// Deep copy sheets
	clonedSheets := make([]Sheet, len(sd.Sheets))
	copy(clonedSheets, sd.Sheets)
	
	// Deep copy formulas
	clonedFormulas := make([]Formula, len(sd.Formulas))
	copy(clonedFormulas, sd.Formulas)
	
	// Create cloned SpreadsheetDocument
	cloned := &SpreadsheetDocument{
		BaseDocument: clonedBase,
		SheetCount:   sd.SheetCount,
		CellCount:    sd.CellCount,
		HasFormulas:  sd.HasFormulas,
		HasCharts:    sd.HasCharts,
		Sheets:       clonedSheets,
		Formulas:     clonedFormulas,
	}
	
	return cloned
}

func (sd *SpreadsheetDocument) GetInfo() string {
	features := ""
	if sd.HasFormulas {
		features += " with formulas"
	}
	if sd.HasCharts {
		features += " with charts"
	}
	return fmt.Sprintf("Spreadsheet Document: '%s' by %s (%d sheets, %d cells)%s",
		sd.Title, sd.Author, sd.SheetCount, sd.CellCount, features)
}

func (sd *SpreadsheetDocument) GetType() string {
	return "SpreadsheetDocument"
}

func (sd *SpreadsheetDocument) AddSheet(name string, rows, cols int) {
	sheet := Sheet{
		Name:     name,
		RowCount: rows,
		ColCount: cols,
		IsHidden: false,
	}
	sd.Sheets = append(sd.Sheets, sheet)
	sd.CellCount += rows * cols
}

func (sd *SpreadsheetDocument) AddFormula(cell, formula string, result interface{}) {
	f := Formula{
		Cell:    cell,
		Formula: formula,
		Result:  result,
	}
	sd.Formulas = append(sd.Formulas, f)
	sd.HasFormulas = true
}

// PresentationDocument represents a presentation document
type PresentationDocument struct {
	*BaseDocument
	SlideCount    int
	HasAnimations bool
	HasTransitions bool
	Theme         string
	Slides        []Slide
	MasterSlides  []MasterSlide
}

type Slide struct {
	Number      int
	Title       string
	Content     string
	Layout      string
	HasImages   bool
	HasVideos   bool
	Transition  string
}

type MasterSlide struct {
	Name   string
	Layout string
	Theme  string
}

// NewPresentationDocument creates a new presentation document
func NewPresentationDocument(title, content, author string, slideCount int) *PresentationDocument {
	return &PresentationDocument{
		BaseDocument:   NewBaseDocument(title, content, author),
		SlideCount:     slideCount,
		HasAnimations:  false,
		HasTransitions: false,
		Theme:          "Default",
		Slides:         make([]Slide, 0),
		MasterSlides:   make([]MasterSlide, 0),
	}
}

// Clone creates a deep copy of the presentation document
func (pd *PresentationDocument) Clone() Prototype {
	// Create new BaseDocument with copied values
	clonedBase := &BaseDocument{
		ID:        generateID(),
		Title:     pd.Title,
		Content:   pd.Content,
		Author:    pd.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
		Tags:      make([]string, len(pd.Tags)),
		Metadata:  make(map[string]interface{}),
	}
	
	// Deep copy tags
	copy(clonedBase.Tags, pd.Tags)
	
	// Deep copy metadata
	for k, v := range pd.Metadata {
		clonedBase.Metadata[k] = v
	}
	
	// Deep copy slides
	clonedSlides := make([]Slide, len(pd.Slides))
	copy(clonedSlides, pd.Slides)
	
	// Deep copy master slides
	clonedMasterSlides := make([]MasterSlide, len(pd.MasterSlides))
	copy(clonedMasterSlides, pd.MasterSlides)
	
	// Create cloned PresentationDocument
	cloned := &PresentationDocument{
		BaseDocument:   clonedBase,
		SlideCount:     pd.SlideCount,
		HasAnimations:  pd.HasAnimations,
		HasTransitions: pd.HasTransitions,
		Theme:          pd.Theme,
		Slides:         clonedSlides,
		MasterSlides:   clonedMasterSlides,
	}
	
	return cloned
}

func (pd *PresentationDocument) GetInfo() string {
	features := ""
	if pd.HasAnimations {
		features += " with animations"
	}
	if pd.HasTransitions {
		features += " with transitions"
	}
	return fmt.Sprintf("Presentation Document: '%s' by %s (%d slides, %s theme)%s",
		pd.Title, pd.Author, pd.SlideCount, pd.Theme, features)
}

func (pd *PresentationDocument) GetType() string {
	return "PresentationDocument"
}

func (pd *PresentationDocument) AddSlide(title, content, layout string) {
	slide := Slide{
		Number:     len(pd.Slides) + 1,
		Title:      title,
		Content:    content,
		Layout:     layout,
		HasImages:  false,
		HasVideos:  false,
		Transition: "None",
	}
	pd.Slides = append(pd.Slides, slide)
}

func (pd *PresentationDocument) SetTheme(theme string) {
	pd.Theme = theme
}

// Helper function to count words in content
func countWords(content string) int {
	if content == "" {
		return 0
	}
	words := 0
	inWord := false
	for _, char := range content {
		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			inWord = false
		} else if !inWord {
			words++
			inWord = true
		}
	}
	return words
}//
 DocumentRegistry manages prototype instances for cloning
type DocumentRegistry struct {
	prototypes map[string]DocumentPrototype
}

// NewDocumentRegistry creates a new document registry
func NewDocumentRegistry() *DocumentRegistry {
	return &DocumentRegistry{
		prototypes: make(map[string]DocumentPrototype),
	}
}

// RegisterPrototype adds a prototype to the registry
func (dr *DocumentRegistry) RegisterPrototype(key string, prototype DocumentPrototype) {
	dr.prototypes[key] = prototype
	fmt.Printf("📋 Registered prototype: %s (%s)\n", key, prototype.GetType())
}

// GetPrototype retrieves and clones a prototype from the registry
func (dr *DocumentRegistry) GetPrototype(key string) (DocumentPrototype, error) {
	prototype, exists := dr.prototypes[key]
	if !exists {
		return nil, fmt.Errorf("prototype '%s' not found in registry", key)
	}
	
	cloned := prototype.Clone()
	if doc, ok := cloned.(DocumentPrototype); ok {
		return doc, nil
	}
	
	return nil, fmt.Errorf("cloned object is not a DocumentPrototype")
}

// ListPrototypes returns all registered prototype keys
func (dr *DocumentRegistry) ListPrototypes() []string {
	keys := make([]string, 0, len(dr.prototypes))
	for key := range dr.prototypes {
		keys = append(keys, key)
	}
	return keys
}

// GetPrototypeInfo returns information about a registered prototype
func (dr *DocumentRegistry) GetPrototypeInfo(key string) (string, error) {
	prototype, exists := dr.prototypes[key]
	if !exists {
		return "", fmt.Errorf("prototype '%s' not found", key)
	}
	return prototype.GetInfo(), nil
}

// RemovePrototype removes a prototype from the registry
func (dr *DocumentRegistry) RemovePrototype(key string) error {
	if _, exists := dr.prototypes[key]; !exists {
		return fmt.Errorf("prototype '%s' not found", key)
	}
	delete(dr.prototypes, key)
	fmt.Printf("🗑️  Removed prototype: %s\n", key)
	return nil
}

// DocumentFactory provides convenient methods for creating documents
type DocumentFactory struct {
	registry *DocumentRegistry
}

// NewDocumentFactory creates a new document factory
func NewDocumentFactory() *DocumentFactory {
	factory := &DocumentFactory{
		registry: NewDocumentRegistry(),
	}
	
	// Register default prototypes
	factory.registerDefaultPrototypes()
	
	return factory
}

// registerDefaultPrototypes sets up common document prototypes
func (df *DocumentFactory) registerDefaultPrototypes() {
	// Text document templates
	basicText := NewTextDocument("Template Text Document", "This is a template for text documents.", "System")
	basicText.AddTag("template")
	basicText.AddTag("text")
	df.registry.RegisterPrototype("basic-text", basicText)
	
	letterTemplate := NewTextDocument("Letter Template", "Dear [Recipient],\n\n[Content]\n\nSincerely,\n[Sender]", "System")
	letterTemplate.AddTag("template")
	letterTemplate.AddTag("letter")
	letterTemplate.SetMetadata("type", "business-letter")
	df.registry.RegisterPrototype("letter-template", letterTemplate)
	
	// PDF document templates
	reportPDF := NewPDFDocument("Report Template", "Executive Summary\n\nIntroduction\n\nFindings\n\nConclusion", "System", 10)
	reportPDF.AddTag("template")
	reportPDF.AddTag("report")
	reportPDF.AddBookmark("Executive Summary")
	reportPDF.AddBookmark("Introduction")
	reportPDF.AddBookmark("Findings")
	reportPDF.AddBookmark("Conclusion")
	df.registry.RegisterPrototype("report-pdf", reportPDF)
	
	contractPDF := NewPDFDocument("Contract Template", "AGREEMENT\n\nParties: [Party A] and [Party B]\n\nTerms and Conditions\n\nSignatures", "System", 5)
	contractPDF.SetEncryption(true, "")
	contractPDF.AddTag("template")
	contractPDF.AddTag("contract")
	contractPDF.AddTag("legal")
	df.registry.RegisterPrototype("contract-pdf", contractPDF)
	
	// Spreadsheet templates
	budgetSheet := NewSpreadsheetDocument("Budget Template", "Monthly Budget Spreadsheet", "System", 3)
	budgetSheet.AddSheet("Income", 20, 5)
	budgetSheet.AddSheet("Expenses", 30, 5)
	budgetSheet.AddSheet("Summary", 10, 3)
	budgetSheet.AddFormula("C21", "=SUM(C2:C20)", 0)
	budgetSheet.AddTag("template")
	budgetSheet.AddTag("budget")
	budgetSheet.AddTag("finance")
	df.registry.RegisterPrototype("budget-spreadsheet", budgetSheet)
	
	inventorySheet := NewSpreadsheetDocument("Inventory Template", "Product Inventory Tracking", "System", 2)
	inventorySheet.AddSheet("Products", 100, 8)
	inventorySheet.AddSheet("Reports", 20, 6)
	inventorySheet.AddFormula("H101", "=SUM(H2:H100)", 0)
	inventorySheet.AddTag("template")
	inventorySheet.AddTag("inventory")
	df.registry.RegisterPrototype("inventory-spreadsheet", inventorySheet)
	
	// Presentation templates
	businessPresentation := NewPresentationDocument("Business Presentation Template", "Company Overview Presentation", "System", 8)
	businessPresentation.SetTheme("Corporate")
	businessPresentation.AddSlide("Title Slide", "Company Name\nPresentation Title\nDate", "title")
	businessPresentation.AddSlide("Agenda", "1. Overview\n2. Products\n3. Market\n4. Financials", "bullet")
	businessPresentation.AddSlide("Company Overview", "About our company...", "content")
	businessPresentation.AddTag("template")
	businessPresentation.AddTag("business")
	df.registry.RegisterPrototype("business-presentation", businessPresentation)
	
	trainingPresentation := NewPresentationDocument("Training Template", "Employee Training Presentation", "System", 12)
	trainingPresentation.SetTheme("Educational")
	trainingPresentation.HasAnimations = true
	trainingPresentation.HasTransitions = true
	trainingPresentation.AddTag("template")
	trainingPresentation.AddTag("training")
	df.registry.RegisterPrototype("training-presentation", trainingPresentation)
}

// CreateDocument creates a new document from a registered prototype
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

// GetRegistry returns the document registry for direct access
func (df *DocumentFactory) GetRegistry() *DocumentRegistry {
	return df.registry
}

// CreateCustomDocument creates a document of a specific type
func (df *DocumentFactory) CreateCustomDocument(docType, title, content, author string) (DocumentPrototype, error) {
	switch docType {
	case "text":
		return NewTextDocument(title, content, author), nil
	case "pdf":
		return NewPDFDocument(title, content, author, 1), nil
	case "spreadsheet":
		return NewSpreadsheetDocument(title, content, author, 1), nil
	case "presentation":
		return NewPresentationDocument(title, content, author, 1), nil
	default:
		return nil, fmt.Errorf("unknown document type: %s", docType)
	}
}