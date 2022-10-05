package main

import (
	"context"
	"log"

	api "timeserver/api/v1"

	t "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var server api.GreeterClient  //the server
var ServerConn *grpc.ClientConn //the server connection

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//use context for timeout on the connection
	timeContext, cancel := context.WithTimeout(context.Background(), t.Second)
	defer cancel() //cancel the connection when we are done

	//dial the server to get a connection to it
	conn, err := grpc.DialContext(timeContext, "localhost:8080", opts...)

	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	// makes a client from the server connection and saves the connection
	// and prints rather or not the connection was is READY
	server = api.NewGreeterClient(conn)
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())

	t1now := t.Now()

	time, err := server.GetRoundTime(context.Background(), &api.RoundTime{})

	t4now := t.Now()

	if err != nil {
		log.Fatal("fuck off")
	}

	serverdiff := time.T3.AsTime().Sub(time.T2.AsTime())

	roundtriptime := t4now.Sub(t1now) - serverdiff

	clientclock := time.T3.AsTime().Add(roundtriptime/2)

	log.Println(clientclock.Local())
}
