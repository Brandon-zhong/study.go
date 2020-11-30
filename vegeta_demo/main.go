package main

import (
	"crypto/md5"
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"net/http"
	"os"
	"sort"
	"time"
)

func constructHeader() *http.Header {
	header := http.Header{}
	header.Set("App-Udid", "aaaaa")
	header.Set("App-Version", "800")
	header.Set("Login-Token", "nologin")
	header.Set("Platform", "ios")
	header.Set("Preid", "29e2485f4dcaca1e8063cb0564855a4e")
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Use-Traditional", "asdf")
	return &header
}

func generateSign(params map[string]string, secretKey string) (string, string) {
	var keyList []string
	for key, _ := range params {
		keyList = append(keyList, key)
	}
	sort.Strings(keyList)

	var str string
	for _, key := range keyList {
		str += fmt.Sprintf("%s=%s&", key, params[key])
	}
	return str, fmt.Sprintf("%x", md5.Sum([]byte(str+"key="+secretKey)))
}

func main() {

	rate := vegeta.Rate{Freq: 3000, Per: time.Second}
	duration := 30 * time.Second

	var params = make(map[string]string)
	params["timestamp"] = "123"
	/*params["post_id"] = "2313011"
	params["last_comment_id"] = "0"
	params["size"] = "6"
	params["order"] = "created"
	params["vote_faction"] = "0"*/

	//计算sign
	var secretKey = "a94c290bc99ffc54307c4f1652a9c41a"
	paramStr, sign := generateSign(params, secretKey)
	paramStr += "sign=" + sign

	targeter := vegeta.NewStaticTargeter(
		vegeta.Target{
			Method: "POST",
			URL:    "http://192.168.131.239:8023/app/bbs/tags",
			Header: *constructHeader(),
			Body:   []byte(paramStr),
		})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	reporter := vegeta.NewTextReporter(&metrics)
	_ = reporter.Report(os.Stdout)

}
