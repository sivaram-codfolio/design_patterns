package builder

type Notification struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Message  string `json:"message"`
	Image    string `json:"image"`
	Icon     string `json:"icon"`
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}
