package rpc_demo

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func TestRpcClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:8090")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hlakjdlfaj", &reply)
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	fmt.Println(reply)
}


