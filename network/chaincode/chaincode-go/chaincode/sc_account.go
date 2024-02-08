package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) GetBankAccount(ctx contractapi.TransactionContextInterface, id int64) (*BankAccount, error) {
	bankAccountJSON, err := s.GetEntityById(ctx, BANK_ACCOUNT_TYPE_NAME, id)

	if err != nil || bankAccountJSON == nil {
		return nil, fmt.Errorf("Bank account with given id %d doesn't exist", id)
	}

	var bankAccount BankAccount
	err = json.Unmarshal(bankAccountJSON, &bankAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bank account: %v", err)
	}

	return &bankAccount, nil
}

// Create bank account
func (s *SmartContract) CreateBankAccount(ctx contractapi.TransactionContextInterface, id int64, personId int64, bankId int64, currency string, balance float64) (*BankAccount, error) {
	exists, err := s.EntityExists(ctx, BANK_ACCOUNT_TYPE_NAME, id)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("Bank account with given id %d already exists", id)
	}

	bankExists, err := s.EntityExists(ctx, BANK_TYPE_NAME, bankId)
	if err != nil {
		return nil, err
	}

	if !bankExists {
		return nil, fmt.Errorf("Bank with given id %d doesn't exist", bankId)
	}

	accountId := ToBankAccountId(id)
	account := BankAccount{
		Id:       accountId,
		PersonId: ToPersonId(personId),
		BankId:   ToBankId(bankId),
		Balance:  balance,
		Currency: currency,
		Cards:    make([]Card, 0),
	}

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}
	err = ctx.GetStub().PutState(account.Id, accountJSON)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *SmartContract) CheckAccountCurrencies(ctx contractapi.TransactionContextInterface, fromAccountId int64, recipientId int64) (bool, error) {
	account, err := s.GetBankAccount(ctx, fromAccountId)
	if err != nil {
		return false, err
	}

	recipient, err := s.GetBankAccount(ctx, recipientId)
	if err != nil {
		return false, err
	}

	hasSameCurrency := account.Currency == recipient.Currency

	return hasSameCurrency, nil
}

// Transfer funds
func (s *SmartContract) TransferFunds(ctx contractapi.TransactionContextInterface, fromAccountId int64, toAccountId int64, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("the transfer amount must be positive")
	}

	fromAccount, err := s.GetBankAccount(ctx, fromAccountId)
	if err != nil {
		return err
	}

	if fromAccount == nil {
		return fmt.Errorf("FromAccount does not exist")
	}

	toAccount, err := s.GetBankAccount(ctx, toAccountId)
	if err != nil {
		return err
	}

	if toAccount == nil {
		return fmt.Errorf("ToAccount does not exist")
	}

	// check source account
	if fromAccount.Balance < amount {
		return fmt.Errorf("the source bank account balance is insufficient")
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

	fromAccountJSON, err := json.Marshal(*fromAccount)
	if err != nil {
		return err
	}
	toAccountJSON, err := json.Marshal(*toAccount)
	if err != nil {
		return err
	}

	// update accounts
	err = ctx.GetStub().PutState(fromAccount.Id, fromAccountJSON)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(toAccount.Id, toAccountJSON)
	if err != nil {
		return err
	}

	return nil
}

// Withdraw funds
func (s *SmartContract) WithdrawFunds(ctx contractapi.TransactionContextInterface, accountId int64, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be a positive number")
	}

	bankAccount, err := s.GetBankAccount(ctx, accountId)
	if err != nil || bankAccount == nil {
		return fmt.Errorf("failed to find bank account")
	}

	// check balance
	if bankAccount.Balance < amount {
		return fmt.Errorf("the bank account doesn't have enough balance")
	}

	bankAccount.Balance -= amount

	bankAccountJSON, err := json.Marshal(*bankAccount)
	if err != nil {
		return err
	}

	// update account
	return ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
}

// Deposit funds into an account
func (s *SmartContract) DepositFunds(ctx contractapi.TransactionContextInterface, accountId int64, currency string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}

	bankAccount, err := s.GetBankAccount(ctx, accountId)
	if err != nil || bankAccount == nil {
		return fmt.Errorf("failed to find bank account")
	}

	if bankAccount.Currency != currency {
		return fmt.Errorf("can't deposit: account currency is %s but deposit is in %s", bankAccount.Currency, currency)
	}

	bankAccount.Balance += amount

	bankAccountJSON, err := json.Marshal(*bankAccount)
	if err != nil {
		return err
	}

	// updated account
	err = ctx.GetStub().PutState(bankAccount.Id, bankAccountJSON)
	if err != nil {
		return err
	}

	return nil
}
