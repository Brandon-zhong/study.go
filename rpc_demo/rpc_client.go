package main

import (
	"fmt"
	"net/rpc"
	"reflect"
)

func main() {
	service, err := DiaHelloService("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("dia error:", err)
	}
	var reply string
	if err = service.Hello("kfjlksdj", &reply); err != nil {
		fmt.Println("call error:", err)
	}

}

type HelloServiceClient struct {
	*rpc.Client
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	typeOf := reflect.TypeOf(h)
	kind := typeOf.Kind()
	fmt.Println(kind)
	return h.Client.Call("fsdf.Hello", request, reply)
}

func DiaHelloService(network, address string) (*HelloServiceClient, error) {
	dial, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: dial}, nil
}
