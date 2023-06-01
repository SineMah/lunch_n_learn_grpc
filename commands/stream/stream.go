package stream

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "mincedmind.com/grpc/proto/count"
)

func Do(args []string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Create a client instance using the connection
	client := pb.NewCounterClient(conn)

	// Create a context with cancellation support
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a stream by calling the server-side streaming method
	stream, err := client.StreamData(cancelCtx, &pb.StreamDataRequest{Data: getInput(args)})
	if err != nil {
		log.Fatalf("failed to call StreamData: %v", err)
	}

	// Receive and process the stream of responses from the server
	for {
		// Receive a response from the server
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}

		// Process the received response
		result := response.GetResult()
		log.Println(result)
	}
}

func getInput(args []string) string {
	input := "Lorem ipsum"

	if len(args) > 0 {
		input = args[0]
	}

	return input
}
