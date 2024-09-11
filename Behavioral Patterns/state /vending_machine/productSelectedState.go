package main

import "fmt"

// ProductSelectedState struct
type ProductSelectedState struct {
	product Product
}

func (s *ProductSelectedState) insertCash(vm *VendingMachine, amount float64) {
	fmt.Println("Product already selected. Please wait for dispensing")
}

func (s *ProductSelectedState) selectProduct(vm *VendingMachine, productId int) {
	fmt.Println("Product already selected. Please wait for dispensing")
}

func (s *ProductSelectedState) cancelTransaction(vm *VendingMachine) {
	fmt.Printf("Cancelling transaction. Returning %.2f \n", vm.currentBalance)
	vm.currentBalance = 0
	vm.setState(&NoCashState{})
}

func (s *ProductSelectedState) dispenseProduct(vm *VendingMachine) {
	fmt.Printf("Dispensing %s\n", s.product.name)
	vm.currentBalance -= s.product.price
	if vm.currentBalance > 0 {
		fmt.Printf("Returning change: %.2f\n", vm.currentBalance)
		vm.currentBalance = 0
	}
	vm.setState(&NoCashState{})
}
