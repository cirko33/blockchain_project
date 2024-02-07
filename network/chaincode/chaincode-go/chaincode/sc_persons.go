package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) CreatePerson(ctx contractapi.TransactionContextInterface, id int64, name string, surname string, email string) (*Person, error) {
	exists, err := s.EntityExists(ctx, PERSON_TYPE_NAME, id)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("Person with given ID %d already exists", id)
	}

	newId := ToPersonId(id)
	newPerson := Person{
		Id:      newId,
		Name:    name,
		Surname: surname,
		Email:   email,
	}

	personJSON, err := json.Marshal(newPerson)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(newPerson.Id, personJSON)
	if err != nil {
		return nil, err
	}

	return &newPerson, nil
}

func (s *SmartContract) GetPerson(ctx contractapi.TransactionContextInterface, id int64) (*Person, error) {
	stringId := ToPersonId(id)
	return s._GetPersonInternal(ctx, stringId)
}

func (s *SmartContract) _GetPersonInternal(ctx contractapi.TransactionContextInterface, id string) (*Person, error) {
	personJSON, err := ctx.GetStub().GetState(id)
	if err != nil || personJSON == nil {
		return nil, fmt.Errorf("failed to read entity with id '%s' from world state: %v", id, err)
	}

	var person Person
	err = json.Unmarshal(personJSON, &person)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal person: %v", err)
	}

	return &person, nil
}

func (s *SmartContract) GetAllPersons(ctx contractapi.TransactionContextInterface) ([]*Person, error) {
	personsIterator, err := ctx.GetStub().GetQueryResult(BuildQueryIdStartsWith(PERSON_TYPE_NAME))
	if err != nil {
		return nil, err
	}

	defer personsIterator.Close()
	var persons []*Person
	for personsIterator.HasNext() {
		queryResponse, err := personsIterator.Next()
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

func (s *SmartContract) GetPersonByBankAccount(ctx contractapi.TransactionContextInterface, bankAccountId int64) (*Person, error) {
	bankAccount, err := s.GetBankAccount(ctx, bankAccountId)
	if err != nil {
		return nil, err
	}

	if bankAccount == nil {
		return nil, fmt.Errorf("Bank account with given ID %d not found", bankAccountId)
	}

	return s._GetPersonInternal(ctx, bankAccount.PersonId)
}
