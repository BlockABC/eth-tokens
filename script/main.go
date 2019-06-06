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

	if *nft { //构建erc721代币信息
		if *build { //从ether scan上拉取排名靠前的50个nft代币的地址，name，symbol，logo url以及说明信息，并从链上拿到name和symbol来更正ether scan的错误信息
			spider, err := ether_scan.Initialize("https://mainnet.infura.io") //"https://mainnet.infura.io" "http://47.52.157.31:8585"
			if err != nil {
				panic(err)
			}
			tokenList, err := spider.BuiltNftFromEtherScan()
			if err != nil {
				panic(err)
			}
			fmt.Println("success get token list:", len(tokenList))
			//用上面拿到的数据构建仓库，写入json文件并下载logo到指定目录
			if err := built.InitializeTokens(`../../nft`, tokenList, false); err != nil {
				panic(err)
			}
		}
		tokens, err := built.CollectTokens(`../../nft`) //从仓库中取出所有代币信息
		if err != nil {
			panic(err)
		}
		if err := built.BuildDist(`../../dist/nft.json`, tokens); err != nil { //用仓库中代币信息构建给前端使用的json文件
			panic(err)
		}
		if err := built.BuildReadme(`../../nft.md`, built.Erc721Path, tokens); err != nil { //构建readme文件，方便查看当前仓库内代币信息
			panic(err)
		}
	}
	if *erc20 { //构建erc20代币信息
		if *g { //从my easy wallet仓库拉取初始数据，不过他们仓库很多垃圾币，因此暂时不用，而是从ether scan拉取排名1000的代币
			tokenList, err := built.TokenListFromGit(built.URLTokenList)
			if err != nil {
				panic(err)
			}
			fmt.Println("success get token list:", len(tokenList))
			if err := built.InitializeTokens(`../../tokens`, tokenList, false); err != nil {
				panic(err)
			}
		}
		if *e { //从ether scan上拉取排名靠前的1000个erc20代币的地址，name，symbol，logo url以及说明信息，并从链上拿到name和symbol来更正ether scan的错误信息
			spider, err := ether_scan.Initialize("https://mainnet.infura.io")
			if err != nil {
				panic(err)
			}

			tokenList, err := spider.BuiltTokensFromEtherScan()
			if err != nil {
				panic(err)
			}
			//ether scan的logo图片非常不清楚，因此从币虎拿到更清楚的logo地址，并更新没有name和symbol的代币
			for index := range tokenList {
				coin_gecko.ReplaceTokenLogoFromCoinGecko(&tokenList[index])
			}
			fmt.Println("success get token list:", len(tokenList))
			eth := built.TokenInfo{
				Symbol:     "ETH",
				Name:       "Ethereum",
				Type:       "",
				Address:    "0x0000000000000000000000000000000000000000",
				EnsAddress: "",
				Decimals:   18,
				Website:    "",
				Logo:       built.Logo{
					Src:      "https://www.cryptocompare.com/media/20646/eth_logo.png?width=200",
					Width:    nil,
					Height:   nil,
					IpfsHash: "",
				},
				Support:    built.Support{},
				Social:     built.Social{},
			}
			tokenList = append([]built.TokenInfo{eth}, tokenList...)
			//用上面拿到的数据构建仓库，写入json文件并下载logo到指定目录
			if err := built.InitializeTokens(`../../tokens`, tokenList, false); err != nil {
				panic(err)
			}
		}

		if *logo { //从ether scan拉取代币的logo，不过ether的logo不清晰，因此暂时不用此段代码
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

		tokens, err := built.CollectTokens(`../../tokens`) //从仓库中取出所有代币信息
		if err != nil {
			panic(err)
		}

		//前端要求ETH必须是第一位，因此需要加上ETH做为ERC20代币，用地址区分

		if err := built.BuildDist(`../../dist/tokens.json`, tokens); err != nil { //用仓库中代币信息构建给前端使用的json文件
			panic(err)
		}
		if err := built.BuildReadme(`../../tokens.md`, built.Erc20Path, tokens); err != nil { //构建readme文件，方便查看当前仓库内代币信息
			panic(err)
		}
	}
}
