package app

import (
	"runtime"

	"github.com/zengqiang96/mattermostclone/model"
)

const (
	broadcastQueueSize = 4096
)

type Hub struct {
	connectionCount int64
	app             *App
	connectionIndex int
	register        chan *WebConn
	unregister      chan *WebConn
	broadcast       chan *model.WebSocketEvent
	stop            chan struct{}
}

func (a *App) NewWebHub() *Hub {
	return &Hub{
		app:        a,
		register:   make(chan *WebConn),
		unregister: make(chan *WebConn),
		broadcast:  make(chan *model.WebSocketEvent, broadcastQueueSize),
		stop:       make(chan struct{}),
	}
}

func (s *Server) Publish(message *model.WebSocketEvent) {
	s.PublishSkipClusterMessage(message)
}

func (s *Server) PublishSkipClusterMessage(event *model.WebSocketEvent) {
	if event.GetBroadcast().UserId != "" {

	} else {
		for _, hub := range s.hubs {
			hub.Broadcast(event)
		}
	}
}

func (a *App) Publish(message *model.WebSocketEvent) {
	a.Srv().Publish(message)
}

func (a *App) HubStart() {
	numberOfHubs := runtime.NumCPU() * 2
	hubs := make([]*Hub, numberOfHubs)

	for i := 0; i < numberOfHubs; i++ {
		hubs[i] = a.NewWebHub()
		hubs[i].connectionIndex = i
		hubs[i].Start()
	}
	a.srv.hubs = hubs
}

func (h *Hub) Start() {
	var doStart func()
	var doRecoverableStart func()
	var doRecover func()

	doStart = func() {
		connIndex := newHubConnectionIndex()
		for {
			select {
			case msg := <-h.broadcast:
				broadcast := func(webConn *WebConn) {

				}

				if msg.GetBroadcast().UserId != "" {

				}
				candidates := connIndex.All()
				for webConn := range candidates {
					broadcast(webConn)
				}
			case <-h.stop:
				for webConn := range connIndex.All() {
					webConn.Close()
				}
			}
		}
	}

	doRecoverableStart = func() {
		defer doRecover()
		doStart()
	}

	doRecover = func() {
		go doRecoverableStart()
	}

	go doRecoverableStart()
}

type hubConnectionIndex struct {
	byUserId     map[string][]*WebConn
	byConnection map[*WebConn]int
}

func newHubConnectionIndex() *hubConnectionIndex {
	return &hubConnectionIndex{
		byUserId:     make(map[string][]*WebConn),
		byConnection: make(map[*WebConn]int),
	}
}

func (i *hubConnectionIndex) All() map[*WebConn]int {
	return i.byConnection
}

func (h *Hub) Broadcast(message *model.WebSocketEvent) {
	if h != nil && message != nil {
		select {
		case h.broadcast <- message:
		case <-h.stop:
		}
	}
}
