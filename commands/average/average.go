package average

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"time"

	pb "mincedmind.com/grpc/proto/average"
)

func Do() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewAverageServiceClient(cc)

	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	for _, value := range generateRandomIntegers(5) {
		req := &pb.IntStream{
			Values: []int32{int32(value)},
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		fmt.Printf(fmt.Sprintf("Sending %d \n", value))
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Printf("Average: %v\n", res.Average)
}

func generateRandomIntegers(n int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	randomIntegers := make([]int, n)
	for i := 0; i < n; i++ {
		randomIntegers[i] = random.Intn(101)
	}

	return randomIntegers
}
