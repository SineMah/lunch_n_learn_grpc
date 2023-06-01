package hello

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "mincedmind.com/grpc/proto/hello"
)

func Do(args []string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := getName(args)

	req := &pb.HelloRequest{Name: name}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	log.Println(res.GetMessage())
}

func getName(names []string) string {
	name := "John Smith"

	if len(names) > 0 {
		name = names[0]
	}

	return name
}
