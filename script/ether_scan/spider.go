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
		name, symbol, decimals, _, err := erc20.ReadTokenInfo(token.Address, s.client, s.url)
		if err != nil {
			log.Error("read token info err:", token.Address, err)
			continue
		}
		tokens[i].Name = name
		tokens[i].Symbol = symbol
		tokens[i].Decimals = int(decimals)
		log.Debug("get token:", name, symbol, decimals, token.Address)
	}

	return tokens, nil
}

func (s *Spider) BuiltNftFromEtherScan() ([]built.TokenInfo, error) {
	var tokens []built.TokenInfo
	for i := 1; i <= pageMax; i++ {
	retry:
		log.Debug("get the token from ether scan:", i)
		ts, err := RequestTokenListByPage(urlEtherScan + fmt.Sprintf("%d", i))
		if err != nil {
			log.Error("RequestTokenListByPage error:", err)
			goto retry
		}
		tokens = append(tokens, ts...)
		time.Sleep(time.Millisecond * 500)
	}
	log.Debug("get tokens from ether scan:", len(tokens))
	for i, token := range tokens {
		name, symbol, decimals, _, err := erc20.ReadTokenInfo(token.Address, s.client, s.url)
		if err != nil {
			log.Error("read token info err:", token.Address, err)
			continue
		}
		tokens[i].Name = name
		tokens[i].Symbol = symbol
		tokens[i].Decimals = 1
		log.Debug("get token:", name, symbol, decimals, token.Address)
	}
	return tokens, nil
}
