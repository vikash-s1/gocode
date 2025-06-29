// Package legacy contains old payment systems with incompatible interfaces
package legacy

import (
	"fmt"
	"time"
)

// PayPalGateway represents the legacy PayPal payment system
// This has its own unique interface that doesn't match our modern standard
type PayPalGateway struct {
	apiKey    string
	isActive  bool
	lastError error
}

// MakePayment is PayPal's specific method for processing payments
func (p *PayPalGateway) MakePayment(dollars float64, currency string) (string, error) {
	if dollars <= 0 {
		p.lastError = fmt.Errorf("invalid amount: %.2f", dollars)
		return "", p.lastError
	}
	
	// Simulate PayPal processing
	transactionID := fmt.Sprintf("PP_%d", time.Now().Unix())
	fmt.Printf("   ✓ PayPal: Processing $%.2f %s (Transaction: %s)\n", dollars, currency, transactionID)
	
	p.isActive = true
	p.lastError = nil
	return transactionID, nil
}

// GetPayPalStatus returns PayPal-specific status
func (p *PayPalGateway) GetPayPalStatus() string {
	if p.lastError != nil {
		return "ERROR: " + p.lastError.Error()
	}
	if p.isActive {
		return "PAYPAL_ACTIVE"
	}
	return "PAYPAL_INACTIVE"
}

// StripeGateway represents the legacy Stripe payment system
type StripeGateway struct {
	secretKey string
	status    string
}

// ChargeCard is Stripe's specific method for processing payments
func (s *StripeGateway) ChargeCard(amountInCents int, cardToken string) (*StripeCharge, error) {
	if amountInCents <= 0 {
		return nil, fmt.Errorf("invalid amount in cents: %d", amountInCents)
	}
	
	// Simulate Stripe processing
	charge := &StripeCharge{
		ID:       fmt.Sprintf("ch_%d", time.Now().Unix()),
		Amount:   amountInCents,
		Currency: "usd",
		Status:   "succeeded",
	}
	
	fmt.Printf("   ✓ Stripe: Charged %d cents (Charge: %s)\n", amountInCents, charge.ID)
	s.status = "STRIPE_SUCCESS"
	return charge, nil
}

// GetStripeStatus returns Stripe-specific status
func (s *StripeGateway) GetStripeStatus() string {
	if s.status == "" {
		return "STRIPE_READY"
	}
	return s.status
}

// StripeCharge represents a Stripe charge object
type StripeCharge struct {
	ID       string
	Amount   int
	Currency string
	Status   string
}

// BitcoinGateway represents a legacy Bitcoin payment system
type BitcoinGateway struct {
	walletAddress string
	networkFee    float64
	confirmations int
}

// SendBitcoin is Bitcoin's specific method for processing payments
func (b *BitcoinGateway) SendBitcoin(btcAmount float64, toAddress string) (string, error) {
	if btcAmount <= 0 {
		return "", fmt.Errorf("invalid BTC amount: %.8f", btcAmount)
	}
	
	// Simulate Bitcoin transaction
	txHash := fmt.Sprintf("btc_%x", time.Now().Unix())
	fmt.Printf("   ✓ Bitcoin: Sent %.8f BTC (TX: %s)\n", btcAmount, txHash)
	
	b.confirmations = 0
	return txHash, nil
}

// GetBitcoinNetworkStatus returns Bitcoin-specific status
func (b *BitcoinGateway) GetBitcoinNetworkStatus() string {
	return fmt.Sprintf("BTC_NETWORK_ACTIVE (Confirmations: %d)", b.confirmations)
}

// GetNetworkFee returns the current network fee
func (b *BitcoinGateway) GetNetworkFee() float64 {
	return b.networkFee
}