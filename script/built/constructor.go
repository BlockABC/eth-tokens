package built

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const URLTokenList = "https://raw.githubusercontent.com/MyEtherWallet/ethereum-lists/master/dist/tokens/eth/tokens-eth.json"

type Logo struct {
	Src      string      `json:"src"`
	Width    interface{} `json:"width"` //这两个字段在返回中有的是字符，有的是数字，需要进行二次解析
	Height   interface{} `json:"height"`
	IpfsHash string      `json:"ipfs_hash"`
}
type Support struct {
	Email string `json:"email"`
	Url   string `json:"url"`
}

type Social struct {
	Blog      string `json:"blog"`
	Chat      string `json:"chat"`
	Facebook  string `json:"facebook"`
	Forum     string `json:"forum"`
	Github    string `json:"github"`
	Gitter    string `json:"gitter"`
	Instagram string `json:"instagram"`
	Linkedin  string `json:"linkedin"`
	Reddit    string `json:"reddit"`
	Slack     string `json:"slack"`
	Telegram  string `json:"telegram"`
	Twitter   string `json:"twitter"`
	Youtube   string `json:"youtube"`
}
type TokenInfo struct {
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Address    string  `json:"address"`
	EnsAddress string  `json:"ens_address"`
	Decimals   int     `json:"decimals"`
	Website    string  `json:"website"`
	Logo       Logo    `json:"logo"`
	Support    Support `json:"support"`
	Social     Social  `json:"social"`
}

func (token *TokenInfo) Bytes() ([]byte, error) {
	t := Token{
		Name:     token.Name,
		Symbol:   token.Symbol,
		Contract: token.Address,
		Decimals: token.Decimals,
		Logo:     token.Logo.Src,
		Desc: struct {
			En string `json:"en"`
			Zh string `json:"zh"`
		}{
			En: token.Type,
			Zh: token.Type,
		},
		WebSite:    token.Social.Blog,
		WhitePaper: "",
		Invalid:    true,
		Links: struct {
			Twitter  string `json:"twitter"`
			Telegram string `json:"telegram"`
		}{
			Twitter:  token.Social.Twitter,
			Telegram: token.Social.Telegram,
		},
	}
	return json.MarshalIndent(t, "", "    ")
}

func TokenListFromGit(url string) (tokenLists []*TokenInfo, err error) {
	requester := gorequest.New().Get(url).Timeout(time.Second*5).Retry(5, time.Second, http.StatusRequestTimeout, http.StatusBadRequest)
	resp, body, errs := requester.EndBytes()
	if errs != nil || resp.StatusCode != http.StatusOK {
		req, err := requester.MakeRequest()
		if err == nil && req != nil && resp != nil {
			fmt.Printf("request status:%d, body:%+v\n", resp.StatusCode, req)
		}
		var errStr string
		for _, e := range errs {
			errStr += e.Error()
		}
		return nil, errors.New(errStr)
	}
	if err := json.Unmarshal(body, &tokenLists); err != nil {
		return nil, err
	}
	for _, token := range tokenLists {
		switch reflect.TypeOf(token.Logo.Width).Kind() {
		case reflect.Float64:
			token.Logo.Width = fmt.Sprintf("%v", token.Logo.Width.(float64))
			token.Logo.Height = fmt.Sprintf("%v", token.Logo.Height.(float64))
		default:
			continue
		}
	}

	return tokenLists, nil
}

func InitializeTokens(dir string, tokenLists []*TokenInfo) error {
	for _, token := range tokenLists {
		time.Sleep(time.Millisecond * 500)
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", dir, strings.ToLower(token.Address)), os.ModePerm); err != nil {
			return err
		}
		if err := WriteTokenInfo(dir, token); err != nil {
			return err
		}
	}
	return nil
}

func WriteTokenInfo(dir string, token *TokenInfo) error {
	f := fmt.Sprintf("%s/%s/token.json", dir, strings.ToLower(token.Address))
	if _, err := os.Stat(f); err == nil || os.IsExist(err) {
		fmt.Println("the contract is exist, rewrite info:", token.Address)
		//return nil
	}
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer checkError(file.Close)
	data, err := token.Bytes()
	if err != nil {
		return err
	}
	if _, err := file.Write(data); err != nil {
		return err
	}
	p := fmt.Sprintf("%s/%s/token.png", dir, strings.ToLower(token.Address))
	if err := RequestIcon(token.Logo.Src, p); err != nil {
		return err
	}
	fmt.Println("write success:", token.Address, FormatSymbol(token.Symbol))
	return nil
}

func RequestIcon(url, p string) error {
	if url == "" {
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(errors.New("get" + url + "error:" + err.Error()))
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("get" + url + "error:" + resp.Status)
		return nil
	}
	defer checkError(resp.Body.Close)
	pix, err := ioutil.ReadAll(resp.Body)
	out, err := os.Create(p)
	if err != nil {
		return errors.New("os create png err:" + err.Error())
	}
	defer checkError(out.Close)
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		return errors.New("io copy err:" + err.Error())
	}
	return nil
}

func FormatSymbol(s string) string {
	if !strings.ContainsAny(s, `;'\"&<>$ф`) {
		return strings.Replace(s, ` `, ``, -1)
	}
	s = strings.Replace(s, `;`, ``, -1)
	s = strings.Replace(s, `'`, ``, -1)
	s = strings.Replace(s, `\`, ``, -1)
	s = strings.Replace(s, `"`, ``, -1)
	s = strings.Replace(s, `&`, ``, -1)
	s = strings.Replace(s, `<`, ``, -1)
	s = strings.Replace(s, `>`, ``, -1)
	s = strings.Replace(s, `$`, ``, -1)
	s = strings.Replace(s, `ф`, ``, -1)
	return strings.Replace(s, ` `, ``, -1)
}

func checkError(f func() error) {
	if err := f(); err != nil {
		fmt.Println(err)
	}
}
