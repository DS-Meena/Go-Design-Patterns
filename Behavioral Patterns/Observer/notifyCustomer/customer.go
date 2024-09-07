package main

import "fmt"

type observer interface {
	onUpdate(data string)
}

type emailAlert struct {
	Email string
}

func (c *emailAlert) onUpdate(data string) {
	fmt.Println("Mail sent on ", c.Email, data)
}

// we can create smsAlert observers too.
// type smsAlert struct {
// 	Phone string
// }

// func (c *smsAlert) onUpdate(data string) {
// 	fmt.Println("SMS sent on ", c.Phone, data)
// }
