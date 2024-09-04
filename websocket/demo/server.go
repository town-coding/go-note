package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 处理函数
	http.HandleFunc("/ws", handleWebSocket)
	go h.run()
	fmt.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
