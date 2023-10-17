package dto

type ResponseData struct {
	ID         string   `json:"id"`
	Indicators []string `json:"indicators"`
	Passion    string   `json:"passion"`
	Time       string   `json:"time"`
}
