package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) SearchPersonsByName(ctx contractapi.TransactionContextInterface, nameQuery string) ([]*Person, error) {
	queryString := BuildQueryFieldContains(PERSON_TYPE_NAME, "name", nameQuery)
	return s.SearchPersonsRaw(ctx, queryString)
}

func (s *SmartContract) SearchPersonsBySurname(ctx contractapi.TransactionContextInterface, surnameQuery string) ([]*Person, error) {
	queryString := BuildQueryFieldContains(PERSON_TYPE_NAME, "surname", surnameQuery)
	return s.SearchPersonsRaw(ctx, queryString)
}

func (s *SmartContract) SearchPersonsBySurnameAndEmail(ctx contractapi.TransactionContextInterface, surname string, email string) ([]*Person, error) {
	surnameSelector := BuildContainsSelector("surname", surname)
	emailSelector := BuildContainsSelector("email", email)
	selectors := fmt.Sprintf("%s, %s", surnameSelector, emailSelector)
	queryString := BuildQueryForEntityType(PERSON_TYPE_NAME, selectors)
	return s.SearchPersonsRaw(ctx, queryString)
}

func (s *SmartContract) SearchPersonsByBankAccount(ctx contractapi.TransactionContextInterface, bankAccountId int64) (*Person, error) {
	bankAccount, err := s.GetBankAccount(ctx, bankAccountId)
	if err != nil {
		return nil, err
	}

	if bankAccount == nil {
		return nil, fmt.Errorf("Bank account with given ID %d not found", bankAccountId)
	}

	return s._GetPersonInternal(ctx, bankAccount.PersonId)
}

func (s *SmartContract) SearchPersonsRaw(ctx contractapi.TransactionContextInterface, queryString string) ([]*Person, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var persons []*Person
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var person Person
		err = json.Unmarshal(queryResponse.Value, &person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}

	return persons, nil
}

// GetBankAccountsWithGTE100 returns all bank accounts with balance greater than or equal to 100
func (s *SmartContract) GetBankAccountsWithGTE100(ctx contractapi.TransactionContextInterface) ([]*BankAccount, error) {
	selectors := `
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
		]`

	queryString := BuildQueryForEntityType(BANK_ACCOUNT_TYPE_NAME, selectors)
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
