package main

import "fmt"

// Mediator
type ChatRoom struct{}

func (cr *ChatRoom) DisplayMessage(user *User, message string) {
	fmt.Printf("[%s]: %s\n", user.name, message)
}

// Colleague
type User struct {
	name     string
	chatRoom *ChatRoom
}

func NewUser(name string, chatRoom *ChatRoom) *User {
	return &User{name: name, chatRoom: chatRoom}
}

func (u *User) Send(message string) {
	u.chatRoom.DisplayMessage(u, message)
}

func main() {
	chatRoom := &ChatRoom{}

	john := NewUser("John smith", chatRoom)
	jane := NewUser("Jane smith", chatRoom)

	john.Send("Hi there!")
	jane.Send("Stop this")
}
