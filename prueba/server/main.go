package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto/github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)
type server struct{
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) GetJugadores(ctx context.Context, in *pb.Jugadores) (*pb.HelloReply, error) {
	fmt.Println("Jugadores recibidos ", in.Name)
	return &pb.HelloReply{Message: "Jugadores recibidos, gracias" + in.Name}, nil
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
	if err = serv.Serve(listener); err != nil{
		panic(err)
	}
}