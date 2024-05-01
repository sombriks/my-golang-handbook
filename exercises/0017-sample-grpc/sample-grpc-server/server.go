package main

import (
	"context"
	"log"
	"sample-grpc-server/db/gen"
	pb "sample-grpc-server/protos"
)

// TodoServer - receiver for protos.TodoServiceServer implementation
type TodoServer struct {
	pb.UnimplementedTodoServiceServer
	database *gen.Queries
	ctx      *context.Context
}

func (s *TodoServer) List(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	var response pb.TodoResponse
	log.Printf("list todos for q=%s\n", *request.Q)
	result, err := s.database.List(*s.ctx, *request.Q)
	if err != nil {
		log.Printf("[WARN] %s\n", err.Error())
	}
	log.Print(result)
	return &response, nil
}

func (s *TodoServer) Insert(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	var response pb.TodoResponse
	log.Printf("insert todo %s\n", *request.Todo)
	return &response, nil
}

func (s *TodoServer) Find(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	var response pb.TodoResponse
	log.Printf("find todo for id=%s\n", *request.Id)
	return &response, nil
}

func (s *TodoServer) Update(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	var response pb.TodoResponse
	log.Printf("update todo for id=%s\n", *request.Id)
	return &response, nil
}

func (s *TodoServer) Delete(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	var response pb.TodoResponse
	log.Printf("delete todo for id=%d\n", *request.Id)
	return &response, nil
}
