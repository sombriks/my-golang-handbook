package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net"
	"sample-grpc-server/db/gen"
	"sample-grpc-server/protos"
)

//go:embed db/schema.sql
var ddl string

// sample-grpc-server entrypoint
func main() {
	flag.Parse()
	port := flag.Int("port", 50051, "The server port")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO move database setup to somewhere else
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}
	// set schema state
	_, err = db.ExecContext(ctx, ddl)
	if err != nil {
		log.Fatal(err)
	}
	queries := gen.New(db)

	// finally provision server
	var todoServer = TodoServer{
		database: queries,
		ctx:      &ctx,
	}
	grpcServer := grpc.NewServer()
	protos.RegisterTodoServiceServer(grpcServer, &todoServer)
	log.Print("server ready")
	log.Fatal(grpcServer.Serve(lis))
}
