#!/bin/bash

if [ $# -lt 2 ]; then
    echo "Usage: ./invoke.sh <channel-num> <function> [args...]"
    exit 1
fi

# Set fabric env vars
source env-vars.sh

PEER_ORG_PATH="${PWD}/organizations/peerOrganizations"

function createCommand() {
    local FUNC_NAME=$1
    local ARGS=""
    if [ $# -gt 1 ]; then
        ARGS="\"$2\""
        for i in ${@:3}; do
            ARGS="$ARGS,\"$i\""
        done
    fi

    COMMAND="{\"function\":\"$FUNC_NAME\",\"Args\":[$ARGS]}"
}

function createPeer0Connections() {
    for (( i=1; i<=$ORGANIZATION_NUMBER; i++ )); do
        local ORG_PATH="${PEER_ORG_PATH}/org${i}.example.com"
        local PEER_PATH="${ORG_PATH}/peers/peer0.org${i}.example.com"
        local PEER_TLS_CERT="${PEER_PATH}/tls/ca.crt"
        local PEER_ADDRESS="localhost:$((6 + $i))051"
        PEER_CONNECTIONS="$PEER_CONNECTIONS --peerAddresses $PEER_ADDRESS --tlsRootCertFiles $PEER_TLS_CERT"
    done
}

createCommand ${@:2}
createPeer0Connections

peer chaincode invoke \
    -o localhost:6050 \
    --ordererTLSHostnameOverride orderer.example.com \
    --tls \
    --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
    -C channel$1\
    -n basic$1 \
    $PEER_CONNECTIONS \
    -c "$COMMAND"