package service

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(t *testing.T) {
	initEchoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	reply, err := (*client).Echo(ctx, &Request{
		Message: "aga custom routing workshop",
	})
	if err != nil {
		t.Logf("%v", reply)
		panic(err)
	}
	t.Logf("%v", reply)
}

var (
	client *EchoServiceClient
)

func initEchoClient() {

	conn, err := grpc.Dial("127.0.0.1:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = new(EchoServiceClient)
	*client = NewEchoServiceClient(conn)
}
