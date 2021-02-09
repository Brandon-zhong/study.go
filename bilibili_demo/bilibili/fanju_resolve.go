package bilibili

import (
	"bufio"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"study.go/util"
)

var SESSDATA string

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

func StartDownloadFanJu(url, folder string, quality int) {

	//quality := 112
	folder = util.GetDownloadDirIfFolderIsNil(folder)

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
	//fmt.Println(data)
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
	fmt.Printf("开始下载，要下载的 %s 一共有 %d 集视频\n", fanjuTitle, len(videoList))
	downloadVideoList(videoList, getPlayListForFanJu)
}

func InputFanJuParam(input *bufio.Scanner) {
	fmt.Print("番剧下载需要用户的SESSDATA，请从你的网页cookie中找出SESSDATA值填入（如果确定要下载的视频不需要大会员可以不填）：\n> ")
	input.Scan()
	SESSDATA = input.Text()
	fmt.Print("请输入要下载番剧的链接，例如（https://www.bilibili.com/bangumi/play/ep90830）：\n> ")
	input.Scan()
	url := input.Text()
	fmt.Println("请输入你要下载视频的清晰度（1080p+:112;1080p:80;720p60:74;720p:64;480p:32;360p:16）:")
	fmt.Print("请输入112或80或74或64或32或16：\n> ")
	input.Scan()
	quality, err := strconv.Atoi(input.Text())
	if err != nil {
		fmt.Println("清晰度输入错误，请输入有效的清晰度数字！")
		return
	}
	StartDownloadFanJu(url, "", quality)
}
