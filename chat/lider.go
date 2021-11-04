package chat

import (
	"context"
	"log"
	// "math/rand"
	"net"
	// "sync"
	// "time"
	pb "github.com/MrAnacletus/Lab2-Distribuidos/protos"
	"google.golang.org/grpc"
)



type server struct{
	pb.UnimplementedLiderServer
}

type Recibo struct {
	Id int32
}

func (s *server) Recibir(ctx context.Context, recibo *pb.Recibo) (*pb.Recibo, error) {
	log.Printf("Recibido: %v", recibo)
	return &pb.Recibo{Id: recibo.Id}, nil
}

func ServerLider(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLiderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main () {
	ServerLider()
}