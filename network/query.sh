#!/bin/bash

if [ $# -lt 1 ]; then
    echo "Usage: ./query.sh <function> [args...]"
    exit 1
fi

# Set env vars
source env-vars.sh

# Get arguments
ARGS="\"$1\""
for i in ${@:2}; do
    ARGS="$ARGS,\"$i\""
done
ARGS="{\"Args\":[$ARGS]}"
echo $ARGS

# Query chaincode
peer chaincode query -C mychannel -n basic -c "$ARGS"