res=$(fuxchaincli tx wasm store ./wasm/erc20/artifacts/cw_erc20-aarch64.wasm --fees 0.01fury --from captain --gas=2000000 -b block -y)
code_id=$(echo "$res" | jq '.logs[0].events[1].attributes[0].value' | sed 's/\"//g')
res=$(fuxchaincli tx wasm instantiate "$code_id" '{"decimals":10,"initial_balances":[{"address":"0x21e283524309CD7eC9B789B43F073e93e43e1B8f","amount":"100000000"}],"name":"my test token", "symbol":"MTT"}' --label test1 --admin did:fury:ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v --fees 0.001fury --from captain -b block -y)
contractAddr=$(echo "$res" | jq '.logs[0].events[0].attributes[0].value' | sed 's/\"//g')
fuxchaincli tx wasm execute "$contractAddr" '{"transfer":{"amount":"100","recipient":"0x78B63831Fb1050841DEaBE5cc785bCaA91AF3478"}}' --fees 0.001fury --from captain -b block -y

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
fuxchaincli query wasm contract-state smart "$contractAddr" '{"balance":{"address":"0x21e283524309CD7eC9B789B43F073e93e43e1B8f"}}'
fuxchaincli query wasm contract-state smart "$contractAddr" '{"balance":{"address":"0x78B63831Fb1050841DEaBE5cc785bCaA91AF3478"}}'
