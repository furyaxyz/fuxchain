#!/bin/bash
set -o errexit -o nounset -o pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

fuxchaincli keys add fred
echo "0-----------------------"
fuxchaincli tx send captain $(fuxchaincli keys show fred -a) 1000okb --fees 0.001okb -y -b block

echo "1-----------------------"
echo "## Add new CosmWasm contract"
RESP=$(fuxchaincli tx wasm store "$DIR/../../../x/wasm/keeper/testdata/hackatom.wasm" \
  --from captain --fees 0.001okb --gas 1500000 -y --node=http://localhost:26657 -b block -o json)

CODE_ID=$(echo "$RESP" | jq -r '.logs[0].events[1].attributes[-1].value')
echo "* Code id: $CODE_ID"
echo "* Download code"
TMPDIR=$(mktemp -t fuxchaincliXXXXXX)
fuxchaincli q wasm code "$CODE_ID" "$TMPDIR"
rm -f "$TMPDIR"
echo "-----------------------"
echo "## List code"
fuxchaincli query wasm list-code --node=http://localhost:26657 -o json | jq

echo "2-----------------------"
echo "## Create new contract instance"
INIT="{\"verifier\":\"$(fuxchaincli keys show captain -a)\", \"beneficiary\":\"$(fuxchaincli keys show fred -a)\"}"
fuxchaincli tx wasm instantiate "$CODE_ID" "$INIT" --admin="$(fuxchaincli keys show captain -a)" \
  --from captain  --fees 0.001okb --amount="100okb" --label "local0.1.0" \
  --gas 1000000 -y -b block -o json | jq

CONTRACT=$(fuxchaincli query wasm list-contract-by-code "$CODE_ID" -o json | jq -r '.contracts[-1]')
echo "* Contract address: $CONTRACT"
echo "### Query all"
RESP=$(fuxchaincli query wasm contract-state all "$CONTRACT" -o json)
echo "$RESP" | jq
echo "### Query smart"
fuxchaincli query wasm contract-state smart "$CONTRACT" '{"verifier":{}}' -o json | jq
echo "### Query raw"
KEY=$(echo "$RESP" | jq -r ".models[0].key")
fuxchaincli query wasm contract-state raw "$CONTRACT" "$KEY" -o json | jq

echo "3-----------------------"
echo "## Execute contract $CONTRACT"
MSG='{"release":{}}'
fuxchaincli tx wasm execute "$CONTRACT" "$MSG" \
  --from captain \
  --gas 1000000 --fees 0.001okb -y  -b block -o json | jq

echo "4-----------------------"
echo "## Set new admin"
echo "### Query old admin: $(fuxchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
echo "### Update contract"
fuxchaincli tx wasm set-contract-admin "$CONTRACT" "$(fuxchaincli keys show fred -a)" \
  --from captain --fees 0.001okb -y -b block -o json | jq
echo "### Query new admin: $(fuxchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"

echo "5-----------------------"
echo "## Migrate contract"
echo "### Upload new code"
RESP=$(fuxchaincli tx wasm store "$DIR/../../../x/wasm/keeper/testdata/burner.wasm" \
  --from captain --fees 0.001okb --gas 1000000 -y --node=http://localhost:26657 -b block -o json)

BURNER_CODE_ID=$(echo "$RESP" | jq -r '.logs[0].events[1].attributes[-1].value')
echo "### Migrate to code id: $BURNER_CODE_ID"

DEST_ACCOUNT=$(fuxchaincli keys show fred -a)
fuxchaincli tx wasm migrate "$CONTRACT" "$BURNER_CODE_ID" "{\"payout\": \"$DEST_ACCOUNT\"}" --from fred  --fees 0.001okb \
 -b block -y -o json | jq

echo "### Query destination account: $BURNER_CODE_ID"
fuxchaincli q account "$DEST_ACCOUNT" -o json | jq
echo "### Query contract meta data: $CONTRACT"
fuxchaincli q wasm contract "$CONTRACT" -o json | jq

echo "### Query contract meta history: $CONTRACT"
fuxchaincli q wasm contract-history "$CONTRACT" -o json | jq

echo "6-----------------------"
echo "## Clear contract admin"
echo "### Query old admin: $(fuxchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
echo "### Update contract"
fuxchaincli tx wasm clear-contract-admin "$CONTRACT" --fees 0.001okb \
  --from fred -y -b block -o json | jq
echo "### Query new admin: $(fuxchaincli q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
