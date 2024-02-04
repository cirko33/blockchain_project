package chaincode

import (
	"fmt"
	"math/rand"
)

func buildMockBanks(count int) []Bank {
	var result []Bank
	for bank_id := 0; bank_id < count; bank_id++ {
		bank := buildMockBank(bank_id)
		result = append(result, bank)
	}
	return result
}

func buildMockBank(id int) Bank {
	persons := buildMockPersons(id, 3)
	accounts := buildMockAccounts(persons)
	return Bank{
		Id:       fmt.Sprint(id),
		Location: fmt.Sprintf("Location_%d", id),
		PIB:      fmt.Sprintf("PIB_%d", id),
		Persons:  persons,
		Accounts: accounts,
	}
}

func buildMockPersons(start_id int, count int) []Person {
	var result []Person

	for i := 0; i < count; i++ {
		id := fmt.Sprintf("%d", start_id+i)
		result = append(result, Person{
			Id:      id,
			Name:    fmt.Sprintf("Person_%s", id),
			Surname: fmt.Sprintf("Personic_%s", id),
			Email:   fmt.Sprintf("mejl_%s@gmail.com", id),
		})
	}

	return result
}

func buildMockAccounts(persons []Person) []BankAccount {
	var result []BankAccount
	currency_labels := []string{"RSD", "EUR"}
	for i, person := range persons {
		account_count := 1 + i%2
		for j := 0; j < account_count; j++ {
			account_id := fmt.Sprintf("%s_%d", person.Id, j)
			result = append(result, BankAccount{
				Id:       account_id,
				PersonId: person.Id,
				Balance:  getRandomBalance(),
				Currency: currency_labels[j],
				Cards: []Card{
					{
						CardNumber:    fmt.Sprintf("%s_123123", account_id),
						BankAccountId: account_id,
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
