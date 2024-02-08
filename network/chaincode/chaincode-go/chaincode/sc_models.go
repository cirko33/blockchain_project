package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Bank struct {
	Id             string `json:"id"`
	Location       string `json:"location"`
	PIB            string `json:"pib"`
	FoundationYear uint32 `json:"foundationYear"`
}

type Person struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type BankAccount struct {
	Id       string          `json:"id"`
	PersonId string          `json:"personId"`
	BankId   string          `json:"bankId"`
	Balance  float64         `json:"balance"`
	Currency string          `json:"currency"`
	Cards    []Card          `json:"cards"`
}

type Card struct {
	Id              string        `json:"id"`
	CardNumber      string        `json:"cardNumber"`
	BankAccountId   string        `json:"bankAccountId"`
}

type ExchangeRate struct {
	BuyingRate  float64 `json:"buyingRate"`
	MiddleRate  float64 `json:"middleRate"`
	SellingRate float64 `json:"sellingRate"`
}
