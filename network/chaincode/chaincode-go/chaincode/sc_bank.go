package chaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

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

// func (s *SmartContract) CreateBank(ctx contractapi.TransactionContextInterface, id int64, location string, pib string) error {
// 	result, err := s.GetEntityById(ctx, BANK_TYPE_NAME, id)
// }

// func (s *SmartContract) DeleteBank(ctx contractapi.TransactionContextInterface, id string) error {
// 	// todo
// }

// func (s *SmartContract) GetBank(ctx contractapi.TransactionContextInterface, id string) error {
// 	// todo
// }
