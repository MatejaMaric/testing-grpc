package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/MatejaMaric/testing-grpc/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedMessengerServer
}

func (Server) GetAge(_ context.Context, person *pb.Person) (*pb.Msg, error) {
	res := &pb.Msg{
		Age: uint32(time.Now().Year()) - person.DateOfBirth,
	}

	return res, nil
}

func main() {
	tcpListener, err := net.Listen("tcp", ":2048")
	if err != nil {
		log.Fatalln("failed to listen on port 2048:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessengerServer(grpcServer, &Server{})

	log.Println("Listening on port 2048")
	if err := grpcServer.Serve(tcpListener); err != nil {
		log.Fatalln("grpc server failed to serve:", err)
	}
}
