package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Bank struct {
	Id       string   `json:"id"`
	Location string   `json:"location"`
	PIB      string   `json:"pib"`
	Persons  []Person `json:"persons"`
}

type Person struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Surname  string        `json:"surname"`
	Email    string        `json:"email"`
	Accounts []BankAccount `json:"accounts"`
}

type BankAccount struct {
	Id       string  `json:"id"`
	PersonId string  `json:"personId"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	Cards    []Card  `json:"cards"`
}

type Card struct {
	CardNumber    string `json:"cardNumber"`
	BankAccountId string `json:"bankAccountId"`
}

type ExchangeRate struct {
	BuyingRate  float64 `json:"buyingRate"`
	MiddleRate  float64 `json:"middleRate"`
	SellingRate float64 `json:"sellingRate"`
}
