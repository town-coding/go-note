package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 调用 Hello 方法
	var reply string
	err = client.Call("ExampleServiceImpl.Hello", "wx", &reply)
	if err != nil {
		log.Fatal("ExampleServiceImpl.Hello error:", err)
	}
	fmt.Println("reply:", reply)

	// 调用Ping方法
	var pingReply int
	err = client.Call("ExampleServiceImpl.Ping", 1, &pingReply)
	if err != nil {
		log.Fatal("ExampleServiceImpl.Ping error:", err)
	}
	fmt.Println("reply:", pingReply)

	// 调用Info方法
	user := User{Username: "Alice", Age: 30}
	var userReply User
	err = client.Call("ExampleServiceImpl.Info", user, &userReply)
	if err != nil {
		log.Fatal("ExampleServiceImpl.Info error:", err)
	}
	fmt.Println("reply:", userReply)

}

type User struct {
	Username string
	Age      int
}
