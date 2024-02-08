package chaincode

import (
	"fmt"
	"math/rand"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) GetEntityById(ctx contractapi.TransactionContextInterface, entityName string, id int64) ([]byte, error) {
	entity, err := ctx.GetStub().GetState(ToEntityId(entityName, id))
	if err != nil {
		return nil, fmt.Errorf("failed to read entity (%s) with id (%d) from world state: %v", entityName, id, err)
	}

	return entity, nil
}

func (s *SmartContract) EntityExists(ctx contractapi.TransactionContextInterface, entityName string, id int64) (bool, error) {
	itemJSON, err := ctx.GetStub().GetState(ToEntityId(entityName, id))
	if err != nil {
		return false, fmt.Errorf("failed to read entity of type '%s' with id '%d' from world state: %v", entityName, id, err)
	}

	return itemJSON != nil, nil
}

func ToEntityId(typeName string, intId int64) string {
	return fmt.Sprintf("%s-%d", typeName, intId)
}

func ToBankId(intId int64) string {
	return ToEntityId(BANK_TYPE_NAME, intId)
}

func ToPersonId(intId int64) string {
	return ToEntityId(PERSON_TYPE_NAME, intId)
}

func ToBankAccountId(intId int64) string {
	return ToEntityId(BANK_ACCOUNT_TYPE_NAME, intId)
}

func ToCardId(intId int64) string {
	return ToEntityId(CARD_TYPE_NAME, intId)
}

func BuildMockBanks(count int64) []Bank {
	var result []Bank
	for bank_id := int64(0); bank_id < count; bank_id++ {
		bank := BuildMockBank(bank_id)
		result = append(result, bank)
	}
	return result
}

func BuildMockBank(id int64) Bank {
	bank_id := ToBankId(id)
	return Bank{
		Id:             bank_id,
		Location:       fmt.Sprintf("Location_%d", id),
		PIB:            fmt.Sprintf("PIB_%d", id),
		FoundationYear: uint32(rand.Intn(50) + 1950),
	}
}

func BuildMockPersons(count int64) []Person {
	var result []Person

	for i := int64(0); i < count; i++ {
		str_id := ToPersonId(i)
		result = append(result, Person{
			Id:      str_id,
			Name:    fmt.Sprintf("Person_%d", i),
			Surname: fmt.Sprintf("Personic_%d", i),
			Email:   fmt.Sprintf("mejl_%d@gmail.com", i),
		})
	}

	return result
}

func BuildMockAccounts(banks []Bank, persons []Person) []BankAccount {
	result := make([]BankAccount, 0)

	id_counter := int64(0)
	for _, bank := range banks {
		for person_idx, person := range persons {
			newAccounts := BuildAccountsForPerson(person, int64(person_idx), bank, id_counter)
			for _, newAcc := range newAccounts {
				result = append(result, newAcc)
				id_counter++
			}
		}
	}

	return result
}

func BuildAccountsForPerson(person Person, index int64, bank Bank, startId int64) []BankAccount {
	var result []BankAccount
	currency_labels := []string{"RSD", "EUR"}
	id_counter := startId
	account_count := 1 + startId%2
	for j := 0; j < int(account_count); j++ {
		account_id := id_counter
		id_counter++
		account_id_str := ToBankAccountId(int64(account_id))
		card_number := fmt.Sprintf("%s_123123", account_id_str)
		result = append(result, BankAccount{
			Id:       account_id_str,
			PersonId: person.Id,
			BankId:   bank.Id,
			Balance:  GetRandomBalance(),
			Currency: currency_labels[j],
			Cards: []Card{
				 {   Id: ToCardId(account_id),
					CardNumber: card_number,
				},
			},
		})
	}
	return result
}

func GetRandomBalance() float64 {
	return 100 * rand.Float64() * (900)
}

func BuildQueryIdStartsWith(prefix string) string {
	return fmt.Sprintf("{\"selector\": {\"_id\": { \"$regex\": \"^(%s-)\" } } }", prefix)
}

func BuildQueryForEntityType(entityName string, selectors string) string {
	return fmt.Sprintf("{\"selector\": {\"_id\": { \"$regex\": \"^(%s-)\" }, %s } }", entityName, selectors)
}

func BuildQueryFieldContains(entityName string, fieldName string, substring string) string {
	fieldSelector := BuildContainsSelector(fieldName, substring)
	return BuildQueryForEntityType(entityName, fieldSelector)
}

func BuildContainsSelector(fieldName string, substring string) string {
	return fmt.Sprintf("\"%s\": { \"$regex\": \"%s\" }", fieldName, substring)
}
