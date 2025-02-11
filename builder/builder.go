package main

import "fmt"

type NotificationBuilder struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Message  string `json:"message"`
	Image    string `json:"image"`
	Icon     string `json:"icon"`
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) setTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) setSubTitle(subTitle string) {
	nb.SubTitle = subTitle
}

func (nb *NotificationBuilder) setMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) setImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) setIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) setPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) setType(notType string) {
	nb.Type = notType
}

func (nb *NotificationBuilder) build() (*Notification, error) {
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("You need to specify a subtitle when using an icon")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("Priority must be 0 to 5")

	}

	return &Notification{
		Title:    nb.Title,
		SubTitle: nb.SubTitle,
		Message:  nb.Message,
		Image:    nb.Image,
		Icon:     nb.Icon,
		Priority: nb.Priority,
		Type:     nb.Type,
	}, nil
}
