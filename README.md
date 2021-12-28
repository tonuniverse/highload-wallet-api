## highload-wallet-api

## Source code

You can always get the source code from the github repository page:

https://github.com/tonuniverse/highload-wallet-api

## Getting started

First, you need to compile the FunC and Fift needed to create the wallet and interact with it. The easiest way to do this is to install `mytonctrl` in `lite` mode from https://github.com/igroman787/mytonctrl

Fift will be built automatically after `mytonctrl` installation. Now do the following to build the FunC:

- `cd /usr/bin/ton/`
- `make func`

Create the first highload TON wallet

- go to `highload-wallet-api` directory
- `cd contract`
- `./wallet.sh`

Get the wallet address from `contract/generated/wallet-info.txt` and send some TON coins to it(0.1 TON will be enough)

Activate your wallet. 

You will need to specify the jsonRPC url, you can use public url depending on the type of network or your own local TON blockchain jsonRPC.

Mainnet: https://toncenter.com/api/v2/jsonRPC __
Testnet: https://testnet.toncenter.com/api/v2/jsonRPC

- `apt install curl`
- `./activate-wallet.sh https://toncenter.com/api/v2/jsonRPC`


## LICENSE

GPL-3.0 License

The original license text can be obtained in the file "LICENSE"