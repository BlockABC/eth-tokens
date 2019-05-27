package erc20

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"unicode"
)

func BalanceAt(address, contract string, client *ethclient.Client) (*big.Int, error) {
	instance, err := NewErc20(common.HexToAddress(contract), client)
	if err != nil {
		return nil, errors.New("new token err:" + err.Error())
	}
	return instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

func PackMessage(name string, args ...interface{}) []byte {
	cAbi, err := abi.JSON(strings.NewReader(string(Erc20ABI)))
	if err != nil {
		fmt.Println("new abi instance err:", err)
		return nil
	}
	data, err := cAbi.Pack(name, args...)
	if err != nil {
		fmt.Println("abi package err:", name, err)
	}
	return data
}

func CallContract(client *ethclient.Client, contract, name string, args ...interface{}) string {
	cAbi, err := abi.JSON(strings.NewReader(string(Erc20ABI)))
	if err != nil {
		fmt.Println("new abi instance err:", err)
		return ""
	}
	data, err := cAbi.Pack(name, args...)
	if err != nil {
		fmt.Println("abi package err:", name, err)
		return ""
	}
	to := common.HexToAddress(contract)
	ret, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From:     common.Address{},
		To:       &to,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     data,
	}, nil)
	if err != nil {
		fmt.Println("call contract err:", err)
		return ""
	}
	return HexToString(hex.EncodeToString(ret))
}

func ReadTokenInfo(address string, client *ethclient.Client, url string) (string, string, uint8, *big.Int, error) {
	to := common.HexToAddress(address)
	instance, err := NewErc20(to, client)
	if err != nil {
		return "", "", 0, nil, errors.New("new erc20 err:" + err.Error())
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		if strings.Contains(err.Error(), "no contract code at given address") {
			symbol = "killed"
		} else {
			symbol, err = instance.SYMBOL(&bind.CallOpts{})
			if err != nil {
				symbol = CallContract(client, address, "symbol")
			}
		}
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		if strings.Contains(err.Error(), "no contract code at given address") {
			name = "killed"
		} else {
			name, err = instance.NAME(&bind.CallOpts{})
			if err != nil {
				name = CallContract(client, address, "name")
			}
		}
	}
	if name == "" {
		name = symbol
	}
	if symbol == "" {
		symbol = name
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		decimals, err = instance.DECIMALS(&bind.CallOpts{})
		if err != nil {
			decimals = 0
		}
	}
	supply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		supply = new(big.Int).SetUint64(0)
	}
	return name, symbol, decimals, supply, nil
}

func HexFormat(s string) string {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return s
}

func HexToString(h string) string {
	n, err := hex.DecodeString(HexFormat(h))
	if err != nil {
		fmt.Println(err)
	}
	return TrimZero(string(n))
}

func TrimZero(s string) string {
	str := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			continue
		}
		str = append(str, v)
	}
	return string(str)
}
