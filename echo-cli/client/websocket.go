package client

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func StartWebsocketTest(host string, port int, count int) {
	logrus.Info("Start Websocket Test")
	conn, err := websocket.Dial(fmt.Sprintf("ws://%s:%d/", host, port), "", fmt.Sprintf("http://%s:%d/", host, port))
	if err != nil {
		logrus.Error("Websocket Dial Error: ", err)
	}
	defer conn.Close()
	var msg = make([]byte, 4096)
	for i := 0; i < count; i++ {
		start := time.Now()
		if _, err := conn.Write([]byte(fmt.Sprintf("Hello World %d", i))); err != nil {
			logrus.Error("Websocket Write Error: ", err)
		}
		n, err := conn.Read(msg)
		if err != nil {
			logrus.Error("Websocket Read Error: ", err)
		}
		end := time.Now()
		logrus.Infof("Use time: %s Websocket Read: %s", end.Sub(start), string(msg[:n]))
		time.Sleep(1 * time.Second)
	}
}
