package websocket

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// writeTimeOut 数据写入时间限制
	writeTimeOut = 10 * time.Second
	// readTimeOut 数据读取时间限制
	readTimeOut = 60 * time.Second
	// heartBeatTime 心跳检测时间
	heartBeatTime = (readTimeOut * 9) / 10
	// 最大消息大小(字节)
	maxMessageSize = 512
)

var (
	LINE  = []byte{'\n'}
	SPACE = []byte{' '}
)

// upgrader websocket升级
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client websocket客户端
type Client struct {
	Manager *Manager
	Conn    *websocket.Conn
	Send    chan []byte
}

// NewClient 创建websocket客户端
func NewClient(manager *Manager, response http.ResponseWriter, request *http.Request) (*Client, error) {
	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Manager: manager,
		Conn:    conn,
		Send:    make(chan []byte, 256),
	}
	client.Manager.Register <- client
	return client, nil
}

// DefaultServeWs 处理websocket请求
func DefaultServeWs(manager *Manager, response http.ResponseWriter, request *http.Request) error {
	client, err := NewClient(manager, response, request)
	if err != nil {
		return err
	}
	go client.defaultRead()
	go client.defaultWrite()
	return nil
}

// ServeWs 处理websocket请求
func ServeWs(manager *Manager, response http.ResponseWriter, request *http.Request, readFunc func(client *Client, messageType int, message []byte, err error) error, writeFunc func(client *Client, messageType int, message []byte) error) error {
	client, err := NewClient(manager, response, request)
	if err != nil {
		return err
	}
	if readFunc != nil {
		go client.read(readFunc)
	} else {
		go client.defaultRead()
	}
	if writeFunc != nil {
		go client.write(writeFunc)
	} else {
		go client.defaultWrite()
	}
	return nil
}

// write 写入数据
func (client *Client) write(fn func(client *Client, messageType int, message []byte) error) {
	ticker := time.NewTicker(heartBeatTime)
	defer func() {
		ticker.Stop()
		client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			// 设置写入截止时间，防止超时
			client.Conn.SetWriteDeadline(time.Now().Add(writeTimeOut))
			// 没有数据
			if !ok {
				client.Conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}
			if err := fn(client, websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C: // 心跳监测
			client.Conn.SetWriteDeadline(time.Now().Add(writeTimeOut))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// defaultWrite 默认写入数据方法
func (client *Client) defaultWrite() {
	client.write(func(client *Client, messageType int, message []byte) error {
		// 有数据
		// 获取写入文本消息的写入器,执行写入操作
		writer, err := client.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return nil
		}
		// 写入消息
		writer.Write(message)
		n := len(client.Send)
		for i := 0; i < n; i++ {
			writer.Write(LINE)
			writer.Write(<-client.Send)
		}
		if err := writer.Close(); err != nil {
			return nil
		}
		return nil
	})
}

// read 读取数据
func (client *Client) read(fn func(client *Client, messageType int, message []byte, err error) error) {
	defer func() {
		client.Manager.UnRegister <- client
		client.Conn.Close()
	}()
	// 设置最大消息大小
	client.Conn.SetReadLimit(maxMessageSize)
	// 设置读取截止时间，防止超时
	client.Conn.SetReadDeadline(time.Now().Add(readTimeOut))
	// 设置心跳消息处理器
	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(readTimeOut))
		return nil
	})
	for {
		messageType, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		err = fn(client, messageType, message, err)
		if err != nil {
			break
		}
	}
}

// defaultRead 默认读取数据方法
func (client *Client) defaultRead() {
	client.read(func(client *Client, messageType int, message []byte, err error) error {
		// 接收到消息处理逻辑
		message = bytes.TrimSpace(bytes.Replace(message, LINE, SPACE, -1))
		client.Manager.Broadcast <- message
		return nil
	})
}
