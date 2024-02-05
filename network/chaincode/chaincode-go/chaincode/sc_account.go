package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Create bank account
func (s *SmartContract) CreateBankAccount(ctx contractapi.TransactionContextInterface, id int64, personId, currency string, balance float64) (*BankAccount, error) {
	var bankAccount BankAccount
	bankAccountBytes, err := s.GetEntityById(ctx, "bankAccount", id)

	if err != nil {
		return nil, err
	}
	if bankAccountBytes != nil {
		err = json.Unmarshal(bankAccountBytes, &bankAccount)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal bank account: %v", err)
		}
		if bankAccount.PersonId == personId {
			return nil, fmt.Errorf("The person with id %s already has an bank account with id %s", personId, bankAccount.Id)
		}
	}

	accountId := toBankAccountId(id)
	account := BankAccount{
		Id:       accountId,
		PersonId: personId,
		Balance:  balance,
		Currency: currency,
		Cards:    []Card{},
	}

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}
	err = ctx.GetStub().PutState(account.Id, accountJSON)
	if err != nil {
		return nil, err
	}

	return &bankAccount, nil
}

func (s *SmartContract) CheckAccountCurrencies(ctx contractapi.TransactionContextInterface, fromAccountId, recipientId string) (bool, error) {
	accountJSON, err := ctx.GetStub().GetState(fromAccountId)
	if err != nil {
		return false, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if accountJSON == nil {
		return false, fmt.Errorf("The bank account %s does not exist", fromAccountId)
	}

	recipientJSON, err := ctx.GetStub().GetState(recipientId)
	if err != nil {
		return false, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if recipientJSON == nil {
		return false, fmt.Errorf("The recipient bank account %s does not exist", recipientId)
	}

	var account BankAccount
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		return false, err
	}

	var recipient BankAccount
	err = json.Unmarshal(recipientJSON, &recipient)
	if err != nil {
		return false, err
	}

	hasSameCurrency := account.Currency == recipient.Currency

	return hasSameCurrency, nil
}

// Transfer funds
func (s *SmartContract) TransferFunds(ctx contractapi.TransactionContextInterface, fromAccountId, toAccountId string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("The transfer amount must be positive")
	}

	fromAccountBytes, err := ctx.GetStub().GetState(fromAccountId)
	if err != nil {
		return fmt.Errorf("Failed to read from world state: %v", err)
	}
	if fromAccountBytes == nil {
		return fmt.Errorf("Source bank account %s does not exist", fromAccountId)
	}

	var fromAccount BankAccount
	err = json.Unmarshal(fromAccountBytes, &fromAccount)
	if err != nil {
		return err
	}

	// check source account
	if fromAccount.Balance < amount {
		return fmt.Errorf("The source bank account balance is insufficient")
	}

	toAccountBytes, err := ctx.GetStub().GetState(toAccountId)
	if err != nil {
		return fmt.Errorf("Failed to read from world state: %v", err)
	}
	if toAccountBytes == nil {
		return fmt.Errorf("The destination bank account %s does not exist", toAccountId)
	}

	var toAccount BankAccount
	err = json.Unmarshal(toAccountBytes, &toAccount)
	if err != nil {
		return err
	}

	if fromAccount.Currency != toAccount.Currency {
		convertedAmount, err := ConvertCurrency(amount, fromAccount.Currency, toAccount.Currency)
		if err != nil {
			return err
		}
		amount *= convertedAmount
	}

	// change balances
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	fromAccountBytes, err = json.Marshal(fromAccount)
	if err != nil {
		return err
	}
	toAccountBytes, err = json.Marshal(toAccount)
	if err != nil {
		return err
	}

	// update accounts
	err = ctx.GetStub().PutState(fromAccountId, fromAccountBytes)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(toAccountId, toAccountBytes)
	if err != nil {
		return err
	}

	return nil
}

// Withdraw funds
func (s *SmartContract) WithdrawFunds(ctx contractapi.TransactionContextInterface, accountId string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdrawal amount must be a positive number")
	}

	accountJSON, err := ctx.GetStub().GetState(accountId)
	if err != nil {
		return fmt.Errorf("Failed to get the account: %s", err.Error())
	}
	if accountJSON == nil {
		return fmt.Errorf("The bank account not found")
	}

	var account BankAccount
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		return err
	}

	// check balance
	if account.Balance < amount {
		return fmt.Errorf("The bank account doesn't have enough balance")
	}

	account.Balance -= amount

	accountJSON, err = json.Marshal(account)
	if err != nil {
		return err
	}

	// update account
	return ctx.GetStub().PutState(accountId, accountJSON)
}

// Deposit funds into an account
func (s *SmartContract) DepositFunds(ctx contractapi.TransactionContextInterface, accountId, currency string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Deposit amount must be positive")
	}

	accountBytes, err := ctx.GetStub().GetState(accountId)
	if err != nil {
		return fmt.Errorf("Failed to read from world state: %v", err)
	}
	if accountBytes == nil {
		return fmt.Errorf("The account %s does not exist", accountId)
	}

	var account BankAccount
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		return err
	}

	if account.Currency != currency {
		return fmt.Errorf("Can't deposit: account currency is %s but deposit is in %s", account.Currency, currency)
	}

	account.Balance += amount

	updatedAccountBytes, err := json.Marshal(account)
	if err != nil {
		return err
	}

	// updated account
	err = ctx.GetStub().PutState(accountId, updatedAccountBytes)
	if err != nil {
		return err
	}

	return nil
}
