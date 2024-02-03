package models

type Account struct {
	Id       string `json:"id"`
	Balance  int    `json:"balance"`
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Cards    []Card `json:"cards"`
}
