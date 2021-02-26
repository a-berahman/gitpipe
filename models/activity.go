package models

//Activity represents Activity data model
type Activity struct {
	ID      int    `json:"id"`
	Note    string `json:"note"`
	Subject string `json:"subject"`
}
