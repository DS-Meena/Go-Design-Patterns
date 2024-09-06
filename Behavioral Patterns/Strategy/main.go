package main

import "fmt"

// Strategy interface
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Conrete strategy: Credit card
type CreditCardPayment struct {}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card 💳", amount)
}

// Concrete strategy: Bhim
type BhimPayment struct {}

func (b *BhimPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using bhim 🌐", amount)
}

// Context
type ShoppingCart struct {
	paymentStrategy PaymentStrategy    // interchangeable
}

// we pass interface to set the algorithm at runtime
func (s *ShoppingCart) SetPaymentStrategy(ps PaymentStrategy) {
	s.paymentStrategy = ps
}

func (s *ShoppingCart) Checkout(amount float64) string {
	return s.paymentStrategy.Pay(amount)
}

func main() {
	cart := &ShoppingCart{}

	// Use Credit Card
	cart.SetPaymentStrategy(&CreditCardPayment{})
	fmt.Println(cart.Checkout(100.89))

	// Switch to Bhim
	cart.SetPaymentStrategy(&BhimPayment{})
	fmt.Println(cart.Checkout(32.74))
}