package ether_scan

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/eager7/eth_tokens/script/built"
	"github.com/parnurzeal/gorequest"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const urlEtherScan = `https://etherscan.io/tokens?ps=100&p=`
const pageMax = 10

func RequestErc20ListByPage(url string) ([]*built.TokenInfo, error) {
	res := gorequest.New().Proxy(UserProxyLists[rand.Intn(len(UserProxyLists))]).Set("user-agent", UserAgentLists[rand.Intn(len(UserAgentLists))])
	ret, body, errs := res.Timeout(time.Second*5).Retry(5, time.Second, http.StatusRequestTimeout, http.StatusBadRequest).Get(url).End()
	if errs != nil || ret.StatusCode != http.StatusOK {
		req, err := res.MakeRequest()
		if err == nil && req != nil && ret != nil {
			fmt.Printf("request status:%d, body:%+v\n", ret.StatusCode, req)
		}
		var errStr string
		for _, e := range errs {
			errStr += e.Error()
		}
		return nil, errors.New("request err:" + errStr)
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, errors.New("NewDocumentFromReader err:" + err.Error())
	}

	var tokens []*built.TokenInfo
	dom.Find("div.media").Each(func(i int, selection *goquery.Selection) {
		ret, err := selection.Html()
		if err != nil {
			fmt.Println("selection to html err:" + err.Error())
			return
		}
		reIcon := regexp.MustCompile(`.*?src="(?P<icon>[^"]*)(?s:".*)`)
		reContract := regexp.MustCompile(`.*?href="/token/(?P<contract>[^"]*)(?s:".*)`)
		icon := reIcon.ReplaceAllString(ret, "$icon")
		contract := reContract.ReplaceAllString(ret, "$contract")
		if !IsValidAddress(contract) {
			fmt.Println("IsValidAddress false:", contract, ret)
		}
		token := built.TokenInfo{
			Address: contract,
			Logo: built.Logo{
				Src: "https://etherscan.io" + icon,
			},
		}
		tokens = append(tokens, &token)
	})

	return tokens, nil
}

func IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
