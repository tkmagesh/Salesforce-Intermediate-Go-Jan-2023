package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

func (svc *appService) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	resp := &proto.AddResponse{
		Result: result,
	}
	return resp, nil
}

func main() {
	asi := &appService{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
