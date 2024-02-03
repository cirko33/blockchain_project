package models

type Card struct {
	CardNumber string `json:"cardNumber"`
	Owner      string `json:"owner"`
	Expiration string `json:"expiration"`
	CVV        string `json:"cvv"`
}
