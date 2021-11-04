package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/tree/main/protos"
	"google.golang.org/grpc"
)



type server struct{
	pb.UnimplementedLiderServer
}

type Jugada struct {
	Id int32
	Jugador int32
	Carta int32
}

type Recibo struct {
	Id int32
}

func (s *Server) recibirJugada (ctx context.Context, jugada *Jugada) (*Jugada, error) {
	recibo := Recibo{Id: jugada.Id}
	return jugada, nil
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