package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Create card
func (s *SmartContract) CreateCard(ctx contractapi.TransactionContextInterface, cardNumber string, bankAccountId int64) (*Card, error) {
	bankAccount, err := s.GetBankAccount(ctx, bankAccountId)
	if err != nil {
		return nil, err
	}

	_, exists := bankAccount.Cards[cardNumber]
	if exists {
		return nil, fmt.Errorf("Card with given card number %s already exists in bank account")
	}

	bankAccount.Cards[cardNumber] = Card{
		CardNumber: cardNumber,
	}

	bankAccountJSON, err := json.Marshal(bankAccount)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return nil, err
	}

	card := bankAccount.Cards[cardNumber]
	return &card, nil
}

// Remove card
func (s *SmartContract) RemoveCard(ctx contractapi.TransactionContextInterface, cardNumber string, bankAccountId int64) error {
	bankAccount, err := s.GetBankAccount(ctx, bankAccountId)
	if err != nil {
		return err
	}

	_, exists := bankAccount.Cards[cardNumber]
	if !exists {
		return fmt.Errorf("Card with given card number %s does not exist in bank account ID %d", cardNumber, bankAccountId)
	}

	delete(bankAccount.Cards, cardNumber)

	bankAccountJSON, err := json.Marshal(*bankAccount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return err
	}

	return nil
}
