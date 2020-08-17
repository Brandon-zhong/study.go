package main

import (
	"fmt"
	"net"
	"net/rpc"
	"reflect"
)

const HelloServiceName = "path/to/pkg,HelloService"

func main() {

	err := RegisterHelloService(new(HelloService))
	if err != nil {
		fmt.Println("register error:", err)
	}
	listen, _ := net.Listen("tcp", "127.0.0.1:8090")
	for {
		conn, _ := listen.Accept()
		rpc.ServeConn(conn)
	}

}

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloServiceDemo interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	typeOf := reflect.TypeOf(svc)
	return rpc.RegisterName(typeOf.Name(), svc)
}

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {

	*reply = "hello:" + request
	return nil
}
