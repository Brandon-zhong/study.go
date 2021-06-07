package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var sl []int
	for i := 0; i < 2048; i++ {
		sl = append(sl, i)
		fmt.Println()
	}
	fmt.Println(sl)
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}

func yingdiReq() {
	params := make(map[string]string)
	params["timestamp"] = "123"
	params["fire"] = "20"
	params["activity_order_id"] = "23134322343"
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	paramStr, sign := generateSign(params, "25f9e794323b453885f5181f1b624d0b")

	req, _ := http.NewRequest("POST", "http://localhost:8023/activity/fire/offer", strings.NewReader(paramStr+"sign="+sign))
	req.Header.Add("App-Udid", "unknow")
	req.Header.Add("App-Version", "826")
	req.Header.Add("Login-Token", "7fb0f281-b94e-4dd5-8fcd-f310491a7b9c")
	req.Header.Add("Platform", "ios")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Use-Traditional", "false")
	req.Header.Add("Activity-Id", "123456789")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
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

func do(l *sync.Mutex) {
}

func getAccountIdFromUserIdAndIndex(userId, index int) int {
	return userId | (index << 27 & 0xffffffff)
}
func getUserIdAndIndexFromAccountId(accountId int) (userId, index int) {
	accountId = accountId & 0xffffffff
	index = accountId >> 27
	userId = accountId & 0x7ffffff
	return
}

func iterateWithDel() {
	type d struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var list []d
	for i := 0; i < 10; i++ {
		list = append(list, d{
			Name: "345",
			Age:  i,
		})
	}
	a := 5
	delete := false
	for i := 0; i < len(list); {
		if list[i].Age == a {
			if delete {
				list = append(list[:i], list[i+1:]...)
				continue
			}
			list[i].Name = "haha"
		}
		i++
	}
	fmt.Println(list)
}

func ldkfj() {
	score := int64(341)<<16 | (int64(274877906943 & 0xffff))
	fmt.Println(score)
	fmt.Println(strconv.FormatInt(int64(274877906943&0xffffffff), 2))
	fmt.Println(strconv.FormatInt(int64(341)<<32, 2))
}
