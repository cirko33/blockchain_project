cd ./network
./network.sh down
./network.sh up 
./network.sh createChannel -c channel1
./network.sh createChannel -c channel2
./network.sh deployCC -c channel1 -ccn basic1 -ccp ./chaincode/chaincode-go -ccl go
./network.sh deployCC -c channel2 -ccn basic2 -ccp ./chaincode/chaincode-go -ccl go

if [[ $1 == "-i" ]]; then
    ./invoke.sh 1 InitLedger
    ./invoke.sh 2 InitLedger
fi
# read -p "Press any key to continue... " -n1 -s
# ./network.sh down