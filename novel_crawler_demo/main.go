package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"study.go/util"
	"sync"
	"time"
)

//只针对37zw.la网址或类型的网页结构的网页有效

func getList(url string) {
	start := time.Now().Unix()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	reader := simplifiedchinese.GBK.NewDecoder().Reader(resp.Body)
	document, _ := goquery.NewDocumentFromReader(reader)

	//下载到一个临时目录中，每章一个文件，完成后再进行合并
	tmpDir := util.GetDownloadDirIfFolderIsNil("E:\\nfs\\download\\tmp")
	_ = os.MkdirAll(tmpDir, 0777)

	title := document.Find("div#info h1").Text()

	downloadConsole := make(chan bool, 8)
	//用于等待所有携程执行完成
	var wait sync.WaitGroup
	document.Find("div#list dl dd a").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Attr("href")
		text := selection.Text()
		//fmt.Println(i, val, text, num, compile.ReplaceAllString(num, ""))
		wait.Add(1)
		go func() {
			//申请执行
			downloadConsole <- true
			defer func() {
				//执行完成，释放资源
				<-downloadConsole
				wait.Done()
			}()
			fmt.Println("start downloading", text)
			filePath := filepath.Join(tmpDir, val[:len(val)-5])
			downloadNovel(url+val, filePath)
		}()
	})
	wait.Wait()
	mergeFile(tmpDir, filepath.Join("E:\\nfs\\download", title+".txt"))
	log.Printf("\nall download has finished, spend time --> %s.", util.ResolveTime(time.Now().Unix()-start))
}

func getList2(url string) {
	start := time.Now().Unix()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	reader := simplifiedchinese.GBK.NewDecoder().Reader(resp.Body)
	document, _ := goquery.NewDocumentFromReader(reader)

	//下载到一个临时目录中，每章一个文件，完成后再进行合并
	tmpDir := util.GetDownloadDirIfFolderIsNil("E:\\nfs\\download\\tmp")
	_ = os.MkdirAll(tmpDir, 0777)

	title := document.Find("div#info h1").Text()

	type info struct {
		url string
		path string
	}
	proc := util.NewConcurrentProc(8, func(i interface{}) {
		f := i.(info)
		downloadNovel(f.url, f.path)
	}).Start()
	document.Find("div#list dl dd a").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Attr("href")
		text := selection.Text()
		fmt.Println("start downloading", text)
		//fmt.Println(i, val, text, num, compile.ReplaceAllString(num, ""))
		proc.AddInfo(info{url: url + val, path: filepath.Join(tmpDir, val[:len(val)-5])})
	})
	proc.Done()
	proc.Wait()
	mergeFile(tmpDir, filepath.Join("E:\\nfs\\download", title+".txt"))
	log.Printf("\nall download has finished, spend time --> %s.", util.ResolveTime(time.Now().Unix()-start))
}



func mergeFile(tmpDir, targetFilePath string) {
	dir, err := os.Open(tmpDir)
	if err != nil {
		log.Println(err)
		return
	}
	fileInfos, err := dir.Readdir(-1)
	sort.Slice(fileInfos, func(i, j int) bool {
		return util.MustInt(strconv.Atoi(fileInfos[i].Name())) < util.MustInt(strconv.Atoi(fileInfos[j].Name()))
	})
	if err != nil {
		log.Println(err)
		return
	}
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer targetFile.Close()
	total := len(fileInfos)
	fmt.Println("start merge file.")
	for index, fileInfo := range fileInfos {
		data, err := ioutil.ReadFile(filepath.Join(tmpDir, fileInfo.Name()))
		if err != nil {
			log.Println(err)
			return
		}
		Bar(index, total)
		_, _ = targetFile.Write(data)
		_, _ = targetFile.WriteString("\n\n\n\n")
	}
	//合并完成后删除文件
	_ = os.RemoveAll(tmpDir)
}

func Bar(index, total int) {
	count := index*100/total + 1
	bar := strings.Repeat("#", count) + strings.Repeat(" ", 100-count)
	fmt.Printf("\r[%s]%3d%%  %8d/%d", bar, count, index, total)
}

func downloadNovel(url, filePath string) bool {

	content := getContent(url)

	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return false
	}
	defer file.Close()
	if _, err = file.WriteString(content); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	reader := simplifiedchinese.GBK.NewDecoder().Reader(resp.Body)
	document, _ := goquery.NewDocumentFromReader(reader)

	title := document.Find("div.bookname h1").Text()
	find := document.Find("div#content")

	var buf bytes.Buffer
	buf.WriteString(title)
	buf.WriteString("\n\n\n\n")

	// Slightly optimized vs calling Each: no single selection object created
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			// Keep newlines and spaces, like jQuery
			buf.WriteString(n.Data)
		} else if n.Type == html.ElementNode && n.Data == "br" {
			buf.WriteString("\n")
		}
		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	for _, n := range find.Nodes {
		f(n)
	}
	return buf.String()
}

func main() {

	//getContent()
	getList2("https://www.37zw.la/22/22055/")
	//mergeFile("E:\\nfs\\download\\tmp", "E:\\nfs\\download\\修真聊天群.txt")

}
