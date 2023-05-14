res=$(fuxchaincli tx wasm store ./wasm/erc20/artifacts/cw_erc20-aarch64.wasm --fees 0.01okb --from captain --gas=2000000 -b block -y)
code_id=$(echo "$res" | jq '.logs[0].events[1].attributes[0].value' | sed 's/\"//g')
res=$(fuxchaincli tx wasm instantiate "$code_id" '{"decimals":10,"initial_balances":[{"address":"ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v","amount":"100000000"}],"name":"my test token", "symbol":"MTT"}' --label test1 --admin ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v --fees 0.001okb --from captain -b block -y)
contractAddr=$(echo "$res" | jq '.logs[0].events[0].attributes[0].value' | sed 's/\"//g')
fuxchaincli tx wasm execute "$contractAddr" '{"transfer":{"amount":"100","recipient":"ex1eutyuqqase3eyvwe92caw8dcx5ly8s544q3hmq"}}' --fees 0.001okb --from captain -b block -y

echo " ========================================================== "
echo "## show all codes uploaded ##"
fuxchaincli query wasm list-code

echo " ========================================================== "
echo "## show contract info by contract addr ##"
fuxchaincli query wasm contract "$contractAddr"

echo " ========================================================== "
echo "## show contract update history by contract addr ##"
fuxchaincli query wasm contract-history "$contractAddr"

echo " ========================================================== "
echo "## query contract state by contract addr ##"
echo "#### all state"
fuxchaincli query wasm contract-state all "$contractAddr"
echo "#### raw state"
fuxchaincli query wasm contract-state raw "$contractAddr" 0006636F6E666967636F6E7374616E7473
echo "#### smart state"
fuxchaincli query wasm contract-state smart "$contractAddr" '{"balance":{"address":"ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v"}}'
fuxchaincli query wasm contract-state smart "$contractAddr" '{"balance":{"address":"ex1eutyuqqase3eyvwe92caw8dcx5ly8s544q3hmq"}}'
