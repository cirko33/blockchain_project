package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SearchPersonsByName returns all persons with given name
func (s *SmartContract) SearchPersonsByName(ctx contractapi.TransactionContextInterface, nameQuery string) ([]*Person, error) {
	queryString := BuildQueryFieldContains(PERSON_TYPE_NAME, "name", nameQuery)
	return s.SearchPersonsRaw(ctx, queryString)
}

// SearchPersonsBySurname returns all persons with given surname
func (s *SmartContract) SearchPersonsBySurname(ctx contractapi.TransactionContextInterface, surnameQuery string) ([]*Person, error) {
	queryString := BuildQueryFieldContains(PERSON_TYPE_NAME, "surname", surnameQuery)
	return s.SearchPersonsRaw(ctx, queryString)
}

// SearchPersonsBySurnameAndEmail returns all persons with given surname and email
func (s *SmartContract) SearchPersonsBySurnameAndEmail(ctx contractapi.TransactionContextInterface, surname string, email string) ([]*Person, error) {
	surnameSelector := BuildContainsSelector("surname", surname)
	emailSelector := BuildContainsSelector("email", email)
	selectors := fmt.Sprintf("%s, %s", surnameSelector, emailSelector)
	queryString := BuildQueryForEntityType(PERSON_TYPE_NAME, selectors)
	return s.SearchPersonsRaw(ctx, queryString)
}

// Not rich query
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

// GetAccountsWithMoreThanBalance returns all bank accounts with balance greater than or equal to parameter balance and in the same currency
func (s *SmartContract) GetAccountsWithMoreThanBalance(ctx contractapi.TransactionContextInterface, currency string, balance float64) ([]*BankAccount, error) {
	queryString := `
	{
		"selector": {
			"$and": [
				{
					"balance": {
						"$gte": ` + fmt.Sprintf("%f", balance) + `
					}
				},
				{
					"currency": {
						"$eq": "` + currency + `"
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

// GetBanksOlderThan returns all banks older than parameter year
func (s *SmartContract) GetBanksOlderThan(ctx contractapi.TransactionContextInterface, year int) ([]*Bank, error) {
	queryString := `
	{
		"selector": {
			"foundationYear": {
				"$lt": ` + fmt.Sprintf("%d", year) + `
			}
		}
	}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

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

// GetBanksByLocation returns all banks in the given location
func (s *SmartContract) GetBanksByLocation(ctx contractapi.TransactionContextInterface, location string) ([]*Bank, error) {
	queryString := `
	{
		"selector": {
			"location": {
				"$eq": "` + location + `"
			}
		}
	}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

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

// GetBankAccountsByPerson returns all bank accounts of the given person
func (s *SmartContract) GetBankAccountsByPerson(ctx contractapi.TransactionContextInterface, personId string) ([]*BankAccount, error) {
	queryString := `
	{
		"selector": {
			"personId": {
				"$eq": "` + personId + `"
			}
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

func (s *SmartContract) GetBankAccountsByBank(ctx contractapi.TransactionContextInterface, bankId string) ([]*BankAccount, error) {
	queryString := `
	{
		"selector": {
			"bankId": {
				"$eq": "` + bankId + `"
			}
		}
	}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var bankAccounts []*BankAccount
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
		bankAccounts = append(bankAccounts, &bankaccount)
	}

	return bankAccounts, nil
}

type BankWithAccounts struct {
	Bank         Bank           `json:"bank"`
	BankAccounts []*BankAccount `json:"bankAccounts"`
}

// GetAllBanksWithAccounts returns all banks with their accounts
func (s *SmartContract) GetAllBanksWithAccounts(ctx contractapi.TransactionContextInterface) ([]*BankWithAccounts, error) {
	banks, err := s.GetAllBanks(ctx)
	if err != nil {
		return nil, err
	}

	var banksWithAccounts []*BankWithAccounts
	for _, bank := range banks {

		bankaccounts, err := s.GetBankAccountsByBank(ctx, bank.Id)
		if err != nil {
			return nil, err
		}

		banksWithAccounts = append(banksWithAccounts, &BankWithAccounts{Bank: *bank, BankAccounts: bankaccounts})
	}

	return banksWithAccounts, nil
}
