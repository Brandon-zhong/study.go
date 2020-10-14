package main

import "fmt"

func main() {

	var d = Demo{name: "asdf"}
	var ha = Haha{Demo: d}
	var ha1 = Haha{Demo: d}
	fmt.Println(ha.name, " --- ", ha1.name)
	ha.name = "zhong"
	d.name = "123456"
	ha1.h()
	fmt.Println(ha.name, " --- ", ha1.name)

	var d2 = Demo{name: "qwopir"}
	var he = Heihei{Demo: &d2}
	var he2 = Heihei{Demo: &d2}
	fmt.Println(he.name," -- ", he2.name)
	he.name = "sheng"
	he2.h()
	fmt.Println(he.name," -- ", he2.name)

}

type Demo struct {
	name string
}

func (d *Demo) h() {
	d.name = d.name + "1"
}

type Haha struct {
	Demo
}

type Heihei struct {
	*Demo
}
