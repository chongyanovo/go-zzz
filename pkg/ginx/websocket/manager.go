package websocket

// Manager ws客户端管理器
type Manager struct {
	Clients    map[*Client]struct{}
	Broadcast  chan []byte
	Register   chan *Client
	UnRegister chan *Client
}

// NewWsManager 创建ws客户端管理器
func NewWsManager() *Manager {
	return &Manager{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]struct{}),
	}
}

// Run 启动ws客户端管理器
func (manager *Manager) Run() {
	for {
		select {
		case client := <-manager.Register:
			manager.Clients[client] = struct{}{}
		case client := <-manager.UnRegister:
			if _, ok := manager.Clients[client]; ok {
				delete(manager.Clients, client)
				close(client.Send)
			}
		case message := <-manager.Broadcast:
			for client := range manager.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(manager.Clients, client)
				}
			}
		}
	}
}
