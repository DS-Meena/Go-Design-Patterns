package main

import (
	"strconv"
)

type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

type iPhone struct {
	observers []emailAlert // list of observers subscribed via email
	price     int
}

func (i *iPhone) priceChange(price int) {
	i.price = price

	// notify customers regarding discount
	i.notifyAll()
}

func (i *iPhone) registerObserver(obs emailAlert) {
	i.observers = append(i.observers, obs)
}

func (i *iPhone) unregisterObserver(sub emailAlert) {
	var newObservers []emailAlert
	for _, obs := range i.observers {
		if obs.Email != sub.Email {
			newObservers = append(newObservers, obs)
		}
	}

	i.observers = newObservers
}

// notify all the observers
func (i *iPhone) notifyAll() {
	for _, obs := range i.observers {
		obs.onUpdate(strconv.Itoa(i.price))
	}
}
