## highload-wallet-api

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tonuniverse/highload-wallet-api)
![GitHub](https://img.shields.io/github/license/tonuniverse/highload-wallet-api)

API wrapper over high-load TON wallet smart contract. Can be useful for cryptocurrency exchanges or any services where mass payments in TON coins are required.

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

Activate your wallet. You will need to specify the jsonRPC url, you can use public url depending on the type of network or your own local TON blockchain jsonRPC.

Mainnet: https://toncenter.com/api/v2/jsonRPC  
Testnet: https://testnet.toncenter.com/api/v2/jsonRPC

- `apt install curl`
- `./activate-wallet.sh https://toncenter.com/api/v2/jsonRPC`

Build the server

- go to project root directory
- `go build`

Create `config.json` in the project root directory

```json
{
    "server": {
        "host": "127.0.0.1",
        "port": "8080",
        "prefork": true
    },
    "ton_net": {
        "json_rpc_url": "https://toncenter.com/api/v2/jsonRPC"
    },
    "fift": {
        "path": "/usr/src/ton/crypto/fift/lib:/usr/src/ton/crypto/smartcont",
        "binary": "/usr/bin/ton/crypto/fift"
    },
    "contract": {
        "new_order_fif": "contract/new-order.fif"
    },
    "temp_path": {
        "orders": "temp/orders",
        "bocs": "temp/bocs"
    }, 
    "wallet": {
        "path": "contract/generated",
        "name": "new-wallet",
        "subwallet_id": "1"
    }
}
```

Run the server:

- `./highload-wallet-api`

Send POST request to `/transfer` endpoint with JSON data:

- transfer_tasks: array; length must be <= 100 elements
- - dest_address: TON base64 address
- - amount_ton: amount in TON format
- - msg: transaction comment; must be between 0 and 123 characters long

Example JSON in POST request:

```json
{
  "transfer_tasks": [
    { 
      "dest_address": "EQCD39VS5jcptHL8vMjEXrzGaRcCVYto7HUn4bpAOg8xqB2N",
      "amount_ton": "0.0001",
      "msg": "test"
    }
  ]
}
```

## API Errors

Possible errors can be found here: `src/api/errors.go`

## Recommendations

### Validating parameters

This API is designed to work locally in your infrastructure. You must validate all parameters reliably before sending them to highload-wallet-api

### JSON RPC

For stable operation of API, we recommend using your own instance of ton full node with a local JSON RPC server. 
Learn more in the `mytonctrl` documentation about it.

## Source code

You can always get the source code from the github repository page:  
https://github.com/tonuniverse/highload-wallet-api

## High-load wallet smart contract

The repository includes TON smart contract from:  
https://github.com/akifoq/highload-wallet

## LICENSE

GPL-3.0 License

The original license text can be obtained in the "LICENSE" file 
