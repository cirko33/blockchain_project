package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}

type BankAccount struct {
	Id        string   `json:"id"`
	PersonId  string   `json:"personId"`
	Balance   float64  `json:"balance"`
	Currency  string   `json:"currency"`
	Cards     []Card   `json:"cards"`
}

type Card struct {
	CardNumber string `json:"cardNumber"`
	BankAccountId  string `json:"bankAccountId"`
}

type ExchangeRate struct {
	BuyingRate  float64 `json:"buyingRate"`
	MiddleRate  float64 `json:"middleRate"`
	SellingRate float64 `json:"sellingRate"`
}
