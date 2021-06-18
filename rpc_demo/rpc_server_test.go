package rpc_demo

import (
	"log"
	"net"
	"net/rpc"
	"testing"
)

func TestRpcServer(t *testing.T) {

	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("ListenTcp error : ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("ListenTcp error : ", err)
		}
		rpc.ServeConn(conn)
	}

}

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {

	*reply = "hello:" + request
	return nil
}
