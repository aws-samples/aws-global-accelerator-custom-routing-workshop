package service

import (
	"fmt"
	"html"
	"net/http"

	"github.com/sirupsen/logrus"
)

// register http handler
func RegisterHttpHandler() {
	http.HandleFunc("/", echo)
}

// new http server
func NewHttpServer(host string, port int) (server *HttpServer, err error) {
	h := &HttpServer{
		Host: host,
		Port: port,
	}
	return h, nil
}

// http server
type HttpServer struct {
	Host string
	Port int
}

// start http server
func (s *HttpServer) Start() {
	// start http listen
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), nil)
	if err != nil {
		logrus.Fatal(err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("received message from %s, content is: %s\n", r.RemoteAddr, r.URL.Path)
	fmt.Fprintf(w, "%s", html.EscapeString(response))
}
