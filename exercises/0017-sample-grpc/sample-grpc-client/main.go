package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "sample-grpc-client/protos"
)

// sample-grpc-client entrypoint
func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTodoServiceClient(conn)
	query := "dishes"
	result, err := client.List(context.TODO(), &pb.TodoRequest{
		Q: &query,
	})
	if err != nil {
		log.Fatalf("fail to call List: %v", err)
	}
	log.Printf("List: %v", result)
}
