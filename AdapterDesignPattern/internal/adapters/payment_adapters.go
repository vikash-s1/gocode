// Package adapters contains adapter implementations that bridge legacy systems to modern interfaces
package adapters

import (
	"fmt"

	"github.com/example/adapter-pattern/internal/legacy"
	"github.com/example/adapter-pattern/internal/payment"
)

// PayPalAdapter adapts the legacy PayPal gateway to the modern payment interface
type PayPalAdapter struct {
	gateway *legacy.PayPalGateway
}

// NewPayPalAdapter creates a new PayPal adapter
func NewPayPalAdapter(gateway *legacy.PayPalGateway) *PayPalAdapter {
	return &PayPalAdapter{gateway: gateway}
}

// ProcessPayment adapts PayPal's MakePayment method to our standard interface
func (p *PayPalAdapter) ProcessPayment(amount float64) error {
	// Convert our standard call to PayPal's specific method
	_, err := p.gateway.MakePayment(amount, "USD")
	return err
}

// GetStatus adapts PayPal's status method to our standard interface
func (p *PayPalAdapter) GetStatus() string {
	return p.gateway.GetPayPalStatus()
}

// GetSupportedCurrencies returns PayPal's supported currencies
func (p *PayPalAdapter) GetSupportedCurrencies() []string {
	return []string{"USD", "EUR", "GBP", "CAD", "AUD"}
}

// StripeAdapter adapts the legacy Stripe gateway to the modern payment interface
type StripeAdapter struct {
	gateway *legacy.StripeGateway
}

// NewStripeAdapter creates a new Stripe adapter
func NewStripeAdapter(gateway *legacy.StripeGateway) *StripeAdapter {
	return &StripeAdapter{gateway: gateway}
}

// ProcessPayment adapts Stripe's ChargeCard method to our standard interface
func (s *StripeAdapter) ProcessPayment(amount float64) error {
	// Convert dollars to cents (Stripe's requirement)
	amountInCents := int(amount * 100)
	
	// Convert our standard call to Stripe's specific method
	_, err := s.gateway.ChargeCard(amountInCents, "card_token_placeholder")
	return err
}

// GetStatus adapts Stripe's status method to our standard interface
func (s *StripeAdapter) GetStatus() string {
	return s.gateway.GetStripeStatus()
}

// GetSupportedCurrencies returns Stripe's supported currencies
func (s *StripeAdapter) GetSupportedCurrencies() []string {
	return []string{"USD", "EUR", "GBP", "JPY", "CAD", "AUD", "CHF", "SEK"}
}

// BitcoinAdapter adapts the legacy Bitcoin gateway to the modern payment interface
type BitcoinAdapter struct {
	gateway   *legacy.BitcoinGateway
	btcToUSD  float64 // Exchange rate for conversion
}

// NewBitcoinAdapter creates a new Bitcoin adapter
func NewBitcoinAdapter(gateway *legacy.BitcoinGateway) *BitcoinAdapter {
	return &BitcoinAdapter{
		gateway:  gateway,
		btcToUSD: 45000.0, // Simulated exchange rate
	}
}

// ProcessPayment adapts Bitcoin's SendBitcoin method to our standard interface
func (b *BitcoinAdapter) ProcessPayment(amount float64) error {
	// Convert USD amount to BTC
	btcAmount := amount / b.btcToUSD
	
	// Convert our standard call to Bitcoin's specific method
	_, err := b.gateway.SendBitcoin(btcAmount, "destination_address_placeholder")
	return err
}

// GetStatus adapts Bitcoin's status method to our standard interface
func (b *BitcoinAdapter) GetStatus() string {
	return b.gateway.GetBitcoinNetworkStatus()
}

// GetSupportedCurrencies returns Bitcoin's supported currencies
func (b *BitcoinAdapter) GetSupportedCurrencies() []string {
	return []string{"BTC", "USD"} // Bitcoin primarily, with USD conversion
}

// PaymentAdapterFactory creates payment adapters based on payment type
type PaymentAdapterFactory struct{}

// CreateAdapter creates the appropriate payment adapter based on payment type
func (f *PaymentAdapterFactory) CreateAdapter(paymentType string) (payment.Processor, error) {
	switch paymentType {
	case "paypal":
		return NewPayPalAdapter(&legacy.PayPalGateway{}), nil
	case "stripe":
		return NewStripeAdapter(&legacy.StripeGateway{}), nil
	case "bitcoin":
		return NewBitcoinAdapter(&legacy.BitcoinGateway{}), nil
	default:
		return nil, fmt.Errorf("unsupported payment type: %s", paymentType)
	}
}