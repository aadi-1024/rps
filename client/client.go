package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aadi-1024/rps/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}
	client := protobuf.NewGameClient(conn)

	var id int
	var c string
	_, _ = fmt.Scanf("%v %v", &id, &c)

	act := new(protobuf.Action)
	act.PlayerId = int32(id)

	switch c {
	case "r":
		act.Move = protobuf.Moves_Rock
	case "p":
		act.Move = protobuf.Moves_Paper
	case "s":
		act.Move = protobuf.Moves_Scissors
	default:
		act.Move = protobuf.Moves_Quit
	}

	res, err := client.PlayMove(context.Background(), act)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(res.GetMsg())
}
