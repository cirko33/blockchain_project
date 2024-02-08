. scripts/utils.sh

function query_function() {
    func=$1
    args=${@:2}
    infoln "Testing query $func"
    ./query.sh 1 $func $args | json_pp
    if [ $? -ne 0 ]; then
        warnln "Failed to query $func"
    else
        successln "Test passed"
    fi
     
    echo ""
    sleep 3
}

function invoke_function() {
    func=$1
    args=${@:2}
    infoln "Testing invoke $func"
    ./invoke.sh 1 $func $args
    if [ $? -ne 0 ]; then
        warnln "Failed to invoke $func"
    else
        successln "Test passed"
    fi
    
    echo ""
    sleep 3
}

successln "Testing chaincode on channel1"
sleep 1

infoln "Testing banks"
invoke_function CreateBank 1 Belgrade 1234 2012
invoke_function CreateBank 2 Belgrade 1234 2013
query_function GetAllBanks
query_function GetBank 1

infoln "Testing bank accounts"
invoke_function CreateBankAccount 1 1 1 RSD 1000.0
invoke_function CreateBankAccount 2 1 1 RSD 1000.0
invoke_function CreateBankAccount 3 1 2 EUR 10.0
invoke_function CreateBankAccount 4 1 2 EUR 10.0
invoke_function CreateBankAccount 5 1 2 RSD 1000.0

query_function GetBankAccount 1
query_function GetBankAccount 2
query_function GetBankAccount 3
query_function GetBankAccount 4
query_function GetBankAccount 5

infoln "Testing bank accounts and currencies"
query_function CheckAccountCurrencies 1 2
query_function CheckAccountCurrencies 1 3
query_function CheckAccountCurrencies 3 4

infoln "Testing transfer funds"
invoke_function TransferFunds 1 2 100.0
query_function GetBankAccount 1
query_function GetBankAccount 2
invoke_function TransferFunds 1 3 100.0
query_function GetBankAccount 1
query_function GetBankAccount 3
invoke_function TransferFunds 3 4 5.0
query_function GetBankAccount 3
query_function GetBankAccount 4

infoln "Testing withdraw funds"
invoke_function WithdrawFunds 5 500.0
query_function GetBankAccount 5

infoln "Testing deposit funds"
invoke_function DepositFunds 5 RSD 1000.0
query_function GetBankAccount 5

infoln "Testing cards"
invoke_function  CreateCard 1234-4321-8765-5678 1 1
invoke_function  CreateCard 1234-4321-8765-5678 2 1
query_function GetBankAccount 1
invoke_function RemoveCard 1 1
query_function GetBankAccount 1
invoke_function RemoveCard 2 1
query_function GetBankAccount 1

infoln "Testing rich queries"
query_function CheckBankAccounts 1 3 RSD
query_function CheckBankAccounts 3 4 RSD
query_function SearchPersonsByName Person_1
query_function SearchPersonsBySurname Personic_1
query_function SearchPersonsBySurnameAndEmail Personic_1 mejl_1@gmail.com
query_function GetAccountsWithMoreThanBalance RSD 100
query_function GetBanksOlderThan 2012
query_function GetBanksByLocation Location_1
query_function GetBankAccountsByPerson person-1
query_function GetBankAccountsByBank bank-1
query_function GetAllBanksWithAccounts

infoln "Testing account functions"