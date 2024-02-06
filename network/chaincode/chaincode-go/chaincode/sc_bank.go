package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) CreateBank(ctx contractapi.TransactionContextInterface, id int64, location string, pib string) (*Bank, error) {
	exists, err := s.EntityExists(ctx, BANK_TYPE_NAME, id)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("Bank with given ID %d already exists", id)
	}

	newId := ToBankId(id)
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
	queryString := BuildQueryIdStartsWith(BANK_TYPE_NAME)
	fmt.Printf("Query string: '%s'", queryString)

	banksIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}

	defer banksIterator.Close()
	var banks []*Bank
	for banksIterator.HasNext() {
		queryResponse, err := banksIterator.Next()
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
