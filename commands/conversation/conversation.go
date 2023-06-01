package conversation

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "mincedmind.com/grpc/proto/conversation"
	"os"
)

var firstShell bool = false

func Do() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewConversationServiceClient(conn)

	stream, err := client.StartConversation(context.Background())
	if err != nil {
		log.Fatalf("Failed to start conversation: %v", err)
	}

	go func() {
		for {
			message := getMessage()

			req := &pb.ConversationMessage{
				Text: message,
			}
			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed to send message: %v", err)
			}
		}
	}()

	// Receive messages
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}

		log.Printf("Received response: %s", resp.Text)
	}
}

func getMessage() string {
	var message string

	if firstShell == false {
		log.Println("Enter your message (or 'quit' to exit):")
	}

	firstShell = true

	_, err := fmt.Scanln(&message)

	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	if message == "quit" {
		os.Exit(0)
	}

	return message
}
