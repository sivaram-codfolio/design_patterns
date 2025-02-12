package main

import (
	"fmt"

	"github.com/sivaram-codfolio/design_patterns/builder"
)

func main() {
	var builder = builder.NewNotificationBuilder()
	builder.SetTitle("New Notification")
	builder.SetIcon("icon.png")
	builder.SetSubTitle("This is a subtitle")
	builder.SetImage("image.jpg")
	builder.SetPriority(5)
	builder.SetMessage("This is a basic notification")
	builder.SetType("alert")

	notification, err := builder.Build()

	if err != nil {
		fmt.Println("Error creating the notification: ", err)
	} else {
		fmt.Printf("Notification: %+v\n", notification)
	}
}
