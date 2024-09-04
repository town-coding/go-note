package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var user_list = []string{}

// 定义一个 WebSocket 升级器，用于将 HTTP 连接升级为 WebSocket 连接
var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{ws: conn, sc: make(chan []byte, 256), data: &Data{}}
	h.register <- c
	go c.writer()
	c.reader()
	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.broadcast <- data_b
		h.unregister <- c
	}()

}

func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.register <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.broadcast <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.broadcast <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.broadcast <- data_b
			h.unregister <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
