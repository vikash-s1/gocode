package strategies

import (
	"fmt"
	"math/rand"
	"time"
)

// PaymentStrategy defines the interface for all payment strategies
type PaymentStrategy interface {
	Pay(amount float64) error
	GetPaymentMethod() string
	ValidatePayment(amount float64) error
}

// PaymentContext holds the current payment strategy and manages payments
type PaymentContext struct {
	strategy PaymentStrategy
	customer string
}

// NewPaymentContext creates a new payment context
func NewPaymentContext(customer string) *PaymentContext {
	return &PaymentContext{
		customer: customer,
	}
}

// SetStrategy changes the payment strategy at runtime
func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
	fmt.Printf("🔄 Payment method changed to: %s\n", strategy.GetPaymentMethod())
}

// ProcessPayment executes payment using the current strategy
func (pc *PaymentContext) ProcessPayment(amount float64) error {
	if pc.strategy == nil {
		return fmt.Errorf("no payment strategy selected")
	}

	fmt.Printf("\n💳 Processing payment for %s\n", pc.customer)
	fmt.Printf("💰 Amount: $%.2f\n", amount)
	fmt.Printf("📱 Method: %s\n", pc.strategy.GetPaymentMethod())

	// Validate payment first
	if err := pc.strategy.ValidatePayment(amount); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Process payment
	return pc.strategy.Pay(amount)
}

// GetCurrentMethod returns the current payment method name
func (pc *PaymentContext) GetCurrentMethod() string {
	if pc.strategy == nil {
		return "None"
	}
	return pc.strategy.GetPaymentMethod()
}
//
 CreditCardStrategy implements payment via credit card
type CreditCardStrategy struct {
	CardNumber string
	CVV        string
	ExpiryDate string
	CardHolder string
}

// NewCreditCardStrategy creates a new credit card payment strategy
func NewCreditCardStrategy(cardNumber, cvv, expiryDate, cardHolder string) *CreditCardStrategy {
	return &CreditCardStrategy{
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpiryDate: expiryDate,
		CardHolder: cardHolder,
	}
}

func (cc *CreditCardStrategy) ValidatePayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: $%.2f", amount)
	}
	if len(cc.CardNumber) != 16 {
		return fmt.Errorf("invalid card number length")
	}
	if len(cc.CVV) != 3 {
		return fmt.Errorf("invalid CVV")
	}
	if cc.CardHolder == "" {
		return fmt.Errorf("card holder name required")
	}
	return nil
}

func (cc *CreditCardStrategy) Pay(amount float64) error {
	fmt.Printf("🔒 Validating credit card: ****-****-****-%s\n", cc.CardNumber[12:])
	time.Sleep(1 * time.Second) // Simulate processing time
	
	// Simulate random payment failure (10% chance)
	if rand.Float32() < 0.1 {
		return fmt.Errorf("credit card payment declined")
	}
	
	fmt.Printf("✅ Credit card payment successful: $%.2f charged to %s\n", amount, cc.CardHolder)
	return nil
}

func (cc *CreditCardStrategy) GetPaymentMethod() string {
	return "Credit Card"
}

// PayPalStrategy implements payment via PayPal
type PayPalStrategy struct {
	Email    string
	Password string
}

// NewPayPalStrategy creates a new PayPal payment strategy
func NewPayPalStrategy(email, password string) *PayPalStrategy {
	return &PayPalStrategy{
		Email:    email,
		Password: password,
	}
}

func (pp *PayPalStrategy) ValidatePayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: $%.2f", amount)
	}
	if pp.Email == "" {
		return fmt.Errorf("PayPal email required")
	}
	if pp.Password == "" {
		return fmt.Errorf("PayPal password required")
	}
	return nil
}

func (pp *PayPalStrategy) Pay(amount float64) error {
	fmt.Printf("🔐 Authenticating PayPal account: %s\n", pp.Email)
	time.Sleep(800 * time.Millisecond) // Simulate processing time
	
	// Simulate random payment failure (5% chance)
	if rand.Float32() < 0.05 {
		return fmt.Errorf("PayPal payment failed - insufficient funds")
	}
	
	fmt.Printf("✅ PayPal payment successful: $%.2f from %s\n", amount, pp.Email)
	return nil
}

func (pp *PayPalStrategy) GetPaymentMethod() string {
	return "PayPal"
}

// CryptoStrategy implements payment via cryptocurrency
type CryptoStrategy struct {
	WalletAddress string
	CryptoType    string
	PrivateKey    string
}

// NewCryptoStrategy creates a new cryptocurrency payment strategy
func NewCryptoStrategy(walletAddress, cryptoType, privateKey string) *CryptoStrategy {
	return &CryptoStrategy{
		WalletAddress: walletAddress,
		CryptoType:    cryptoType,
		PrivateKey:    privateKey,
	}
}

func (cs *CryptoStrategy) ValidatePayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: $%.2f", amount)
	}
	if cs.WalletAddress == "" {
		return fmt.Errorf("wallet address required")
	}
	if cs.CryptoType == "" {
		return fmt.Errorf("cryptocurrency type required")
	}
	if len(cs.WalletAddress) < 26 {
		return fmt.Errorf("invalid wallet address format")
	}
	return nil
}

func (cs *CryptoStrategy) Pay(amount float64) error {
	fmt.Printf("⛓️  Connecting to %s blockchain...\n", cs.CryptoType)
	time.Sleep(1500 * time.Millisecond) // Simulate blockchain processing time
	
	// Simulate random payment failure (15% chance - crypto can be volatile)
	if rand.Float32() < 0.15 {
		return fmt.Errorf("cryptocurrency transaction failed - network congestion")
	}
	
	fmt.Printf("✅ %s payment successful: $%.2f from wallet %s...%s\n", 
		cs.CryptoType, amount, cs.WalletAddress[:6], cs.WalletAddress[len(cs.WalletAddress)-4:])
	return nil
}

func (cs *CryptoStrategy) GetPaymentMethod() string {
	return fmt.Sprintf("Cryptocurrency (%s)", cs.CryptoType)
}

// BankTransferStrategy implements payment via bank transfer
type BankTransferStrategy struct {
	AccountNumber string
	RoutingNumber string
	BankName      string
	AccountHolder string
}

// NewBankTransferStrategy creates a new bank transfer payment strategy
func NewBankTransferStrategy(accountNumber, routingNumber, bankName, accountHolder string) *BankTransferStrategy {
	return &BankTransferStrategy{
		AccountNumber: accountNumber,
		RoutingNumber: routingNumber,
		BankName:      bankName,
		AccountHolder: accountHolder,
	}
}

func (bt *BankTransferStrategy) ValidatePayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: $%.2f", amount)
	}
	if len(bt.AccountNumber) < 8 {
		return fmt.Errorf("invalid account number")
	}
	if len(bt.RoutingNumber) != 9 {
		return fmt.Errorf("invalid routing number")
	}
	if bt.AccountHolder == "" {
		return fmt.Errorf("account holder name required")
	}
	return nil
}

func (bt *BankTransferStrategy) Pay(amount float64) error {
	fmt.Printf("🏦 Initiating bank transfer from %s\n", bt.BankName)
	fmt.Printf("📋 Account: ****%s\n", bt.AccountNumber[len(bt.AccountNumber)-4:])
	time.Sleep(2 * time.Second) // Simulate bank processing time
	
	// Simulate random payment failure (8% chance)
	if rand.Float32() < 0.08 {
		return fmt.Errorf("bank transfer failed - insufficient funds or account locked")
	}
	
	fmt.Printf("✅ Bank transfer successful: $%.2f from %s's account\n", amount, bt.AccountHolder)
	return nil
}

func (bt *BankTransferStrategy) GetPaymentMethod() string {
	return "Bank Transfer"
}