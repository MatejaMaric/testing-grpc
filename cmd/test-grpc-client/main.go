package main

import (
	"context"
	"log"
	"time"

	"github.com/MatejaMaric/protobuf-testing/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":2048", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("grpc client failed to connect:", err)
	}
	defer conn.Close()

	client := pb.NewMessengerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	msg, err := client.GetAge(ctx, &pb.Person{Name: "Test", DateOfBirth: 1997})
	if err != nil {
		log.Fatalln("grpc client failed to GetAge:", err)
	}

	log.Println("gRPC response: age:", msg.Age)
}
