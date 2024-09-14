package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// ExampleService 定义接口
type ExampleService interface {
	Hello(request string, reply *string) error
	Ping(request int, reply *int) error
	Info(request User, reply *User) error
}

// ExampleServiceImpl 结构体
type ExampleServiceImpl struct {
	Name string
}

var _ ExampleService = (*ExampleServiceImpl)(nil)

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
	// 将服务 ExampleService 注册到 RPC 服务器
	err := server.Register(&ExampleServiceImpl{
		Name: "coding",
	})
	if err != nil {
		log.Println("rpc register error", err)
	}

	// 监听TCP连接并处理请求
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("net listen error", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept error", err)
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
