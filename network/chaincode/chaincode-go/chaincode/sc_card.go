package chaincode

// Create card
func (s *SmartContract) CreateCard(ctx contractapi.TransactionContextInterface, id int64, cardNumber, bankAccountId string) (*Card, error) {
	var card Card
	cardBytes, err := s.GetEntityById(ctx, "card", id)

	if err != nil {
		return nil, err
	}
	if cardBytes != nil{
		err = json.Unmarshal(cardBytes, &card)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal card: %v", err)
		}
	    return nil, fmt.Errorf("The card %s already has an car with id %s", cardNumber, card.Id)
	}

	var bankAccount BankAccount
	bankAccountBytes, err := s.GetEntityById(ctx, "bankAccount", bankAccountId)

	if err != nil {
		return nil, err
	}
	if bankAccountBytes != nil{
		err = json.Unmarshal(bankAccountBytes, &bankAccount)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal bank account: %v", err)
		}
	}

	cardId := toCardId(id)
	card := Card{
		Id:       accountId,
		CardNumber: cardNumber,
		BankAccountId: bankAccountId
	}

    bankAccount.Cards.put(card)
	
	bankAccountJSON, err := json.Marshal(bankAccount)
	if err != nil {
		return nil, err
	}
	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return nil, err
	}

	return &card, nil
}


// Remove card
func (s *SmartContract) RemoveCard(ctx contractapi.TransactionContextInterface, cardId int64, , bankAccountId string) (*Card, error) {
	var card Card
	cardBytes, err := s.GetEntityById(ctx, "card", cardId)

	if err != nil {
		return nil, err
	}
	if cardBytes != nil{
		err = json.Unmarshal(cardBytes, &card)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal card: %v", err)
		}
	    return nil, fmt.Errorf("The card %s already has an car with id %s", cardNumber, card.Id)
	}

	var bankAccount BankAccount
	bankAccountBytes, err := s.GetEntityById(ctx, "bankAccount", id)

	if err != nil {
		return nil, err
	}
	if bankAccountBytes != nil{
		err = json.Unmarshal(bankAccountBytes, &bankAccount)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal bank account: %v", err)
		}
	}

	cardId := toCardId(id)
	card := Card{
		Id:       accountId,
		CardNumber: cardNumber,
		BankAccountId: bankAccountId
	}

    RemoveCard(&bankAccount, cardId)
	
	bankAccountJSON, err := json.Marshal(bankAccount)
	if err != nil {
		return nil, err
	}
	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return nil, err
	}
	err = ctx.GetStub().PutState(card.Id, nil)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func RemoveCard(account *BankAccount, cardNumber string) {
    for i, card := range account.Cards {
        if card.Number == cardNumber {
            account.Cards = append(account.Cards[:i], account.Cards[i+1:]...)
            return
        }
    }
}