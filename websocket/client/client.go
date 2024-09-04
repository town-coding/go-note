package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	// 连接 WebSocket 服务器
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Dial Error:", err)
	}
	defer conn.Close()

	// 在新协程中读取消息
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read Error:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// 发送消息到服务器
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello Server!"))
	if err != nil {
		log.Println("Write Error:", err)
		return
	}

	// 捕获中断信号以关闭连接
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 等待中断信号以退出程序
	<-interrupt

	log.Println("Interrupt received, closing connection...")

	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Write Close Error:", err)
		return
	}

	select {
	case <-done:
	}
}
