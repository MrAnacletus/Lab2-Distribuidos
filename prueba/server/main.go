package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)
type server struct{
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ")
	return &pb.HelloReply{Message: "Hello "}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ")
	return &pb.HelloReply{Message: "Hello "}, nil
}

func (s *server) GetJugadas(ctx context.Context, in *pb.Jugadores) (*pb.HelloReply, error) {
	fmt.Println("Jugadas recibidas ")
	return &pb.HelloReply{Message: "Jugadas recibidas, gracias"}, nil
}

func main(){
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	serv := grpc.NewServer()
	pb.RegisterHelloServiceServer(serv, &server{})
	if err = serv.Serve(listener); err != nil{
		panic(err)
	}
}