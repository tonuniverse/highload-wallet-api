RESULT_DIR="generated"
FC_FILE="highload-wallet-v2-code.fc"

mkdir $RESULT_DIR
cp $FC_FILE $RESULT_DIR

FULL_DIR="`pwd`/$RESULT_DIR"

/usr/bin/ton/compile.sh $FULL_DIR 0
rm $RESULT_DIR/$FC_FILE

FIFT_FILE="new-highload-wallet-v2.fif"
FIFT_PATH="/usr/bin/ton/crypto/fift"

cp $FIFT_FILE $RESULT_DIR

cd $RESULT_DIR

${FIFT_PATH} -s ${FIFT_FILE} 0 1 >> res.txt
rm ${FIFT_FILE}
rm "highload-wallet-v2-code.fif"

echo "Highload wallet address: 

-- Non-bounceable: `awk '/Non-bounceable address/{print $NF}' res.txt`
-- Bounceable:     `awk '/Bounceable address/{print $NF}' res.txt`
-- Raw:            `awk '/new wallet address/{print $NF}' res.txt`

Some more info:

-- Init query_id is:  `awk '/Init query_id/{print $NF}' res.txt`
-- Signing message:   `awk '/signing message/{print $NF}' res.txt`
" >> h.txt

rm res.txt