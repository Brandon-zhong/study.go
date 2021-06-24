package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"study.go/util"
	"time"
)

//只针对37zw.la网址或类型的网页结构的网页有效

type value struct {
	url     string
	title   string
	content *string
}

func downloadNovel(url, fileDir string) {
	start := time.Now().Unix()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	reader := simplifiedchinese.GBK.NewDecoder().Reader(resp.Body)
	document, _ := goquery.NewDocumentFromReader(reader)
	title := document.Find("div#info h1").Text()

	pool := util.NewSeqConcurrentProc(8, func(val interface{}) interface{} {
		v := val.(value)
		content := getContent(v.url)
		v.content = &content
		return v
	})

	in, out := pool.Process()

	go writeFileFromChan(out, filepath.Join(fileDir, title+".txt"))

	document.Find("div#list dl dd a").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Attr("href")
		text := selection.Text()
		fmt.Println("start downloading", text, val)
		in <- value{
			url:   url + val,
			title: text,
		}
	})
	pool.CancelFn()
	log.Printf("\nall download has finished, spend time --> %s.", util.ResolveTime(time.Now().Unix()-start))
}

func writeFileFromChan(ch <-chan interface{}, fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		panic(err.Error())
	}

	for val := range ch {
		content := val.(value).content
		if content == nil {
			continue
		}
		_, _ = file.WriteString(*content + "\n\n\n")
	}
	_ = file.Close()
}

func Bar(index, total int) {
	count := index*100/total + 1
	bar := strings.Repeat("#", count) + strings.Repeat(" ", 100-count)
	fmt.Printf("\r[%s]%3d%%  %8d/%d", bar, count, index, total)
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
	return util.BytesToString(buf.Bytes())
}

func main() {

	//getContent()
	downloadNovel("https://www.777zw.la/22/22055/", "E:\\nfs\\download")
	//mergeFile("E:\\nfs\\download\\tmp", "E:\\nfs\\download\\修真聊天群.txt")

}
