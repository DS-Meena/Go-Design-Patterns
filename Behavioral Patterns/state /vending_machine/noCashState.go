package main

import "fmt"

// No cash state struct
type NoCashState struct{}

func (s *NoCashState) insertCash(vm *VendingMachine, amount float64) {
	vm.currentBalance += amount
	fmt.Printf("Insert %.2f. Current balance: %.2f\n", amount, vm.currentBalance)
	vm.setState(&HasCashState{})
}

func (s *NoCashState) selectProduct(vm *VendingMachine, productId int) {
	fmt.Println("Please insert cash first")
}

func (s *NoCashState) cancelTransaction(vm *VendingMachine) {
	fmt.Println("No transaction to cancel")
}

func (s *NoCashState) dispenseProduct(vm *VendingMachine) {
	fmt.Println("Please insert cash and select a product first")
}
