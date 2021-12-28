#!/bin/bash

help(){
    echo "Usage: 
    ./activate-wallet.sh {jsonRPC url}"
    exit 2
}

JSON_RPC_URL=$1
BOC_FILENAME="generated/new-wallet1-query.boc"

[[ -z "$JSON_RPC_URL" ]] && help
[ ! -f $BOC_FILENAME ] && echo "${BOC_FILENAME} does not exist" &&  exit 2

B64_BOC=`base64 -w 0 generated/new-wallet1-query.boc`

printf "JSON RPC response:\n\n"

curl -X POST $JSON_RPC_URL -H 'content-type: application/json' \
    -d '{"jsonrpc":"2.0","method":"sendBoc","params":{"boc":"'${B64_BOC}'"}}'

echo ""