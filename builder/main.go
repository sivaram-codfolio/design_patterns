package main

import "fmt"

func main() {
	var builder = newNotificationBuilder()
	builder.setTitle("New Notification")
	builder.setIcon("icon.png")
	builder.setSubTitle("This is a subtitle")
	builder.setImage("image.jpg")
	builder.setPriority(5)
	builder.setMessage("This is a basic notification")
	builder.setType("alert")

	notification, err := builder.build()

	if err != nil {
		fmt.Println("Error creating the notification: ", err)
	} else {
		fmt.Printf("Notification: %+v\n", notification)
	}
}
