// Package payment defines the modern payment processing interface
package payment

// Processor defines the standard interface for payment processing
// This is our target interface that all payment methods should implement
type Processor interface {
	// ProcessPayment processes a payment for the given amount
	ProcessPayment(amount float64) error
	
	// GetStatus returns the current status of the payment processor
	GetStatus() string
	
	// GetSupportedCurrencies returns list of supported currencies
	GetSupportedCurrencies() []string
}

// PaymentResult represents the result of a payment operation
type PaymentResult struct {
	TransactionID string
	Amount        float64
	Currency      string
	Status        string
	Message       string
}

// PaymentError represents payment-specific errors
type PaymentError struct {
	Code    string
	Message string
	Details map[string]interface{}
}

func (e *PaymentError) Error() string {
	return e.Message
}