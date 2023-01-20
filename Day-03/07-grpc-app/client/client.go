package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doRequestResponse(ctx, client)
	// doServerStreaming(ctx, client)
	// doClientStreaming(ctx, client)
	doBiDirectionalStreaming(ctx, client)

}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	/*
		addCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		fmt.Println("Hit ENTER to stop...")
		go func() {
			fmt.Scanln()
			cancel()
		}()
	*/

	addCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := client.Add(addCtx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Result : %d\n", res.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All responses received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Prime Number received :", res.GetPrimeNo())
	}
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	data := []int32{3, 1, 4, 2, 5, 9, 7, 8, 6}
	clientStream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range data {
		fmt.Printf("Sending no : %d\n", no)
		req := &proto.AverageRequest{
			No: no,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Average : %d\n", res.GetAverage())
}

func doBiDirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	personNames := []proto.PersonName{
		proto.PersonName{FirstName: "Magesh", LastName: "Kuppan"},
		proto.PersonName{FirstName: "Suresh", LastName: "Kannan"},
		proto.PersonName{FirstName: "Rajesh", LastName: "Pandit"},
		proto.PersonName{FirstName: "Ganesh", LastName: "Kumar"},
		proto.PersonName{FirstName: "Ramesh", LastName: "Jayaraman"},
	}
	clientStream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{})
	go func() {
		for {
			resp, err := clientStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(resp.GetGreetMessage())
		}
		done <- struct{}{}
	}()
	for _, personName := range personNames {
		// time.Sleep(500 * time.Millisecond)
		req := &proto.GreetRequest{
			Person: &personName,
		}
		fmt.Println("Sending : ", personName.FirstName, personName.LastName)
		clientStream.Send(req)
		time.Sleep(time.Second)
	}
	clientStream.CloseSend()
	<-done
}
