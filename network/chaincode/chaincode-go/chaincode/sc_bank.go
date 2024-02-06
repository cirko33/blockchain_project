package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) GetBank(ctx contractapi.TransactionContextInterface, id int64) (*Bank, error) {
	bankJSON, err := s.GetEntityById(ctx, BANK_TYPE_NAME, id)

	if err != nil || bankJSON == nil {
		return nil, err
	}

	var bank Bank
	err = json.Unmarshal(bankJSON, &bank)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bank: %v", err)
	}

	return &bank, nil
}

func (s *SmartContract) GetAllBanks(ctx contractapi.TransactionContextInterface) ([]*Bank, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}

	defer resultsIterator.Close()
	var banks []*Bank
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var bank Bank
		err = json.Unmarshal(queryResponse.Value, &bank)
		if err != nil {
			return nil, err
		}
		banks = append(banks, &bank)
	}

	return banks, nil
}

func (s *SmartContract) CreateBank(ctx contractapi.TransactionContextInterface, id int64, location string, pib string) (*Bank, error) {
	exists, err := s.EntityExists(ctx, BANK_TYPE_NAME, id)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("Bank with given ID %d already exists", id)
	}

	newId := toBankId(id)
	newBank := Bank{
		Id:       newId,
		Location: location,
		PIB:      pib,
	}

	bankJSON, err := json.Marshal(newBank)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(newBank.Id, bankJSON)
	if err != nil {
		return nil, err
	}

	return &newBank, nil
}
