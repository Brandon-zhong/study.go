package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	for i := 0; i < 1000; i++ {
		//fmt.Println(uuid.NewV4().String())
		fmt.Println("demo --> ", NewDemo().String())
	}

}

func NewDemo() *demo {
	d := demo{}
	d.safeRandom(d[:])
	d.SetVersion(uuid.V4)
	d.SetVariant(uuid.VariantRFC4122)
	return &d
}

type demo [16]byte

func (d *demo) safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}

func (d *demo) SetVersion(v byte) {
	d[6] = (d[6] & 0x0f) | (v << 4)
}

func (d *demo) SetVariant(v byte) {
	switch v {
	case uuid.VariantNCS:
		d[8] = d[8]&(0xff>>1) | (0x00 << 7)
	case uuid.VariantRFC4122:
		d[8] = d[8]&(0xff>>2) | (0x02 << 6)
	case uuid.VariantMicrosoft:
		d[8] = d[8]&(0xff>>3) | (0x06 << 5)
	case uuid.VariantFuture:
		fallthrough
	default:
		d[8] = d[8]&(0xff>>3) | (0x07 << 5)
	}
}

func (d demo) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], d[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], d[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], d[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], d[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], d[10:])

	return string(buf)
}
