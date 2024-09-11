package main

// State interface
type State interface {
	insertCash(vm *VendingMachine, amount float64)
	selectProduct(vm *VendingMachine, productId int)
	cancelTransaction(vm *VendingMachine)
	dispenseProduct(vm *VendingMachine)
}
