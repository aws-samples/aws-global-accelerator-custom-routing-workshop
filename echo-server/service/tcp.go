package service

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

func NewTcpServer(host string, port int) (*TcpServer, error) {
	s := &TcpServer{
		Host: host,
		Port: port,
	}
	return s, nil
}

type TcpServer struct {
	Host string
	Port int
}

func (s *TcpServer) Start() {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		logrus.Fatal("Error listening: ", err.Error())
	}
	// Close the listener when the application closes.
	defer l.Close()
	logrus.Infof("Listening on %s:%d", s.Host, s.Port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			logrus.Error("Error accepting: ", err.Error())
		}
		logrus.Infof("Received message from %s", conn.RemoteAddr().String())
		// Handle connections in a new goroutine.
		go s.handleRequest(conn)
	}
}

// Handles incoming requests.
func (s *TcpServer) handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	for {
		len, err := conn.Read(buf)
		if err != nil {
			logrus.Error("Error reading: ", err.Error())
			break
		}
		// Send a response back to person contacting us.
		message := string(buf[:len])
		response := fmt.Sprintf("received message from %s, content is: %s\n", conn.RemoteAddr().String(), message)
		conn.Write([]byte(response))
		if string(buf[:4]) == "exit" {
			conn.Write([]byte("bye\n"))
			break
		}
	}
	conn.Close()
}
