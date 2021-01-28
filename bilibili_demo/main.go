package main

import (
	"crypto/md5"
	"fmt"
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
	"sync/atomic"
	"time"
)

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

	startDownload("BV1BE411D7ii")
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
	var videoInfoList []VideoInfo
	for i := range array {
		//封装下载信息
		videoInfo := VideoInfo{
			aid:      aid,
			cid:      array[i].Get("cid").Int(),
			title:    array[i].Get("part").String(),
			page:     int(array[i].Get("page").Int()),
			baseUrl:  baseUrl,
			quality:  quality,
			filePath: filder,
		}
		videoInfoList = append(videoInfoList, videoInfo)
	}

	size := len(array)
	//开始下载，使用指定容量的管道来控制同时下载的数量
	videoInfoChan := make(chan bool, 16)
	//用于等待所有携程执行完成
	var wait sync.WaitGroup
	wait.Add(size)
	var count int64 = -1
	for i := 0; i < size; i++ {
		go func() {
			//获取一个值
			videoInfo := videoInfoList[atomic.AddInt64(&count, 1)]
			//通过管道申请开始下载，
			videoInfoChan <- true
			defer func() {
				//处理完成后空出管道的位置，通知其他等待的携程开始处理
				<-videoInfoChan
				wait.Done()
			}()
			//fmt.Println(strconv.Itoa(int(newSize)) + " -- " + videoInfo.title)
			//time.Sleep(time.Second * 3)
			downloadVideoWithRetry(&videoInfo, 3)
		}()
	}
	wait.Wait()
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
func getPlayUrl(videoInfo *VideoInfo) (flag bool) {
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
		return false
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read respond fail, error --> %s", err)
		return false
	}
	var videoList []string
	for _, i := range gjson.GetBytes(all, "durl").Array() {
		videoList = append(videoList, i.Get("url").String())
	}
	videoInfo.urlList = videoList
	return true
}

func downloadVideoWithRetry(videoInfo *VideoInfo, retry int) {
	var flag = false
	for i := 0; i < retry && !flag; i++ {
		flag = downloadVideo(videoInfo)
	}
}

func downloadVideo(videoInfo *VideoInfo) (flag bool) {

	//获取下载链接
	flag = getPlayUrl(videoInfo)

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
			flag = false
			continue
		}
		//生成完整的文件目录
		fileName := videoInfo.title
		if urlSize != 1 {
			fileName = videoInfo.title + "-" + strconv.Itoa(index)
		}
		filePath := filepath.Join(videoInfo.filePath, fileName+".mp4")
		_ = os.MkdirAll(videoInfo.filePath, 0777)
		_ = os.Remove(filePath)
		file, err := os.Create(filePath)
		if err != nil {
			log.Printf("open file error, err --> %s", err)
			flag = false
			continue
		}
		log.Println(fileName + " is downloading.   file size --> " + strconv.FormatInt(resp.ContentLength, 10))
		if _, err = io.Copy(file, resp.Body); err != nil {
			log.Printf("download video error, err --> %s", err)
			flag = false
			continue
		}
		log.Println(fileName + " has finished.")
		resp.Body.Close()
		_ = file.Close()
	}
	return
}
