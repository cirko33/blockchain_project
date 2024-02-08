package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Get card
func (s *SmartContract) GetCard(ctx contractapi.TransactionContextInterface, id int64) (*Card, error) {
	cardJSON, err := s.GetEntityById(ctx, CARD_TYPE_NAME, id)
	if err != nil {
		return nil, err
	}

	if  cardJSON == nil {
		return nil, fmt.Errorf("Card with given id %d does not exist", id)
	}

	var card Card
	err = json.Unmarshal(cardJSON, &card)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal card: %v", err)
	}

	return &card, nil
}

// Create card
func (s *SmartContract) CreateCard(ctx contractapi.TransactionContextInterface, cardNumber string, cardId, bankAccountId int64) (*Card, error) {
	bankAccount, err := s.GetBankAccount(ctx, bankAccountId)
	if err != nil {
		return nil, err
	}

	if bankAccount == nil {
		return nil, fmt.Errorf("Bank account does not exist")
	}

	for _, c := range bankAccount.Cards {
        if c.Id == ToCardId(cardId) {
			return nil, fmt.Errorf("Card with given card number %s already exists in bank account", cardNumber)
        }
    }

	card := Card{
		Id: ToCardId(cardId),
		CardNumber: cardNumber,
		BankAccountId: ToBankAccountId(bankAccountId),
	}

	bankAccount.Cards = append(bankAccount.Cards,card)

	bankAccountJSON, err := json.Marshal(bankAccount)
	if err != nil {
		return nil, err
	}

	cardJSON, err := json.Marshal(card)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(card.Id, cardJSON)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

// Remove card
func (s *SmartContract) RemoveCard(ctx contractapi.TransactionContextInterface, cardId, bankAccountId int64) (*Card, error) {
	bankAccount, err := s.GetBankAccount(ctx, bankAccountId)
	if err != nil {
		return nil, err
	}

	card, err1 := s.GetCard(ctx,cardId)

	if err1 != nil {
		return nil, err1
	}

	if card.BankAccountId != ToBankAccountId(bankAccountId){
		return nil, fmt.Errorf("Card with given card id %d does not exist in bank account with id %d", cardId, bankAccountId)
	}

	index, found := FindCardIndexById(bankAccount.Cards, ToCardId(cardId))
    if found {
        fmt.Printf("Found card index: %d\n", index)
    } else {
        return nil, fmt.Errorf("Card not found.")
    }

	if index < 0 || index >= len(bankAccount.Cards) {
        return nil, fmt.Errorf("Index out of range")
    }

    bankAccount.Cards = append(bankAccount.Cards[:index], bankAccount.Cards[index+1:]...)

	bankAccountJSON, err := json.Marshal(*bankAccount)
	if err != nil {
		return nil, err
	}

	cardJSON, err1 := json.Marshal(*card)
	if err1 != nil {
		return nil, err1
	}

	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(card.Id, cardJSON)
	if err != nil {
		return nil, err
	}

	return card, nil
}


func FindCardIndexById(cards []Card, id string) (int, bool) {
    for index, card := range cards {
        if card.Id == id {
            return index, true
        }
    }
    return -1, false
}