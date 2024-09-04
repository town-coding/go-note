package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 定义一个 WebSocket 升级器，用于将 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	// 允许所有来源
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理 WebSocket 连接的函数
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	for {
		// 读取客户端发送的消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		// 打印接收到的消息
		log.Printf("Received: %s", message)

		// 发送消息回客户端
		err = conn.WriteMessage(messageType, []byte("Hello from Server!"))
		if err != nil {
			log.Println("Write Error:", err)
			break
		}
	}
}

func main() {
	// 处理函数
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
