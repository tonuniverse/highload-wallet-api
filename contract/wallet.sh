#!/bin/bash

RESULT_DIR="generated"
FC_FILE="highload-wallet-v2-code.fc"
WALLET_FIF_FILE="highload-wallet-v2-code.fif"
NEW_WALLET_FIFT_FILE="new-highload-wallet-v2.fif"

fiftPath="/usr/bin/ton/crypto/fift"
smartcontPath="/usr/src/ton/crypto/smartcont"

export FIFTPATH="/usr/src/ton/crypto/fift/lib:${smartcontPath}"

mkdir $RESULT_DIR

/usr/bin/ton/crypto/func \
 -SPA "${smartcontPath}/stdlib.fc" ${FC_FILE} -o ${RESULT_DIR}/${WALLET_FIF_FILE}

cp $NEW_WALLET_FIFT_FILE $RESULT_DIR
cd $RESULT_DIR

${fiftPath} -s ${NEW_WALLET_FIFT_FILE} 0 1 >> res.txt
rm ${NEW_WALLET_FIFT_FILE} ${WALLET_FIF_FILE}

echo "Highload wallet address: 

-- Non-bounceable: `awk '/Non-bounceable address/{print $NF}' res.txt`
-- Bounceable:     `awk '/Bounceable address/{print $NF}' res.txt`
-- Raw:            `awk '/new wallet address/{print $NF}' res.txt`

More info:

-- Init query_id is:  `awk '/Init query_id/{print $NF}' res.txt`
-- Signing message:   `awk '/signing message/{print $NF}' res.txt`
" >> wallet-info.txt

rm res.txt