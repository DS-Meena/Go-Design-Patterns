package main

import "fmt"

func main() {
	// Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()

	// Use the builder object to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetSubTitle("This is a subtitle")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(4)
	bldr.SetMessage("This is a basic Notification with some text in it")
	bldr.SetType("alert")

	// Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}

}
