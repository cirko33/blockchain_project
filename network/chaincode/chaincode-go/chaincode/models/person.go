package models

type Person struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	//Accounts map[string]Account `json:"accounts"`
}
