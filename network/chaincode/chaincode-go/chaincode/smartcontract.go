package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	banks := BuildMockBanks(4)
	persons := BuildMockPersons(12)
	bankAccounts := BuildMockAccounts(banks, persons)

	for _, bank := range banks {
		bankJSON, err := json.Marshal(bank)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(bank.Id, bankJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, person := range persons {
		personJSON, err := json.Marshal(person)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(person.Id, personJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, bankAccount := range bankAccounts {
		bankAccountJSON, err := json.Marshal(bankAccount)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}
