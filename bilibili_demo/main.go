package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var pool, _ = ants.NewPool(100)

func main() {

}

type VideoInfo struct {
	aid     int    //视频id
	cid     int    //视频单P id
	title   string //单P 的标题
	page    int    //单P 号
	referer string //一定要加
}



func genCheckRedirectFun(referer string) func(req *http.Request, via []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", referer)
		return nil
	}
}

func download() {
	referer := "https://api.bilibili.com/x/web-interface/view?aid=245576115/?p=1"
	url := "http://upos-sz-mirrorcos.bilivideo.com/upgcxcode/26/05/262640526/262640526-1-80.flv?e=ig8euxZM2rNcNbR1hbUVhoM1hWNBhwdEto8g5X10ugNcXBlqNxHxNEVE5XREto8KqJZHUa6m5J0SqE85tZvEuENvNC8xNEVE9EKE9IMvXBvE2ENvNCImNEVEK9GVqJIwqa80WXIekXRE9IMvXBvEuENvNCImNEVEua6m2jIxux0CkF6s2JZv5x0DQJZY2F8SkXKE9IB5QK==&deadline=1611311779&gen=playurl&nbs=1&oi=3230711203&os=cosbv&platform=pc&trid=8c48d764e4b247a9ab3e504fc911b1f0&uipk=5&upsig=d82aff5dd460cf2db4a5bcdfd09d6160&uparams=e,deadline,gen,nbs,oi,os,platform,trid,uipk&mid=0"
	cid := 262640526
	var aid int64 = 245576115
	title := "王道训练营-C语言教程"
	page := 25
	order := 1

	client := http.Client{CheckRedirect: genCheckRedirectFun(referer)}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(url, err)
		return
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:56.0) Gecko/20100101 Firefox/56.0")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Range", "bytes=0-")
	request.Header.Set("Referer", referer)
	request.Header.Set("Origin", "https://www.bilibili.com")
	request.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("Fail to download the video %d,err is %s", cid, err)
		return
	}

	if resp.StatusCode != http.StatusPartialContent {
		log.Fatalf("Fail to download the video %d,status code is %d", cid, resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	aidPath := GetAidFileDownloadDir(aid, title)
	filename := fmt.Sprintf("%d_%d.flv", page, order)
	file, err := os.Create(filepath.Join(aidPath, filename))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer file.Close()

	log.Println(title + ":" + filename + " is downloading.")
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Printf("Failed to download video %d", cid)
		return
	}
	log.Println(title + ":" + filename + " has finished.")
}

func GetAidFileDownloadDir(aid int64, title string) string {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	title = strings.Replace(title, ":", "", -1)
	title = strings.Replace(title, "\\", "", -1)
	title = strings.Replace(title, "/", "", -1)
	title = strings.Replace(title, "*", "", -1)
	title = strings.Replace(title, "?", "", -1)
	title = strings.Replace(title, "\"", "", -1)
	title = strings.Replace(title, "<", "", -1)
	title = strings.Replace(title, ">", "", -1)
	title = strings.Replace(title, "|", "", -1)
	// remove special symbal
	fullDirPath := filepath.Join(curDir, "download", fmt.Sprintf("%d_%s", aid, title))
	err = os.MkdirAll(fullDirPath, 0777)
	if err != nil {
		panic(err)
	}
	return fullDirPath
}
