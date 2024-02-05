#!/bin/bash

# Set env vars
source env-vars.sh

# Run query
FUNC_NAME=$1
peer chaincode query -C mychannel -n basic -c "{\"Args\":[\"${FUNC_NAME}\", \"1\", \"2\"]}"