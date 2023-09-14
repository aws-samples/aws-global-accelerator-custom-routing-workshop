package service

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// new websocket server
func NewWebSocketServer(host string, port int) (server *WebSocketServer, err error) {
	w := &WebSocketServer{
		Host: host,
		Port: port,
	}
	return w, nil
}

// websocket server
type WebSocketServer struct {
	Host string
	Port int
}

// start websocket server
func (s *WebSocketServer) Start() (err error) {
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("/", s.echo)
	server := &http.Server{Handler: httpServeMux}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		logrus.Fatal(err)
	}
	err = server.Serve(listener)
	if err != nil {
		logrus.Fatal(err)
	}
	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *WebSocketServer) echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Websocket Upgrade error(%v), userAgent(%s)", err, r.UserAgent())
		return
	}
	defer ws.Close()
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Websocket ReadMessage error(%v), userAgent(%s)", err, r.UserAgent())
			return
		}
		// Send a response back to person contacting us.
		response := fmt.Sprintf("received message from %s, content is: %s\n", ws.RemoteAddr().String(), p)
		err = ws.WriteMessage(messageType, []byte(response))
		if err != nil {
			logrus.Errorf("Websocket WriteMessage error(%v), userAgent(%s)", err, r.UserAgent())
			return
		}
	}
}
