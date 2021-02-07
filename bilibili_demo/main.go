package main

import (
	"bufio"
	"fmt"
	"os"
	"study.go/bilibili_demo/bilibili"
)

func main() {

	//terminalDownload()

	bilibili.SESSDATA = "2d68f213%2C1626184663%2C485c1*11"
	bilibili.StartDownloadFanJu("https://www.bilibili.com/bangumi/play/ep90830", "E:\\nfs\\download", 112)

}

func terminalDownload() {
again:
	fmt.Print("请选择你要下载视频的模式（1-普通视频，2-番剧视频）：\n> ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	mode := scan.Text()
	if mode == "1" {
		bilibili.InputVideoParam(scan)
	} else if mode == "2" {
		bilibili.InputFanJuParam(scan)
	} else {
		fmt.Println("未识别的模式，请重新输入！")
		goto again
	}
}
