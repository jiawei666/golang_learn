package main

import (
	"jiawei666/golang_learn/advance_go/mytype"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request *mytype.String, reply *mytype.String) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request *mytype.String, reply *mytype.String) error {
	reply.Value = "hello:" + request.Value
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
