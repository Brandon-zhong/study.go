package bilibili

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"study.go/util"
	"sync"
	"sync/atomic"
	"time"
)

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

func downloadVideoList(videoInfoList []VideoInfo, getPlayUrl func(videoInfo *VideoInfo) (flag bool)) {

	startTime := time.Now().Unix()
	size := len(videoInfoList)
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
			downloadVideoWithRetry(&videoInfo, 3, getPlayUrl)
		}()
	}
	wait.Wait()
	log.Printf("all video has finished, spend time --> %s.", util.ResolveTime(time.Now().Unix()-startTime))
	time.Sleep(2 * time.Second)
}

func downloadVideoWithRetry(videoInfo *VideoInfo, retry int, getPlayUrl func(videoInfo *VideoInfo) (flag bool)) {
	var flag = false
	for i := 0; i < retry && !flag; i++ {
		flag = downloadVideo(videoInfo, getPlayUrl)
	}
}

func downloadVideo(videoInfo *VideoInfo, getPlayUrl func(videoInfo *VideoInfo) (flag bool)) (flag bool) {

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
		filePath := filepath.Join(videoInfo.filePath, fileName+".flv")
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
