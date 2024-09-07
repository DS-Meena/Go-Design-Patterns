package main

func main() {
	customer1 := emailAlert{
		Email: "dsm@gmail.com",
	}
	customer2 := emailAlert{
		Email: "random@gmail.com",
	}

	product := &iPhone{
		price: 100,
	}

	// register customers
	product.registerObserver(customer1)
	product.registerObserver(customer2)

	// customers will get notifications for price change
	product.priceChange(90)
	product.priceChange(80)
	product.unregisterObserver(customer2)

	product.priceChange(50)
}
