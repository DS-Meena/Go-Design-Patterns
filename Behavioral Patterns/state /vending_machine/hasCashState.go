package main

import "fmt"

// Hash Cash state struct
type HasCashState struct{}

func (s *HasCashState) insertCash(vm *VendingMachine, amount float64) {
	vm.currentBalance += amount
	fmt.Printf("Insert %.2f. Current balance: %.2f\n", amount, vm.currentBalance)
}

func (s *HasCashState) selectProduct(vm *VendingMachine, productId int) {
	for _, product := range vm.products {
		if product.id == productId {
			if vm.currentBalance >= product.price {
				fmt.Printf("Selected product: %s \n", product.name)
				vm.setState(&ProductSelectedState{product: product})
				return
			} else {
				fmt.Printf("Insufficient balance. Please insert %.2f more \n", product.price-vm.currentBalance)
				return
			}
		}
	}

	fmt.Println("Invalid product selection")
}

func (s *HasCashState) cancelTransaction(vm *VendingMachine) {
	fmt.Printf("Returning %.2f\n", vm.currentBalance)
	vm.currentBalance = 0
	vm.setState(&NoCashState{})
}
func (s *HasCashState) dispenseProduct(vm *VendingMachine) {
	fmt.Println("Pelase select a product first")
}
