package main

// Esse arquivo serve para simular o comportamento de uma controller de uma api rest para a inserção de dados que foram implementados nos services
import (
	"context"
	"fmt"
	"go-grpc/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "haha",
		Email: "haha@gmail",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "1",
		Name:  "Igor",
		Email: "igotaraujo@gmail",
	}
	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err != nil {
			log.Fatalf("Could not receive response %v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println("Status: ", stream.Status, "-", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Igor",
			Email: "igotaraujo@gmail",
		},
		&pb.User{
			Id:    "2",
			Name:  "Allan",
			Email: "Allan@gmail",
		},
		&pb.User{
			Id:    "3",
			Name:  "Daniel",
			Email: "Daniel@gmail",
		},
		&pb.User{
			Id:    "4",
			Name:  "Joao",
			Email: "Joao@gmail",
		},
	}
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatal("Error creating request: %v", err)
	}
	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving request: %v", err)
	}
	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "Wesley",
			Email: "wes@wes.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "Wesley 2",
			Email: "wes2@wes.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "Wesley 3",
			Email: "wes3@wes.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "Wesley 4",
			Email: "wes4@wes.com",
		},
		&pb.User{
			Id:    "w5",
			Name:  "Wesley 5",
			Email: "wes5@wes.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Recebendo user %v com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect grpc ERROR %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// Adiciona um user
	// AddUser(client)
	// Adicionar User com stream client
	// AddUserVerbose(client)
	// Adiciona User com stream server e retorna um array das mensagens enviadas
	AddUsers(client)
	// Adicionar User com stream e o server retornar assim que o stream terminar
	// AddUserStreamBoth(client)
}
