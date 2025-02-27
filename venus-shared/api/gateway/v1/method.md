# Sample code of curl

```bash
# <Inputs> corresponding to the value of Inputs Tag of each API
curl http://<ip>:<port>/rpc/v1 -X POST -H "Content-Type: application/json"  -H "Authorization: Bearer <token>"  -d '{"method": "Gateway.<method>", "params": <Inputs>, "id": 0}'
```
# Groups

* [Gateway](#gateway)
  * [Version](#version)
* [MarketClient](#marketclient)
  * [IsUnsealed](#isunsealed)
  * [ListMarketConnectionsState](#listmarketconnectionsstate)
  * [SectorsUnsealPiece](#sectorsunsealpiece)
* [MarketServiceProvider](#marketserviceprovider)
  * [ListenMarketEvent](#listenmarketevent)
  * [ResponseMarketEvent](#responsemarketevent)
* [ProofClient](#proofclient)
  * [ComputeProof](#computeproof)
  * [ListConnectedMiners](#listconnectedminers)
  * [ListMinerConnection](#listminerconnection)
* [ProofServiceProvider](#proofserviceprovider)
  * [ListenProofEvent](#listenproofevent)
  * [ResponseProofEvent](#responseproofevent)
* [WalletClient](#walletclient)
  * [ListWalletInfo](#listwalletinfo)
  * [ListWalletInfoByWallet](#listwalletinfobywallet)
  * [WalletHas](#wallethas)
  * [WalletSign](#walletsign)
* [WalletServiceProvider](#walletserviceprovider)
  * [AddNewAddress](#addnewaddress)
  * [ListenWalletEvent](#listenwalletevent)
  * [RemoveAddress](#removeaddress)
  * [ResponseWalletEvent](#responsewalletevent)
  * [SupportNewAccount](#supportnewaccount)

## Gateway

### Version
Version provides information about API provider


Perms: read

Inputs: `[]`

Response:
```json
{
  "Version": "string value",
  "APIVersion": 131840
}
```

## MarketClient

### IsUnsealed


Perms: admin

Inputs:
```json
[
  "f01234",
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  {
    "ID": {
      "Miner": 1000,
      "Number": 9
    },
    "ProofType": 8
  },
  10,
  1032
]
```

Response: `true`

### ListMarketConnectionsState


Perms: admin

Inputs: `[]`

Response:
```json
[
  {
    "Addr": "f01234",
    "Conn": {
      "Connections": [
        {
          "Addrs": [
            "f01234"
          ],
          "ChannelId": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
          "Ip": "string value",
          "RequestCount": 123,
          "CreateTime": "0001-01-01T00:00:00Z"
        }
      ],
      "ConnectionCount": 123
    }
  }
]
```

### SectorsUnsealPiece


Perms: admin

Inputs:
```json
[
  "f01234",
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  {
    "ID": {
      "Miner": 1000,
      "Number": 9
    },
    "ProofType": 8
  },
  10,
  1032,
  "string value"
]
```

Response: `{}`

## MarketServiceProvider

### ListenMarketEvent


Perms: read

Inputs:
```json
[
  {
    "Miner": "f01234"
  }
]
```

Response:
```json
{
  "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  "Method": "string value",
  "Payload": "Ynl0ZSBhcnJheQ=="
}
```

### ResponseMarketEvent


Perms: read

Inputs:
```json
[
  {
    "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
    "Payload": "Ynl0ZSBhcnJheQ==",
    "Error": "string value"
  }
]
```

Response: `{}`

## ProofClient

### ComputeProof


Perms: admin

Inputs:
```json
[
  "f01234",
  [
    {
      "SealProof": 8,
      "SectorNumber": 9,
      "SectorKey": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "SealedCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      }
    }
  ],
  "Bw==",
  10101,
  24
]
```

Response:
```json
[
  {
    "PoStProof": 8,
    "ProofBytes": "Ynl0ZSBhcnJheQ=="
  }
]
```

### ListConnectedMiners


Perms: admin

Inputs: `[]`

Response:
```json
[
  "f01234"
]
```

### ListMinerConnection


Perms: admin

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
{
  "Connections": [
    {
      "Addrs": [
        "f01234"
      ],
      "ChannelId": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
      "Ip": "string value",
      "RequestCount": 123,
      "CreateTime": "0001-01-01T00:00:00Z"
    }
  ],
  "ConnectionCount": 123
}
```

## ProofServiceProvider

### ListenProofEvent


Perms: read

Inputs:
```json
[
  {
    "MinerAddress": "f01234"
  }
]
```

Response:
```json
{
  "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  "Method": "string value",
  "Payload": "Ynl0ZSBhcnJheQ=="
}
```

### ResponseProofEvent


Perms: read

Inputs:
```json
[
  {
    "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
    "Payload": "Ynl0ZSBhcnJheQ==",
    "Error": "string value"
  }
]
```

Response: `{}`

## WalletClient

### ListWalletInfo


Perms: admin

Inputs: `[]`

Response:
```json
[
  {
    "Account": "string value",
    "SupportAccounts": [
      "string value"
    ],
    "ConnectStates": [
      {
        "Addrs": [
          "f01234"
        ],
        "ChannelId": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
        "Ip": "string value",
        "RequestCount": 123,
        "CreateTime": "0001-01-01T00:00:00Z"
      }
    ]
  }
]
```

### ListWalletInfoByWallet


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response:
```json
{
  "Account": "string value",
  "SupportAccounts": [
    "string value"
  ],
  "ConnectStates": [
    {
      "Addrs": [
        "f01234"
      ],
      "ChannelId": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
      "Ip": "string value",
      "RequestCount": 123,
      "CreateTime": "0001-01-01T00:00:00Z"
    }
  ]
}
```

### WalletHas


Perms: admin

Inputs:
```json
[
  "string value",
  "f01234"
]
```

Response: `true`

### WalletSign


Perms: admin

Inputs:
```json
[
  "string value",
  "f01234",
  "Ynl0ZSBhcnJheQ==",
  {
    "Type": "message",
    "Extra": "Ynl0ZSBhcnJheQ=="
  }
]
```

Response:
```json
{
  "Type": 2,
  "Data": "Ynl0ZSBhcnJheQ=="
}
```

## WalletServiceProvider

### AddNewAddress


Perms: read

Inputs:
```json
[
  "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  [
    "f01234"
  ]
]
```

Response: `{}`

### ListenWalletEvent


Perms: read

Inputs:
```json
[
  {
    "SupportAccounts": [
      "string value"
    ],
    "SignBytes": "Ynl0ZSBhcnJheQ=="
  }
]
```

Response:
```json
{
  "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  "Method": "string value",
  "Payload": "Ynl0ZSBhcnJheQ=="
}
```

### RemoveAddress


Perms: read

Inputs:
```json
[
  "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  [
    "f01234"
  ]
]
```

Response: `{}`

### ResponseWalletEvent


Perms: read

Inputs:
```json
[
  {
    "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
    "Payload": "Ynl0ZSBhcnJheQ==",
    "Error": "string value"
  }
]
```

Response: `{}`

### SupportNewAccount


Perms: read

Inputs:
```json
[
  "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  "string value"
]
```

Response: `{}`

