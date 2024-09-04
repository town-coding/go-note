package main

import "encoding/json"

type hub struct {
	connections map[*connection]bool // 维护所有活跃连接的映射
	broadcast   chan []byte          // 广播消息通道，所有连接都会收到通过这个通道传输的数据
	register    chan *connection     // 注册新连接的通道
	unregister  chan *connection     // 注销连接的通道
}

var h = hub{
	connections: make(map[*connection]bool),
	broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.sc)
			}
		case data := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.sc <- data:
					c.data.UserList = user_list
				default:
					delete(h.connections, c)
					close(c.sc)
				}
			}
		}
	}
}
