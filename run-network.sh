cd ./network
./network.sh down
./network.sh up createChannel
./network.sh deployCC -ccn basic -ccp ./chaincode/chaincode-go -ccl go
if [ $1 = "-i" ]; then
    ./invoke.sh InitLedger
fi
read -p "Press any key to continue... " -n1 -s
./network.sh down