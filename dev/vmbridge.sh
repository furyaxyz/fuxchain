res=$(fuxchaincli tx wasm store ./wasm/erc20/artifacts/cw_erc20.wasm --fees 0.01fury --from captain --gas=20000000 -b block -y)
echo "store--------------"
echo $res
code_id=$(echo "$res" | jq '.logs[0].events[1].attributes[0].value' | sed 's/\"//g')
res=$(fuxchaincli tx wasm instantiate "$code_id" '{"decimals":10,"initial_balances":[{"address":"did:fury:ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v","amount":"100000000"}],"name":"my test token", "symbol":"MTT"}' --label test1 --admin did:fury:ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v --fees 0.001fury --from captain -b block -y)
contractAddr=$(echo "$res" | jq '.logs[0].events[0].attributes[0].value' | sed 's/\"//g')
echo $contractAddr
