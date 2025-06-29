package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"strategypattern/internal/strategies"
)

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("💳 Welcome to the Strategy Pattern Payment System Demo!")
	fmt.Println(strings.Repeat("=", 55))

	// Interactive demo
	runInteractiveDemo()

	fmt.Println("\n🎯 Automated Demo:")
	fmt.Println(strings.Repeat("=", 30))

	// Automated demo to show all strategies
	runAutomatedDemo()
}

func runInteractiveDemo() {
	scanner := bufio.NewScanner(os.Stdin)

	// Create payment context
	fmt.Print("Enter customer name: ")
	scanner.Scan()
	customerName := strings.TrimSpace(scanner.Text())
	if customerName == "" {
		customerName = "John Doe"
	}

	paymentContext := strategies.NewPaymentContext(customerName)

	fmt.Printf("\n👋 Welcome %s! Choose your payment method:\n", customerName)
	fmt.Println("1 - Credit Card")
	fmt.Println("2 - PayPal")
	fmt.Println("3 - Cryptocurrency")
	fmt.Println("4 - Bank Transfer")
	fmt.Println("5 - Switch Payment Method")
	fmt.Println("6 - Process Payment")
	fmt.Println("q - Quit to Automated Demo")

	var amount float64 = 99.99 // Default amount

	for {
		fmt.Printf("\nCurrent Method: %s | Amount: $%.2f\n", paymentContext.GetCurrentMethod(), amount)
		fmt.Print("Enter choice: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			setupCreditCard(paymentContext, scanner)
		case "2":
			setupPayPal(paymentContext, scanner)
		case "3":
			setupCrypto(paymentContext, scanner)
		case "4":
			setupBankTransfer(paymentContext, scanner)
		case "5":
			switchPaymentMethod(paymentContext, scanner)
		case "6":
			fmt.Print("Enter amount to pay: $")
			scanner.Scan()
			if amountStr := strings.TrimSpace(scanner.Text()); amountStr != "" {
				if parsedAmount, err := strconv.ParseFloat(amountStr, 64); err == nil {
					amount = parsedAmount
				}
			}
			processPayment(paymentContext, amount)
		case "q":
			return
		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}

func setupCreditCard(ctx *strategies.PaymentContext, scanner *bufio.Scanner) {
	fmt.Println("\n💳 Setting up Credit Card payment...")
	
	// Use default values for demo
	strategy := strategies.NewCreditCardStrategy(
		"1234567890123456", // Card number
		"123",              // CVV
		"12/25",            // Expiry
		"John Doe",         // Cardholder
	)
	
	ctx.SetStrategy(strategy)
	fmt.Println("✅ Credit Card configured with demo data")
}

func setupPayPal(ctx *strategies.PaymentContext, scanner *bufio.Scanner) {
	fmt.Println("\n💙 Setting up PayPal payment...")
	
	strategy := strategies.NewPayPalStrategy(
		"john.doe@email.com", // Email
		"securepassword",     // Password
	)
	
	ctx.SetStrategy(strategy)
	fmt.Println("✅ PayPal configured with demo data")
}

func setupCrypto(ctx *strategies.PaymentContext, scanner *bufio.Scanner) {
	fmt.Println("\n₿ Setting up Cryptocurrency payment...")
	
	strategy := strategies.NewCryptoStrategy(
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", // Bitcoin wallet
		"Bitcoin",                              // Crypto type
		"private_key_demo",                     // Private key
	)
	
	ctx.SetStrategy(strategy)
	fmt.Println("✅ Cryptocurrency configured with demo data")
}

func setupBankTransfer(ctx *strategies.PaymentContext, scanner *bufio.Scanner) {
	fmt.Println("\n🏦 Setting up Bank Transfer payment...")
	
	strategy := strategies.NewBankTransferStrategy(
		"123456789",    // Account number
		"021000021",    // Routing number
		"Demo Bank",    // Bank name
		"John Doe",     // Account holder
	)
	
	ctx.SetStrategy(strategy)
	fmt.Println("✅ Bank Transfer configured with demo data")
}

func switchPaymentMethod(ctx *strategies.PaymentContext, scanner *bufio.Scanner) {
	fmt.Println("\n🔄 Available payment methods:")
	fmt.Println("1 - Credit Card")
	fmt.Println("2 - PayPal") 
	fmt.Println("3 - Cryptocurrency")
	fmt.Println("4 - Bank Transfer")
	
	fmt.Print("Choose new method: ")
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())
	
	switch choice {
	case "1":
		setupCreditCard(ctx, scanner)
	case "2":
		setupPayPal(ctx, scanner)
	case "3":
		setupCrypto(ctx, scanner)
	case "4":
		setupBankTransfer(ctx, scanner)
	default:
		fmt.Println("❌ Invalid choice")
	}
}

func processPayment(ctx *strategies.PaymentContext, amount float64) {
	fmt.Println(strings.Repeat("-", 40))
	
	if err := ctx.ProcessPayment(amount); err != nil {
		fmt.Printf("❌ Payment failed: %v\n", err)
	} else {
		fmt.Println("🎉 Payment completed successfully!")
	}
	
	fmt.Println(strings.Repeat("-", 40))
}

func runAutomatedDemo() {
	customer := "Alice Smith"
	paymentContext := strategies.NewPaymentContext(customer)
	
	// Demo different payment strategies
	strategies := []struct {
		name     string
		strategy strategies.PaymentStrategy
		amount   float64
	}{
		{
			"Credit Card Payment",
			strategies.NewCreditCardStrategy("4532123456789012", "456", "03/26", "Alice Smith"),
			149.99,
		},
		{
			"PayPal Payment", 
			strategies.NewPayPalStrategy("alice.smith@email.com", "mypassword123"),
			75.50,
		},
		{
			"Cryptocurrency Payment",
			strategies.NewCryptoStrategy("bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", "Bitcoin", "demo_private_key"),
			299.99,
		},
		{
			"Bank Transfer Payment",
			strategies.NewBankTransferStrategy("987654321", "111000025", "First National Bank", "Alice Smith"),
			500.00,
		},
	}
	
	for i, demo := range strategies {
		fmt.Printf("\n🎬 Demo %d: %s\n", i+1, demo.name)
		fmt.Println(strings.Repeat("-", 35))
		
		// Set strategy and process payment
		paymentContext.SetStrategy(demo.strategy)
		
		if err := paymentContext.ProcessPayment(demo.amount); err != nil {
			fmt.Printf("❌ Payment failed: %v\n", err)
		} else {
			fmt.Println("🎉 Payment completed successfully!")
		}
		
		// Add delay between demos
		time.Sleep(500 * time.Millisecond)
	}
	
	// Demonstrate strategy switching at runtime
	fmt.Println("\n🔄 Runtime Strategy Switching Demo:")
	fmt.Println(strings.Repeat("-", 40))
	
	// Start with credit card
	paymentContext.SetStrategy(strategies[0].strategy)
	fmt.Printf("Initial method: %s\n", paymentContext.GetCurrentMethod())
	
	// Switch to PayPal
	paymentContext.SetStrategy(strategies[1].strategy)
	fmt.Printf("Switched to: %s\n", paymentContext.GetCurrentMethod())
	
	// Switch to Crypto
	paymentContext.SetStrategy(strategies[2].strategy)
	fmt.Printf("Switched to: %s\n", paymentContext.GetCurrentMethod())
	
	// Process final payment
	fmt.Println("\n💰 Processing final payment with current strategy:")
	paymentContext.ProcessPayment(199.99)
	
	fmt.Println("\n✨ Strategy Pattern Demo Complete!")
}