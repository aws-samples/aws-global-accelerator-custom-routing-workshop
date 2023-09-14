package client

import (
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

func StartUdpTest(host string, port int, count int) {
	logrus.Info("Start UDP test")
	// connect to udb server and send message
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()

	var buf = make([]byte, 4096)
	for i := 0; i < count; i++ {
		// send message
		msg := fmt.Sprintf("Hello World %d", i)
		startTime := time.Now()
		_, err = conn.Write([]byte(msg))
		if err != nil {
			logrus.Error("Error sending message", err.Error())
			return
		}
		fmt.Println("Message sent: ", msg)

		// read message time out process
		c1 := make(chan string, 1)
		var n int
		go func() {
			n, err = conn.Read(buf)
			c1 <- "done"
		}()

		select {
		case <-c1:
		case <-time.After(5 * time.Second):
			logrus.Error("received message time out")
			time.Sleep(1 * time.Second)
			continue
		}
		if err != nil {
			logrus.Error("Error reading message", err.Error())
			return
		}
		endTime := time.Now()
		logrus.Infof("Use time: %s Message received: %s", endTime.Sub(startTime), string(buf[:n]))
		time.Sleep(1 * time.Second)
	}
}
