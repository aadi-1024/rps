package main

import (
	"log"
	"net"

	"github.com/aadi-1024/rps/protobuf"
	"github.com/aadi-1024/rps/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	srv := grpc.NewServer()

	gameServer, resolver := server.NewGameServer()
	protobuf.RegisterGameServer(srv, &gameServer)

	go resolver()
	if err := srv.Serve(lis); err != nil {
		log.Fatalln(err.Error())
	}
}
