//go:build server
// +build server

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pbAverage "mincedmind.com/grpc/proto/average"
	pbCount "mincedmind.com/grpc/proto/count"
	pbHello "mincedmind.com/grpc/proto/hello"
)

type server struct {
	pbAverage.UnimplementedAverageServiceServer
	pbCount.UnimplementedCounterServer
	pbHello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pbHello.HelloRequest) (*pbHello.HelloResponse, error) {
	name := req.GetName()
	message := "Hello, " + name
	return &pbHello.HelloResponse{Message: message}, nil
}

func (s *server) StreamData(request *pbCount.StreamDataRequest, stream pbCount.Counter_StreamDataServer) error {
	data := request.GetData()

	for i := 0; i < 10; i++ {
		response := &pbCount.StreamDataResponse{
			Result: fmt.Sprintf("%d Processed data: %s", i, data),
		}

		if err := stream.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func (*server) CalculateAverage(stream pbAverage.AverageService_CalculateAverageServer) error {
	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				average := float32(sum) / float32(count)
				return stream.SendAndClose(&pbAverage.AverageResponse{
					Average: average,
				})
			}
			log.Fatalf("Error receiving client stream: %v", err)
		}

		for _, value := range req.Values {
			sum += int(value)
			count++
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbAverage.RegisterAverageServiceServer(s, &server{})
	pbCount.RegisterCounterServer(s, &server{})
	pbHello.RegisterGreeterServer(s, &server{})

	log.Println("Server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
