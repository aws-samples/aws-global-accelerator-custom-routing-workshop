package client

import (
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

func StartTcpTest(host string, port int, count int) {
	logrus.Info("Start Tcp Test")
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logrus.Errorf("Error dialing: %s", err.Error())
		return
	}
	defer conn.Close()
	var buf [4096]byte
	for i := 0; i < count; i++ {
		startTime := time.Now()
		_, err := conn.Write([]byte(fmt.Sprintf("Hello World %d", i)))
		if err != nil {
			logrus.Errorf("Error writing: %s", err.Error())
			return
		}

		// read message time out process
		c1 := make(chan string, 1)
		var n int
		go func() {
			n, err = conn.Read(buf[0:])
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
			logrus.Errorf("Error reading: %s", err.Error())
			return
		}
		endTime := time.Now()
		logrus.Infof("Use Time: %s, Received: %s", endTime.Sub(startTime), string(buf[0:n]))
		time.Sleep(1 * time.Second)
	}
}
