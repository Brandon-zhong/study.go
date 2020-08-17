package main

import "fmt"

var a = "G"

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()

}
type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type hahah interface {

}
