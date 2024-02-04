package chaincode

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) GetEntityById(ctx contractapi.TransactionContextInterface, entityName string, id int64) ([]byte, error) {
	entity, err := ctx.GetStub().GetState(fmt.Sprintf("%s-%d", entityName, id))
	if err != nil {
		return false, fmt.Errorf("Failed to read entity (%s) with id (%s) from world state: %v", entityName, id, err)
	}

	return entity, nil
}

func toBankId(intId int64) string {
	return fmt.Sprintf("%s-%d", BANK_TYPE_NAME, intId)
}

func toPersonId(intId int64) string {
	return fmt.Sprintf("%s-%d", PERSON_TYPE_NAME, intId)
}

func toBankAccountId(intId int64) string {
	return fmt.Sprintf("%s-%d", BANK_ACCOUNT_TYPE_NAME, intId)
}

func toCardId(intId int64) string {
	return fmt.Sprintf("%s-%d", CARD_TYPE_NAME, intId)
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
	persons := buildMockPersons(id, 3)
	accounts := buildMockAccounts(persons)
	bank_id := toBankId(id)
	return Bank{
		Id:       bank_id,
		Location: fmt.Sprintf("Location_%s", bank_id),
		PIB:      fmt.Sprintf("PIB_%s", bank_id),
		Persons:  persons,
		Accounts: accounts,
	}
}

func buildMockPersons(start_id int64, count int64) []Person {
	var result []Person

	for i := int64(0); i < count; i++ {
		id := start_id + i
		str_id := toPersonId(id)
		result = append(result, Person{
			Id:      str_id,
			Name:    fmt.Sprintf("Person_%s", str_id),
			Surname: fmt.Sprintf("Personic_%s", str_id),
			Email:   fmt.Sprintf("mejl_%s@gmail.com", str_id),
		})
	}

	return result
}

func buildMockAccounts(persons []Person) []BankAccount {
	var result []BankAccount
	currency_labels := []string{"RSD", "EUR"}
	id_counter := 0
	for i, person := range persons {
		account_count := 1 + i%2
		for j := 0; j < account_count; j++ {
			account_id := id_counter
			id_counter++
			account_id_str := toBankAccountId(int64(account_id))
			result = append(result, BankAccount{
				Id:       account_id_str,
				PersonId: person.Id,
				Balance:  getRandomBalance(),
				Currency: currency_labels[j],
				Cards: []Card{
					{
						CardNumber:    fmt.Sprintf("%s_123123", account_id_str),
						BankAccountId: account_id_str,
					},
				},
			})
		}
	}
	return result
}

func getRandomBalance() float64 {
	return rand.Float64() * (1000 - 100)
}
