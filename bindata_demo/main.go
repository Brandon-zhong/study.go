package main

import "fmt"

//go:embed data/data.info
var data []byte

func main() {
	embedTest()
}

//2020-10-30提交的实现，需要在go1.16才能使用
//TODO 以后试一下embed的使用
func embedTest() {
	fmt.Println(string(data))
}

//bindata框架做的就是将文件转成go文件，
//文件的内容以字节数组的形式保存，并且封装了一些方法，
//通过入口方法直接通过相对路径返回文件的字节数组
func binDataTest() {
	bytes, _ := Asset("data/data.info")
	fmt.Println(string(bytes))
}
