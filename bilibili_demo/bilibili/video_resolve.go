package bilibili

import (
	"crypto/md5"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

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

//获取是的播放地址
func getVideoPlayUrl(videoInfo *VideoInfo) (flag bool) {
	//获取appKey和secret
	key, sec := GetAppKey("rbMCKn@KuamXWlPMoJGsKcbiJKUfkPF_8dABscJntvqhRSETg")
	//生成参数和校验和
	quality := strconv.Itoa(videoInfo.quality)
	params := fmt.Sprintf("appkey=%s&cid=%s&otype=json&qn=%s&quality=%s&type=", key, strconv.FormatInt(videoInfo.cid, 10), quality, quality)
	checkSum := fmt.Sprintf("%x", md5.Sum([]byte(params+sec)))
	urlApi := fmt.Sprintf("https://interface.bilibili.com/v2/playurl?%s&sign=%s", params, checkSum)

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

func StartDownloadVideo(videoId, folder string) {

	//folder := "E:\\nfs\\download"
	folder = getDownloadDirIfFolderIsNil(folder)
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
	folder = filepath.Join(folder, data.Get("title").String())
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
			filePath: folder,
		}
		videoInfoList = append(videoInfoList, videoInfo)
	}

	downloadVideoList(videoInfoList, getVideoPlayUrl)
}
