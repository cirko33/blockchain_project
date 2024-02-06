package chaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GetBankAccountsWithGTE100 returns all bank accounts with balance greater than or equal to 100
func (s *SmartContract) GetBankAccountsWithGTE100(ctx contractapi.TransactionContextInterface) ([]*BankAccount, error) {
	queryString := `{
		"selector": {
			"$and": [
				{
					"balance": {
						"$gte": 100
					}
				},
				{
					"currency": {
						"$eq": "EUR"
					}
				}
			]
		}
	}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var bankaccounts []*BankAccount
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var bankaccount BankAccount
		err = json.Unmarshal(queryResponse.Value, &bankaccount)
		if err != nil {
			return nil, err
		}
		bankaccounts = append(bankaccounts, &bankaccount)
	}

	return bankaccounts, nil
}
