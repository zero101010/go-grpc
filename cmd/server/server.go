package main

import (
	"go-grpc/pb"
	"go-grpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Cria o listener e a porta que irá ouvir
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Cold not connect: %v", err)
	}
	grpcServer := grpc.NewServer()
	// Registra serviço
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error grpc server: %v", err)
	}
}
