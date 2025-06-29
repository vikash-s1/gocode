package states

// State interface defines the contract for all concrete states
type State interface {
	InsertCoin(machine *VendingMachine)
	SelectProduct(machine *VendingMachine)
	DispenseProduct(machine *VendingMachine)
	Cancel(machine *VendingMachine)
	GetStateName() string
}

// VendingMachine represents the context that maintains state
type VendingMachine struct {
	currentState State
	coinInserted bool
	productCount int
}

// NewVendingMachine creates a new vending machine instance
func NewVendingMachine(productCount int) *VendingMachine {
	vm := &VendingMachine{
		productCount: productCount,
		coinInserted: false,
	}
	// Set initial state to IdleState
	vm.currentState = &IdleState{}
	return vm
}

// State transition methods
func (vm *VendingMachine) SetState(state State) {
	vm.currentState = state
}

func (vm *VendingMachine) GetCurrentState() string {
	return vm.currentState.GetStateName()
}

// Context methods that delegate to current state
func (vm *VendingMachine) InsertCoin() {
	vm.currentState.InsertCoin(vm)
}

func (vm *VendingMachine) SelectProduct() {
	vm.currentState.SelectProduct(vm)
}

func (vm *VendingMachine) DispenseProduct() {
	vm.currentState.DispenseProduct(vm)
}

func (vm *VendingMachine) Cancel() {
	vm.currentState.Cancel(vm)
}

// Helper methods
func (vm *VendingMachine) HasProducts() bool {
	return vm.productCount > 0
}

func (vm *VendingMachine) DecrementProductCount() {
	if vm.productCount > 0 {
		vm.productCount--
	}
}

func (vm *VendingMachine) GetProductCount() int {
	return vm.productCount
}

func (vm *VendingMachine) SetCoinInserted(inserted bool) {
	vm.coinInserted = inserted
}

func (vm *VendingMachine) IsCoinInserted() bool {
	return vm.coinInserted
}
// Idl
eState - Initial state when machine is waiting for coin
type IdleState struct{}

func (s *IdleState) InsertCoin(machine *VendingMachine) {
	println("💰 Coin inserted! Please select a product.")
	machine.SetCoinInserted(true)
	machine.SetState(&CoinInsertedState{})
}

func (s *IdleState) SelectProduct(machine *VendingMachine) {
	println("❌ Please insert a coin first!")
}

func (s *IdleState) DispenseProduct(machine *VendingMachine) {
	println("❌ Please insert a coin and select a product first!")
}

func (s *IdleState) Cancel(machine *VendingMachine) {
	println("ℹ️  Nothing to cancel. Machine is idle.")
}

func (s *IdleState) GetStateName() string {
	return "Idle"
}

// CoinInsertedState - State when coin is inserted but product not selected
type CoinInsertedState struct{}

func (s *CoinInsertedState) InsertCoin(machine *VendingMachine) {
	println("ℹ️  Coin already inserted. Please select a product.")
}

func (s *CoinInsertedState) SelectProduct(machine *VendingMachine) {
	if !machine.HasProducts() {
		println("❌ Sorry, no products available. Returning coin...")
		machine.SetCoinInserted(false)
		machine.SetState(&IdleState{})
		return
	}
	println("🎯 Product selected! Dispensing...")
	machine.SetState(&ProductSelectedState{})
}

func (s *CoinInsertedState) DispenseProduct(machine *VendingMachine) {
	println("❌ Please select a product first!")
}

func (s *CoinInsertedState) Cancel(machine *VendingMachine) {
	println("🔄 Transaction cancelled. Returning coin...")
	machine.SetCoinInserted(false)
	machine.SetState(&IdleState{})
}

func (s *CoinInsertedState) GetStateName() string {
	return "CoinInserted"
}

// ProductSelectedState - State when product is selected and being dispensed
type ProductSelectedState struct{}

func (s *ProductSelectedState) InsertCoin(machine *VendingMachine) {
	println("ℹ️  Product is being dispensed. Please wait...")
}

func (s *ProductSelectedState) SelectProduct(machine *VendingMachine) {
	println("ℹ️  Product already selected and being dispensed...")
}

func (s *ProductSelectedState) DispenseProduct(machine *VendingMachine) {
	println("📦 Product dispensed! Thank you for your purchase!")
	machine.DecrementProductCount()
	machine.SetCoinInserted(false)
	
	if machine.HasProducts() {
		machine.SetState(&IdleState{})
	} else {
		println("⚠️  Machine is now out of products!")
		machine.SetState(&OutOfStockState{})
	}
}

func (s *ProductSelectedState) Cancel(machine *VendingMachine) {
	println("❌ Cannot cancel. Product is being dispensed...")
}

func (s *ProductSelectedState) GetStateName() string {
	return "ProductSelected"
}

// OutOfStockState - State when machine has no products
type OutOfStockState struct{}

func (s *OutOfStockState) InsertCoin(machine *VendingMachine) {
	println("❌ Sorry, machine is out of stock. Cannot accept coins.")
}

func (s *OutOfStockState) SelectProduct(machine *VendingMachine) {
	println("❌ Sorry, machine is out of stock.")
}

func (s *OutOfStockState) DispenseProduct(machine *VendingMachine) {
	println("❌ Sorry, machine is out of stock.")
}

func (s *OutOfStockState) Cancel(machine *VendingMachine) {
	println("ℹ️  Nothing to cancel. Machine is out of stock.")
}

func (s *OutOfStockState) GetStateName() string {
	return "OutOfStock"
}