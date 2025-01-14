package main

import (
	"reflect"
)

func NilDiffs() {
	println(
		"Difference between a normal 'nil'",
		"and an empty interface:\n ",
	)
	AdultNil := NewAdult(17) // Returns nil
	// AdultNil.BuyWithCard() // Panic nil pointer

	AdultInterface := NewAdult(23)
	AdultInterface.BuyWithCard()

	println(
		"Using reflect type,\n 'pure nil':",
		reflect.TypeOf(AdultNil),
		"vs 'interface nil':",
		reflect.TypeOf(&AdultInterface),
	) // different values

	println("pure nil == interface nil?\n",
		AdultNil == AdultInterface,
	) // false

	println("Printing both values:\n",
		AdultNil,
		AdultInterface,
	) // same value printed

	// Nil channels
	var ch chan int // nil channel

	select {
	case msg := <-ch:
		println("Received:", msg)
	case <-ch: // This will never be selected as the channel is nil
		println("Channel is closed or nil.")
	default:
		println("No data received, channel is nil.")
	}
}

type CreditCardOwner interface {
	BuyWithCard()
}

type AdultWithCreditCard struct{}

// Needs to be assign to pointer to be called from interface
func (i *AdultWithCreditCard) BuyWithCard() {
	// If the pointer is nill but not used, the method still can be called
	print("Buying with credit card")
}

func NewAdult(n int) CreditCardOwner {
	if n < 18 {
		// Returns a nil value with the interface type
		return nil // This should be avoided
	}
	var a *AdultWithCreditCard = nil
	// Returns a pointer to a nil value with the struct type
	return a
}
