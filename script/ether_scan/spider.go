package ether_scan

import (
	"errors"
	"fmt"
	"github.com/eager7/elog"
	"github.com/eager7/eth_tokens/script/built"
	"github.com/eager7/eth_tokens/script/erc20"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

var log = elog.NewLogger("spider", elog.DebugLevel)

type Spider struct {
	url    string
	client *ethclient.Client
}

func Initialize(url string) (*Spider, error) { //"http://47.52.157.31:8585"
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("initialize eth client err:%v", err))
	}
	return &Spider{url: url, client: client}, nil
}

func (s *Spider) BuiltTokensFromEtherScan() ([]built.TokenInfo, error) {
	var tokens []built.TokenInfo
	for i := 1; i <= pageMax; i++ {
	retry:
		log.Debug("get the token from ether scan:", i)
		ts, err := RequestTokenListByPage(urlNftEtherScan + fmt.Sprintf("%d", i))
		if err != nil {
			log.Error("RequestTokenListByPage error:", err)
			goto retry
		}
		tokens = append(tokens, ts...)
		time.Sleep(time.Millisecond * 500)
	}
	log.Debug("get tokens from ether scan:", len(tokens))
	for i, token := range tokens {
		times := 0
	read:
		name, symbol, decimals, _, err := erc20.ReadTokenInfo(token.Address, s.client, s.url)
		if err != nil {
			times++
			log.Error("read token info err:", token.Address, err)
			if times <= 3 {
				goto read
			}
			continue
		}
		if name != "" {
			tokens[i].Name = name
		}
		if symbol != "" {
			tokens[i].Symbol = symbol
		}
		tokens[i].Type = "ERC20"
		tokens[i].Decimals = int(decimals)
		log.Debug("get token:", tokens[i].Name, tokens[i].Symbol, tokens[i].Decimals, tokens[i].Type, token.Address)
	}

	return tokens, nil
}

func (s *Spider) BuiltNftFromEtherScan() ([]built.TokenInfo, error) {
	var tokens []built.TokenInfo
	for i := 1; i <= 1; i++ {
	retry:
		log.Debug("get the token from ether scan:", i)
		ts, err := RequestTokenListByPage(urlNftEtherScan + fmt.Sprintf("%d", i))
		if err != nil {
			log.Error("RequestTokenListByPage error:", err)
			goto retry
		}
		tokens = append(tokens, ts...)
		time.Sleep(time.Millisecond * 500)
	}
	log.Debug("get tokens from ether scan:", len(tokens))
	for i, token := range tokens {
		name, symbol, _, _, err := erc20.ReadTokenInfo(token.Address, s.client, s.url)
		if err != nil {
			log.Error("read token info err:", token.Address, err)
			continue
		}
		if name != "" {
			tokens[i].Name = name
		}
		if symbol != "" {
			tokens[i].Symbol = symbol
		}
		tokens[i].Decimals = 1
		tokens[i].Type = "ERC721"
		log.Debug("get token:", tokens[i].Name, tokens[i].Symbol, tokens[i].Decimals, tokens[i].Type, token.Address)
	}
	return tokens, nil
}
