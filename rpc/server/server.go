package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type ExampleService interface {
	Hello(request string, reply *string) error
	Ping(request int, reply *int) error
	Info(request User, reply *User) error
}

// ExampleServiceImpl 结构体
type ExampleServiceImpl struct {
	Name string
}

func (e ExampleServiceImpl) Hello(request string, reply *string) error {
	fmt.Printf("receive msg:%s\n", request)
	*reply = fmt.Sprintf("Hello %s", e.Name)
	return nil
}

func (e ExampleServiceImpl) Ping(request int, reply *int) error {
	fmt.Printf("receive msg:%d\n", request)
	*reply = 111
	return nil
}

func (e ExampleServiceImpl) Info(request User, reply *User) error {
	fmt.Printf("receive msg:%v\n", request)
	*reply = User{
		Username: "服务器：" + request.Username,
		Age:      request.Age + 1,
	}
	return nil
}

func main() {
	// 创建了一个新的RPC服务器实例
	server := rpc.NewServer()
	// 将服务HelloService注册到RPC服务器
	err := server.Register(&ExampleServiceImpl{
		Name: "coding",
	})
	if err != nil {
		fmt.Println("rpc register error")
		log.Fatal("register error", err)
	}
	// 监听TCP连接并处理请求
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net listen error")
		log.Fatal("net listen error", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error")
			log.Fatal("accept error", err)
		}
		go func() {
			server.ServeConn(conn)
		}()
	}
}

type User struct {
	Username string
	Age      int
}
