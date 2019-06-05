package main

import (
	"flag"
	"fmt"
	"github.com/eager7/eth_tokens/script/built"
	"github.com/eager7/eth_tokens/script/coin_gecko"
	"github.com/eager7/eth_tokens/script/ether_scan"
	"time"
)

func main() {
	var g = flag.Bool("g", false, "")
	var e = flag.Bool("e", false, "")
	var logo = flag.Bool("logo", false, "")
	var nft = flag.Bool("nft", false, "")
	var erc20 = flag.Bool("erc20", false, "")
	var build = flag.Bool("build", false, "")
	flag.Parse()

	if *nft {
		if *build {
			spider, err := ether_scan.Initialize("https://mainnet.infura.io")//"https://mainnet.infura.io" "http://47.52.157.31:8585"
			if err != nil {
				panic(err)
			}
			tokenList, err := spider.BuiltNftFromEtherScan()
			if err != nil {
				panic(err)
			}
			fmt.Println("success get token list:", len(tokenList))
			if err := built.InitializeTokens(`../../nft`, tokenList, false); err != nil {
				panic(err)
			}
		}
		tokens, err := built.CollectTokens(`../../nft`)
		if err != nil {
			panic(err)
		}
		if err := built.BuildDist(`../../dist/nft.json`, tokens); err != nil {
			panic(err)
		}
		if err := built.BuildReadme(`../../nft.md`, tokens); err != nil {
			panic(err)
		}
	}
	if *erc20 {
		if *g {
			tokenList, err := built.TokenListFromGit(built.URLTokenList)
			if err != nil {
				panic(err)
			}
			fmt.Println("success get token list:", len(tokenList))
			if err := built.InitializeTokens(`../../tokens`, tokenList, false); err != nil {
				panic(err)
			}
		}
		if *e {
			spider, err := ether_scan.Initialize("http://47.52.157.31:8585")
			if err != nil {
				panic(err)
			}
			tokenList, err := spider.BuiltTokensFromEtherScan()
			if err != nil {
				panic(err)
			}
			for index := range tokenList {
				coin_gecko.ReplaceTokenLogoFromCoinGecko(&tokenList[index])
			}
			fmt.Println("success get token list:", len(tokenList))
			if err := built.InitializeTokens(`../../tokens`, tokenList, false); err != nil {
				panic(err)
			}
		}

		if *logo {
			tokens, err := built.CollectTokens(`../../tokens`)
			if err != nil {
				panic(err)
			}
			fmt.Println("total len of token:", len(tokens))
			var tokenList []built.TokenInfo
			for _, t := range tokens {
				if t.Logo == "" {
					tokenList = append(tokenList, built.TokenInfo{
						Symbol:     t.Symbol,
						Name:       t.Name,
						Type:       "",
						Address:    t.Contract,
						EnsAddress: "",
						Decimals:   t.Decimals,
						Logo:       built.Logo{},
					})
				}
			}
			fmt.Println("no logo token list:", len(tokenList))
			for i, t := range tokenList {
			retry:
				if err := ether_scan.RequestTokenLogo(&tokenList[i]); err != nil {
					fmt.Println("request err:", t.Symbol, t.Address)
					time.Sleep(time.Second * 1)
					goto retry
				}
			}
			if err := built.InitializeTokens(`../../tokens`, tokenList, false); err != nil {
				panic(err)
			}
		}

		tokens, err := built.CollectTokens(`../../tokens`)
		if err != nil {
			panic(err)
		}

		eth := built.Token{
			Name:     "Ethereum",
			Symbol:   "ETH",
			Contract: "0x0000000000000000000000000000000000000000",
			Decimals: 18,
			Logo:     "https://www.cryptocompare.com/media/20646/eth_logo.png?width=200",
			Invalid:  true,
		}
		tokens = append([]*built.Token{&eth}, tokens...)
		if err := built.BuildDist(`../../dist/tokens.json`, tokens); err != nil {
			panic(err)
		}
		if err := built.BuildReadme(`../../tokens.md`, tokens); err != nil {
			panic(err)
		}
	}
}
