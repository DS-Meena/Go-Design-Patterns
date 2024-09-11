package main

func main() {
	vm := &VendingMachine{
		currentState: &NoCashState{},
		products: []Product{
			{id: 1, name: "Cola", price: 15},
			{id: 2, name: "Chips", price: 10},
			{id: 3, name: "Candy", price: 5},
		},
	}

	vm.insertCash(10)
	vm.selectProduct(2)
	vm.dispenseProduct()

	vm.insertCash(7)
	vm.selectProduct(3)
	vm.dispenseProduct()

	vm.insertCash(5)
	vm.selectProduct(2)
	vm.insertCash(5)
	vm.selectProduct(2)
	vm.dispenseProduct()

	vm.insertCash(15)
	vm.selectProduct(1)
	vm.cancelTransaction()
}
