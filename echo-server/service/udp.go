package service

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

func NewUdpServer(host string, port int) (server *UdpServer, err error) {
	u := &UdpServer{
		Host: host,
		Port: port,
	}
	return u, nil
}

type UdpServer struct {
	Host string
	Port int
}

func (s *UdpServer) Start() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		logrus.Fatal(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()
	for {
		s.handleClient(conn)
	}
}

// Handles incoming requests.
func (s *UdpServer) handleClient(conn *net.UDPConn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	len, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		logrus.Error("Error reading:", err.Error())
	}
	message := string(buf[:len])
	// Send a response back to person contacting us.
	response := fmt.Sprintf("received message from %s, content is: %s\n", addr.String(), message)
	_, err = conn.WriteToUDP([]byte(response), addr)
	if err != nil {
		logrus.Error("Error writing:", err.Error())
	}
}
