#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-fuxchaind.log}

##
## Run binary with all parameters
##
export FUXCHAINDHOME="/fuxchaind/node${ID}/fuxchaind"

if [ -d "$(dirname "${FUXCHAINDHOME}"/"${LOG}")" ]; then
  fuxchaind --chain-id fuxchain-1 --home "${FUXCHAINDHOME}" "$@" | tee "${FUXCHAINDHOME}/${LOG}"
else
  fuxchaind --chain-id fuxchain-1 --home "${FUXCHAINDHOME}" "$@"
fi

