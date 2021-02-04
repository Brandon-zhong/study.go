package bilibili

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const SESSDATA = "ba234a5b%2C1620707829%2Ccd50d*b1"

func getPlayListForFanJu(video *VideoInfo) (flag bool) {
	urlApi := fmt.Sprintf("https://api.bilibili.com/x/player/playurl?cid=%d&avid=%d&qn=%d", video.cid, video.aid, video.quality)
	request, _ := http.NewRequest("GET", urlApi, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	request.Header.Set("Cookie", "SESSDATA="+SESSDATA)
	request.Header.Set("Host", "api.bilibili.com")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bytes))
	data := gjson.ParseBytes(bytes)
	if data.Get("code").Int() != 0 {
		fmt.Println("注意!当前集数为B站大会员专享,若想下载,Cookie中请传入大会员的SESSDATA")
		os.Exit(0)
		return false
	}
	var urlList []string
	for _, result := range data.Get("data.durl").Array() {
		urlList = append(urlList, result.Get("url").String())
	}
	video.urlList = urlList
	return true
}

func StartDownloadFanJu(url, folder string) {

	quality := 112
	folder = getDownloadDirIfFolderIsNil(folder)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(string(bytes))
	compile, _ := regexp.Compile("INITIAL_STATE__=(.*?\"]});")
	find := compile.Find(bytes)
	j := string(find)
	data := gjson.Parse(j[len("INITIAL_STATE__=") : len(j)-1])
	fanjuTitle := data.Get("mediaInfo.title").String()
	folder = filepath.Join(folder, fanjuTitle)

	array := data.Get("epList").Array()
	var videoList []VideoInfo
	for index := range array {
		result := array[index]
		videoList = append(videoList, VideoInfo{
			aid:      result.Get("aid").Int(),
			cid:      result.Get("cid").Int(),
			title:    strings.TrimSpace(result.Get("titleFormat").String() + " " + result.Get("longTitle").String()),
			page:     int(result.Get("i").Int()),
			baseUrl:  url,
			quality:  quality,
			filePath: folder,
		})
	}
	downloadVideoList(videoList, getPlayListForFanJu)
}
