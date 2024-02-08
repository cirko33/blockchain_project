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
invoke_function CreateBank 31 Belgrade 1234 2012
invoke_function CreateBank 32 Belgrade 1234 2013
query_function GetAllBanks
query_function GetBank 31

infoln "Testing bank accounts"
invoke_function CreateBankAccount 51 1 31 RSD 1000.0
invoke_function CreateBankAccount 52 1 31 RSD 1000.0
invoke_function CreateBankAccount 53 1 31 EUR 10.0
invoke_function CreateBankAccount 54 1 31 EUR 10.0
invoke_function CreateBankAccount 55 1 32 RSD 1000.0

query_function GetBankAccount 51
query_function GetBankAccount 52
query_function GetBankAccount 53
query_function GetBankAccount 54
query_function GetBankAccount 55

infoln "Testing bank accounts and currencies"
query_function CheckAccountCurrencies 51 52
query_function CheckAccountCurrencies 51 53
query_function CheckAccountCurrencies 53 54

infoln "Testing transfer funds"
invoke_function TransferFunds 51 52 100.0
query_function GetBankAccount 51
query_function GetBankAccount 52
invoke_function TransferFunds 51 53 100.0
query_function GetBankAccount 51
query_function GetBankAccount 53
invoke_function TransferFunds 53 54 5.0
query_function GetBankAccount 53
query_function GetBankAccount 54

infoln "Testing withdraw funds"
invoke_function WithdrawFunds 55 500.0
query_function GetBankAccount 55

infoln "Testing deposit funds"
invoke_function DepositFunds 55 RSD 1000.0
query_function GetBankAccount 55

infoln "Testing cards"
invoke_function  CreateCard 1234-4321-8765-5678 301 51
invoke_function  CreateCard 1234-4321-8765-5678 302 51
query_function GetBankAccount 51
invoke_function RemoveCard 301 51
query_function GetBankAccount 51
invoke_function RemoveCard 302 51
query_function GetBankAccount 51

infoln "Testing rich queries"
query_function CheckBankAccounts 51 53 RSD
query_function CheckBankAccounts 53 54 RSD
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