package chaincode

import (
	"fmt"
	"math/rand"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) GetEntityById(ctx contractapi.TransactionContextInterface, entityName string, id int64) ([]byte, error) {
	entity, err := ctx.GetStub().GetState(toEntityId(entityName, id))
	if err != nil {
		return nil, fmt.Errorf("failed to read entity (%s) with id (%d) from world state: %v", entityName, id, err)
	}

	return entity, nil
}

func (s *SmartContract) EntityExists(ctx contractapi.TransactionContextInterface, entityName string, id int64) (bool, error) {
	itemJSON, err := ctx.GetStub().GetState(toEntityId(entityName, id))
	if err != nil {
		return false, fmt.Errorf("failed to read entity of type '%s' with id '%d' from world state: %v", entityName, id, err)
	}

	return itemJSON != nil, nil
}

func toEntityId(typeName string, intId int64) string {
	return fmt.Sprintf("%s-%d", typeName, intId)
}

func toBankId(intId int64) string {
	return toEntityId(BANK_TYPE_NAME, intId)
}

func toPersonId(intId int64) string {
	return toEntityId(PERSON_TYPE_NAME, intId)
}

func toBankAccountId(intId int64) string {
	return toEntityId(BANK_ACCOUNT_TYPE_NAME, intId)
}

func buildMockBanks(count int64) []Bank {
	var result []Bank
	for bank_id := int64(0); bank_id < count; bank_id++ {
		bank := buildMockBank(bank_id)
		result = append(result, bank)
	}
	return result
}

func buildMockBank(id int64) Bank {
	bank_id := toBankId(id)
	return Bank{
		Id:       bank_id,
		Location: fmt.Sprintf("Location_%s", bank_id),
		PIB:      fmt.Sprintf("PIB_%s", bank_id),
	}
}

func buildMockPersons(count int64) []Person {
	var result []Person

	for i := int64(0); i < count; i++ {
		str_id := toPersonId(i)
		result = append(result, Person{
			Id:      str_id,
			Name:    fmt.Sprintf("Person_%s", str_id),
			Surname: fmt.Sprintf("Personic_%s", str_id),
			Email:   fmt.Sprintf("mejl_%s@gmail.com", str_id),
		})
	}

	return result
}

func buildMockAccounts(banks []Bank, persons []Person) []BankAccount {
	result := make([]BankAccount, 0)

	id_counter := int64(0)
	for _, bank := range banks {
		for person_idx, person := range persons {
			newAccounts := buildAccountsForPerson(person, int64(person_idx), bank, id_counter)
			for _, newAcc := range newAccounts {
				result = append(result, newAcc)
				id_counter++
			}
		}
	}

	return result
}

func buildAccountsForPerson(person Person, index int64, bank Bank, startId int64) []BankAccount {
	var result []BankAccount
	currency_labels := []string{"RSD", "EUR"}
	id_counter := startId
	account_count := 1 + startId%2
	for j := 0; j < int(account_count); j++ {
		account_id := id_counter
		id_counter++
		account_id_str := toBankAccountId(int64(account_id))
		card_number := fmt.Sprintf("%s_123123", account_id_str)
		result = append(result, BankAccount{
			Id:       account_id_str,
			PersonId: person.Id,
			BankId:   bank.Id,
			Balance:  getRandomBalance(),
			Currency: currency_labels[j],
			Cards: map[string]Card{
				card_number: Card{CardNumber: card_number},
			},
		})
	}
	return result
}

func getRandomBalance() float64 {
	return 100 * rand.Float64() * (900)
}
