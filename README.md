# ETH tokens
> The ultimate collection of all ETH tokens (PR welcome !).

## Why

There are over 184,137 tokens on ETH mainnet (According to [EtherScan Tokens](https://etherscan.io/tokens)).

Unfortunately, it's hard for developers to collect all the tokens' logo and other info. Thus limited the chance to provide a better user experience for users.

This lib aims to collect all the tokens and their data including their **logo/decimals/desc/website, etc.**

## Contributing
PR is welcome!

0. Star this repository

1. Create a folder under [/tokens](/tokens) named with token's contract address(lowercase) `/tokens/${CONTRACT_ADDRESS}`

2. Then add yor token LOGO and JSON file under the folder. `/tokens/${CONTRACT_ADDRESS}/{TOKEN_UPPERCASE}.png` `/tokens/${CONTRACT_ADDRESS}/{TOKEN_UPPERCASE}.json`

3. Create PR, and leave your contact (Telegram/WeChat/Email) in comment for further support.

Please refer to [BNB](https://github.com/eager7/eth_tokens/tree/master/tokens/0xB8c77482e45F1F44dE1745F52C74426C631bDD52) for token example.

Feel free to submit tokens if you are the token's owner or not, the community needs your contribution :).

Thanks to [MyEtherWallet](https://github.com/MyEtherWallet/ethereum-lists) for initial data.

## Template
Please submit your token's JSON file follow template below. [BNB](./tokens/0xB8c77482e45F1F44dE1745F52C74426C631bDD52)

```json
{
    "name": "BNB",
    "symbol": "BNB",
    "contract": "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
    "decimals": 18,
    "logo": "https://raw.githubusercontent.com/eager7/eth_tokens/master/tokens/0xB8c77482e45F1F44dE1745F52C74426C631bDD52/BNB.png",
    "desc": {
        "en": "",
        "zh": ""
    },
    "website": "https://www.binance.com/cn",
    "whitepaper": "",
    "invalid": true,
    "links": {
        "twitter": "",
        "telegram": ""
    }
}
```

## Note
1. For better compatibility, please provide PNG format as possible.
2. `invalid` field is only used to indicate that this token is invalid on chain. If your token is valid, please don't fill in this field.

## Token List
 [tokens.md](./tokens.md)
