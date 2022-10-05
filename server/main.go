package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	t "time"
	api "timeserver/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	api.UnimplementedGreeterServer
	port string
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
	p, _ := peer.FromContext(ctx)
	log.Println(p.Addr.String())

	log.Println("Client invoked GetRoundTime")

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

var name = flag.String("name", "localhost", "Senders name")
var port = flag.String("port", "8080", "Tcp server")

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *name, *port))
	log.Printf("listen on %s:%s\n", *name, *port)

	if err != nil {
		log.Fatal("woops")
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	server := &Server{
		port: *port,
	}

	api.RegisterGreeterServer(grpcServer, server)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("noob")
	}
}
