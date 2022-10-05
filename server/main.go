package main

import (
	"context"
	"log"
	"net"
	t "time"
	api "timeserver/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	api.UnimplementedGreeterServer
}

func (s *Server) GetTime(ctx context.Context, time *api.Time) (*api.Time, error) {
	now := t.Now()

	log.Println(now)

	ts := timestamppb.New(now)

	log.Println(ts)

	response := &api.Time{
		Time: ts,
	}

	return response, nil
}

func (s *Server) GetRoundTime(ctx context.Context, time *api.RoundTime) (*api.RoundTime, error) {
	// t2
	t2now := t.Now()
	t2ts := timestamppb.New(t2now)

	// t3
	t3now := t.Now()
	t3ts := timestamppb.New(t3now)

	response := &api.RoundTime{
		T2: t2ts,
		T3: t3ts,
	}

	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	log.Println("listen on :8080")

	if err != nil {
		log.Fatal("woops")
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	server := &Server{}

	api.RegisterGreeterServer(grpcServer, server)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("noob")
	}
}
