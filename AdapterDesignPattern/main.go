package main

import (
	"fmt"
	"log"

	"github.com/example/adapter-pattern/internal/adapters"
	"github.com/example/adapter-pattern/internal/legacy"
	"github.com/example/adapter-pattern/internal/modern"
	"github.com/example/adapter-pattern/internal/payment"
)

func main() {
	fmt.Println("=== Adapter Design Pattern Demo ===\n")

	// Scenario 1: Payment Processing
	demonstratePaymentAdapters()

	fmt.Println("\n" + "="*50 + "\n")

	// Scenario 2: Database Connection
	demonstrateDatabaseAdapters()

	fmt.Println("\n" + "="*50 + "\n")

	// Scenario 3: Media Player
	demonstrateMediaAdapters()
}

// demonstratePaymentAdapters shows how to adapt different payment systems
func demonstratePaymentAdapters() {
	fmt.Println("🏦 Payment Processing Adapters")
	fmt.Println("------------------------------")

	// Create different payment processors
	paypalGateway := &legacy.PayPalGateway{}
	stripeGateway := &legacy.StripeGateway{}
	bitcoinGateway := &legacy.BitcoinGateway{}

	// Adapt them to our unified interface
	paypalAdapter := adapters.NewPayPalAdapter(paypalGateway)
	stripeAdapter := adapters.NewStripeAdapter(stripeGateway)
	bitcoinAdapter := adapters.NewBitcoinAdapter(bitcoinGateway)

	// Use them through the same interface
	processors := []payment.Processor{
		paypalAdapter,
		stripeAdapter,
		bitcoinAdapter,
	}

	amount := 99.99
	for i, processor := range processors {
		fmt.Printf("\n%d. Processing $%.2f payment:\n", i+1, amount)
		
		if err := processor.ProcessPayment(amount); err != nil {
			log.Printf("Payment failed: %v", err)
		}
		
		fmt.Printf("   Status: %s\n", processor.GetStatus())
	}
}

// demonstrateDatabaseAdapters shows how to adapt different database systems
func demonstrateDatabaseAdapters() {
	fmt.Println("🗄️  Database Connection Adapters")
	fmt.Println("--------------------------------")

	// Legacy database systems
	mysqlDB := &legacy.MySQLDatabase{Host: "mysql.example.com", Port: 3306}
	postgresDB := &legacy.PostgreSQLDatabase{ConnectionString: "postgres://user:pass@localhost/db"}
	mongoClient := &legacy.MongoClient{URI: "mongodb://localhost:27017"}

	// Adapt them to modern interface
	mysqlAdapter := adapters.NewMySQLAdapter(mysqlDB)
	postgresAdapter := adapters.NewPostgreSQLAdapter(postgresDB)
	mongoAdapter := adapters.NewMongoAdapter(mongoClient)

	// Use them through unified interface
	databases := []modern.Database{
		mysqlAdapter,
		postgresAdapter,
		mongoAdapter,
	}

	for i, db := range databases {
		fmt.Printf("\n%d. Testing database connection:\n", i+1)
		
		if err := db.Connect(); err != nil {
			log.Printf("Connection failed: %v", err)
			continue
		}
		
		// Perform operations
		if err := db.ExecuteQuery("SELECT * FROM users LIMIT 5"); err != nil {
			log.Printf("Query failed: %v", err)
		}
		
		if err := db.Disconnect(); err != nil {
			log.Printf("Disconnect failed: %v", err)
		}
	}
}

// demonstrateMediaAdapters shows how to adapt different media formats
func demonstrateMediaAdapters() {
	fmt.Println("🎵 Media Player Adapters")
	fmt.Println("------------------------")

	// Legacy media players
	mp3Player := &legacy.MP3Player{}
	wavPlayer := &legacy.WAVPlayer{}
	flacPlayer := &legacy.FLACPlayer{}

	// Adapt them to modern interface
	mp3Adapter := adapters.NewMP3Adapter(mp3Player)
	wavAdapter := adapters.NewWAVAdapter(wavPlayer)
	flacAdapter := adapters.NewFLACAdapter(flacPlayer)

	// Use them through unified interface
	players := []modern.MediaPlayer{
		mp3Adapter,
		wavAdapter,
		flacAdapter,
	}

	songs := []string{
		"song1.mp3",
		"song2.wav",
		"song3.flac",
	}

	for i, player := range players {
		fmt.Printf("\n%d. Playing %s:\n", i+1, songs[i])
		
		if err := player.Play(songs[i]); err != nil {
			log.Printf("Playback failed: %v", err)
			continue
		}
		
		player.SetVolume(75)
		fmt.Printf("   Volume set to: %d%%\n", player.GetVolume())
		
		if err := player.Stop(); err != nil {
			log.Printf("Stop failed: %v", err)
		}
	}
}