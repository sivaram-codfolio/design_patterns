package builder

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

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subTitle string) {
	nb.SubTitle = subTitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.Type = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("you need to specify a subtitle when using an icon")
	}

	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority must be 0 to 5")

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
