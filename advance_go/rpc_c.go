package main

import (
	"fmt"
	"jiawei666/golang_learn/advance_go/mytype"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceInterface = interface {
	Hello(request *mytype.String, reply *mytype.String) error
}

type HelloServiceClient struct {
	Client *rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

const HelloServiceName = "path/to/pkg.HelloService"

func (p *HelloServiceClient) Hello(request *mytype.String, reply *mytype.String) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply mytype.String
	err = client.Call(HelloServiceName+".Hello", mytype.String{Value: "jiawei"}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
