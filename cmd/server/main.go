package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/MatejaMaric/testing-grpc/pkg/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	go runHttpGateway()
	runGrpcServer()
}

func runGrpcServer() {
	tcpListener, err := net.Listen("tcp", ":2048")
	if err != nil {
		log.Fatalln("failed to listen on port 2048:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessengerServer(grpcServer, &Server{})

	log.Println("grpc listening on port 2048")
	if err := grpcServer.Serve(tcpListener); err != nil {
		log.Fatalln("grpc server failed to serve:", err)
	}
}

func runHttpGateway() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	endpoint := "localhost:2048"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := pb.RegisterMessengerHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalln("http gateway failed to register handler:", err)
	}

	log.Println("http gateway listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("http gateway failed to serve:", err)
	}
}
