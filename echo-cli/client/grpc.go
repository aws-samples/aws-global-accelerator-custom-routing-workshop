package client

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/aws-samples/aws-global-accelerator-custom-routing-workshop/echo-server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGrpcTest(host string, port int, count int) {
	logrus.Infof("StartGrpcTest host:%s port:%d count:%d", host, port, count)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Error(err)
	}
	defer conn.Close()
	for i := 0; i < count; i++ {
		client := new(service.EchoServiceClient)
		*client = service.NewEchoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		start := time.Now()
		resp, err := (*client).Echo(ctx, &service.Request{Message: fmt.Sprintf("Hello World %d", i)})
		if err != nil {
			logrus.Error(err)
		}
		end := time.Now()
		logrus.Infof("Use time: %s received message: %s", end.Sub(start), resp.GetMessage())
		cancel()
		time.Sleep(1 * time.Second)
	}
}
