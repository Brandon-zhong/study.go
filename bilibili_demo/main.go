package main

import (
	"crypto/md5"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var pool, _ = ants.NewPool(100)

func main() {
	/*videoInfo := VideoInfo{
		aid:      70211798,
		cid:      223418336,
		title:    "【2021版】2.1.2_BCD码",
		page:     9,
		baseUrl:  "https://www.bilibili.com/video/av70211798",
		quality:  80,
		filePath: "E:\\nfs\\download",
	}
	videoInfo := VideoInfo{
		aid:      498807073,
		cid:      209502920,
		title:    "克隆史（上）",
		page:     9,
		baseUrl:  "https://www.bilibili.com/video/av498807073",
		quality:  80,
		filePath: "E:\\nfs\\download",
	}
	getPlayUrl(&videoInfo)
	downloadVideo(&videoInfo)*/

	startDownload("BV1X7411b7Yc")
}

type VideoInfo struct {
	aid      int64    //视频id
	cid      int64    //视频单P id
	title    string   //单P 的标题
	page     int      //单P 号
	baseUrl  string   //一定要加, 基础url
	quality  int      //视频质量 1080p:80;720p:64;480p:32;360p:16
	filePath string   //视频下载的地址，只是目录地址
	urlList  []string //视频的url地址
}

func genCheckRedirectFun(referer string) func(req *http.Request, via []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", referer)
		return nil
	}
}

const entropy = "rbMCKn@KuamXWlPMoJGsKcbiJKUfkPF_8dABscJntvqhRSETg"
const playUrlApi = "https://interface.bilibili.com/v2/playurl?%s&sign=%s"
const paramsTmp = "appkey=%s&cid=%s&otype=json&qn=%s&quality=%s&type="

func GetAppKey(entropy string) (appkey, sec string) {
	revEntropy := ReverseRunes([]rune(entropy))
	for i := range revEntropy {
		revEntropy[i] = revEntropy[i] + 2
	}
	ret := strings.Split(string(revEntropy), ":")

	return ret[0], ret[1]
}
func ReverseRunes(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func startDownload(videoId string) {

	startTime := time.Now().Unix()
	filder := "E:\\nfs\\download"
	quality := 80

	params := "aid=" + videoId
	if strings.HasPrefix(videoId, "BV") {
		params = "bvid=" + videoId
	}
	//获取视频列表
	baseUrl := fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?%s", params)
	request, _ := http.NewRequest("GET", baseUrl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("get base url error, err --> %s", err)
		return
	}
	defer resp.Body.Close()
	all, _ := ioutil.ReadAll(resp.Body)
	data := gjson.GetBytes(all, "data")
	aid := data.Get("aid").Int()
	filder = filepath.Join(filder, data.Get("title").String())
	array := data.Get("pages").Array()
	//开始下载，生成16个协程的协程池
	pool, _ := ants.NewPool(32)
	videoSize := len(array)
	videoInfoChan := make(chan VideoInfo)
	go func() {
		for i := range array {
			j := i
			_ = pool.Submit(func() {
				//封装下载信息
				videoInfo := VideoInfo{
					aid:      aid,
					cid:      array[j].Get("cid").Int(),
					title:    array[j].Get("part").String(),
					page:     int(array[j].Get("page").Int()),
					baseUrl:  baseUrl,
					quality:  quality,
					filePath: filder,
				}
				getPlayUrl(&videoInfo)
				fmt.Println("get play url --> " + videoInfo.title)
				videoInfoChan <- videoInfo
			})
		}
	}()

	var waitGroup sync.WaitGroup
	waitGroup.Add(videoSize)
	for i := 0; i < videoSize; i++ {
		_ = pool.Submit(func() {
			defer waitGroup.Done()
			videoInfo := <-videoInfoChan
			downloadVideo(&videoInfo)
		})
	}
	waitGroup.Wait()
	log.Printf("all video has finished, spend time --> %s.", resolveTime(int(time.Now().Unix()-startTime)))

}

func resolveTime(spendTime int) string {
	second := spendTime % 60
	minute := spendTime / 60
	if minute == 0 {
		return fmt.Sprintf("%ds", second)
	}
	if minute < 60 {
		return fmt.Sprintf("%dm%ds", minute, second)
	}
	hour := minute / 60
	return fmt.Sprintf("%dh%dm%ds", hour, minute, second)
}

//获取是的播放地址
func getPlayUrl(videoInfo *VideoInfo) {
	//获取appKey和secret
	key, sec := GetAppKey(entropy)
	//生成参数和校验和
	quality := strconv.Itoa(videoInfo.quality)
	params := fmt.Sprintf(paramsTmp, key, strconv.FormatInt(videoInfo.cid, 10), quality, quality)
	checkSum := fmt.Sprintf("%x", md5.Sum([]byte(params+sec)))
	urlApi := fmt.Sprintf(playUrlApi, params, checkSum)

	request, _ := http.NewRequest("GET", urlApi, nil)
	request.Header.Set("Referer", videoInfo.baseUrl)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("get play url error, err --> %s,  status code --> %d", err, resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read respond fail, error --> %s", err)
		return
	}
	var videoList []string
	for _, i := range gjson.GetBytes(all, "durl").Array() {
		videoList = append(videoList, i.Get("url").String())
	}
	videoInfo.urlList = videoList
}

func downloadVideo(videoInfo *VideoInfo) {
	urlSize := len(videoInfo.urlList)
	//当视频和多视频下载
	if urlSize == 0 {
		return
	}
	for index, url := range videoInfo.urlList {
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:56.0) Gecko/20100101 Firefox/56.0")
		request.Header.Set("Accept", "*/*")
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Accept-Encoding", "gzip, deflate, br")
		request.Header.Set("Range", "bytes=0-")
		request.Header.Set("Referer", videoInfo.baseUrl)
		request.Header.Set("Origin", "https://www.bilibili.com")
		request.Header.Set("Connection", "keep-alive")
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Printf("get play url error, err --> %s", err)
			continue
		}
		fmt.Println("文件大小：" + strconv.FormatInt(resp.ContentLength, 10))
		//生成完整的文件目录
		fileName := videoInfo.title
		if urlSize != 1 {
			fileName = videoInfo.title + "-" + strconv.Itoa(index)
		}
		filePath := filepath.Join(videoInfo.filePath, fileName+".mp4")
		_ = os.MkdirAll(videoInfo.filePath, 0777)
		file, err := os.Create(filePath)
		if err != nil {
			log.Printf("open file error, err --> %s", err)
			continue
		}
		log.Println(fileName + " is downloading.")
		if _, err = io.Copy(file, resp.Body); err != nil {
			log.Printf("download video error, err --> %s", err)
			continue
		}
		log.Println(fileName + " has finished.")
		resp.Body.Close()
		_ = file.Close()
	}

}
