package main

// Vending Machine struct
type VendingMachine struct {
	currentState   State
	products       []Product
	currentBalance float64
}

// Vending machine methods
func (vm *VendingMachine) setState(state State) {
	vm.currentState = state
}

func (vm *VendingMachine) insertCash(amount float64) {
	vm.currentState.insertCash(vm, amount)
}

func (vm *VendingMachine) selectProduct(productId int) {
	vm.currentState.selectProduct(vm, productId)
}

func (vm *VendingMachine) cancelTransaction() {
	vm.currentState.cancelTransaction(vm)
}

func (vm *VendingMachine) dispenseProduct() {
	vm.currentState.dispenseProduct(vm)
}
