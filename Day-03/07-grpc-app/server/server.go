package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

func (s *appService) Add(ctx context.Context, req *proto.AddRequest) (res *proto.AddResponse, err error) {
	x := req.GetX()
	y := req.GetY()
	if dl, ok := ctx.Deadline(); ok {
		fmt.Printf("Request received with deadline : %v\n", dl)
	}
	fmt.Printf("Processing [Add] x = %d and y = %d\n", x, y)
	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Cancel signal received...")
		err = errors.New("operation cancelled")
		return
	default:
		result := x + y
		fmt.Printf("Responding [Add] x = %d y = %d and result = %d\n", x, y, result)
		res = &proto.AddResponse{
			Result: result,
		}
		return
	}
	return
}

func (s *appService) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Processing [GeneratePrimes] Start = %d and End = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Printf("Sending Prime No : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			err := serverStream.Send(res)

			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("All prime numbers generated")
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *appService) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			fmt.Println("All the data received... calculating average...")
			avg := sum / count
			res := &proto.AverageResponse{
				Average: avg,
			}
			err = serverStream.SendAndClose(res)
			return err
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		fmt.Printf("Received No : %d\n", no)
		sum += no
		count++
	}
}

func (s *appService) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := req.GetPerson()
		msg := fmt.Sprintf("Hi %s %s!", person.GetFirstName(), person.GetLastName())
		resp := &proto.GreetResponse{
			GreetMessage: msg,
		}
		time.Sleep(500 * time.Millisecond)
		e := serverStream.Send(resp)
		if e != nil {
			log.Fatalln(err)
		}
	}
	return nil
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
