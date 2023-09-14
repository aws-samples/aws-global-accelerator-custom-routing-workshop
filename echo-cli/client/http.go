package client

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func StartHttpTest(host string, port int, count int) {
	logrus.Info("start http test")

	for i := 0; i < count; i++ {
		client := &http.Client{}
		startTime := time.Now()
		result, err := client.Get(fmt.Sprintf("http://%s:%d/hello_world_%d", host, port, i))
		if err != nil {
			logrus.Error("http get error: ", err)
			return
		}
		message, err := io.ReadAll(result.Body)
		if err != nil {
			logrus.Error("read body error: ", err)
			return
		}
		endTime := time.Now()
		logrus.Infof("Use time: %s http get result: %s", endTime.Sub(startTime), string(message))
		client.CloseIdleConnections()
		time.Sleep(1 * time.Second)
	}
}
