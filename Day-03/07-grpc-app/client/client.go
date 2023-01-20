package main

import (
	"context"
	"fmt"
	"log"

	"grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)

	ctx := context.Background()
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	resp, err := client.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.GetResult())
}
